package account

import (
	"context"
	"crypto/sha256"
	"fmt"
	"strings"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"

	"github.com/mmrath/gobase/common/auth"
	"github.com/mmrath/gobase/common/crypto"
	"github.com/mmrath/gobase/common/error_util"
	"github.com/mmrath/gobase/model"
)

type service struct {
	notifier      Notifier
	db            *model.DB
	credentialDao model.UserCredentialDao
	userDao       model.UserDao
}

type Service interface {
	SignUp(*model.SignUpRequest) (*model.User, error)
	Activate(token string) error
	InitiatePasswordReset(email string) error
	ResetPassword(*model.ResetPasswordRequest) error
	ChangePassword(ctx context.Context, data *model.ChangePasswordRequest) error
}

type Notifier interface {
	NotifyActivation(user *model.User, token string) error
	NotifyPasswordChange(user *model.User) error
	NotifyPasswordResetInit(user *model.User, token string) error
}

func (s *service) ChangePassword(ctx context.Context, data *model.ChangePasswordRequest) error {

	id := auth.UserIdFromContext(ctx)

	err := s.db.Tx(ctx, func(ctx context.Context) error {
		uc, err := s.credentialDao.Get(ctx, id)

		if err != nil {
			if model.IsNoDataFound(err) {
				return error_util.NewBadRequest("password cannot be changed as user does not exist")
			}
			return err
		}

		currentPasswordSha := crypto.SHA256([]byte(data.CurrentPassword))
		matched, err := crypto.CheckPassword(currentPasswordSha, uc.PasswordHash)

		if err != nil {
			return err
		}

		if !matched {
			if uc.InvalidAttempts >= 3 {
				err := s.credentialDao.IncrementInvalidAttempts(ctx, id, true)
				if err != nil {
					return err
				}
			} else {
				err := s.credentialDao.IncrementInvalidAttempts(ctx, id, false)
				if err != nil {
					return err
				}
			}
			return error_util.NewUnauthorized("invalid current password")
		}

		newPasswordHash, err := crypto.HashPassword(crypto.SHA256([]byte(data.NewPassword)))
		if err != nil {
			return error_util.NewInternal(err, "failed to hash password")
		}

		err = s.credentialDao.ChangePassword(ctx, uc.ID, newPasswordHash)
		return err
	})

	return err
}

func (s *service) InitiatePasswordReset(email string) error {

	var user *model.User
	var err error
	resetToken := uuid.New().String()
	resetTokenSha := fmt.Sprintf("%x", sha256.Sum256([]byte(resetToken)))
	expiresAt := time.Now().Add(20 * time.Minute)

	err = s.db.Tx(context.Background(), func(ctx context.Context) error {
		user, err = s.userDao.FindByEmail(ctx, email)
		if err != nil {
			if model.IsNoDataFound(err) {
				return error_util.NewBadRequest("user not found")
			}
			return err
		}
		_, err = s.credentialDao.Get(ctx, user.ID)

		if err != nil {
			if model.IsNoDataFound(err) {
				cred := model.UserCredential{
					ID:                user.ID,
					ResetKey:          resetTokenSha,
					ResetKeyExpiresAt: expiresAt,
				}
				err = s.credentialDao.Insert(ctx, &cred)
				return err
			} else {
				return err
			}
		} else {
			err = s.credentialDao.UpdateResetKey(ctx, user.ID, resetTokenSha, expiresAt)
			return err
		}
	})

	if err != nil {
		return err
	}
	err = s.notifier.NotifyPasswordResetInit(user, resetToken)

	if err != nil {
		log.Error().Err(err).Int64("id", user.ID).Msg("failed to send password reset email")
		return error_util.NewInternal(err, "Failed to send password reset email")
	}

	log.Info().Int64("id", user.ID).Msg("successfully sent password reset email")
	return nil
}

func (s *service) ResetPassword(passwordResetRequest *model.ResetPasswordRequest) error {
	resetTokenSha := fmt.Sprintf("%x", sha256.Sum256([]byte(passwordResetRequest.ResetToken)))

	err := s.db.Tx(context.Background(), func(ctx context.Context) error {
		uc, err := s.credentialDao.FindByResetKey(ctx, resetTokenSha)

		if err != nil {
			if model.IsNoDataFound(err) {
				return error_util.NewBadRequest("reset key is invalid")
			}
			return err
		}

		if uc.ResetKeyExpiresAt.Before(time.Now()) {
			return error_util.NewBadRequest("reset key is expired")
		}

		passwordHash, err := crypto.HashPassword(crypto.SHA256([]byte(passwordResetRequest.NewPassword)))
		if err != nil {
			return error_util.NewInternal(err, "failed to hash password")
		}

		err = s.credentialDao.ResetPassword(ctx, uc.ID, passwordHash)
		return err
	})

	return err
}

func (s *service) Activate(token string) error {
	err := validation.Validate(token,
		validation.Required,
		validation.Length(4, 128),
	)
	if err != nil {
		return err
	}

	tokenHash := fmt.Sprintf("%x", sha256.Sum256([]byte(token)))

	err = s.db.Tx(context.Background(), func(ctx context.Context) error {

		uc, err := s.credentialDao.GetByActivationKey(ctx, tokenHash)
		if err != nil {
			if model.IsNoDataFound(err) {
				return error_util.NewBadRequest("Invalid activation token")
			}
			return err
		}
		if !uc.Activated {
			if uc.ActivationKeyExpiresAt.Before(time.Now()) {
				return error_util.NewBadRequest("Activation token is expired, sign up again")
			}
			err = s.credentialDao.Activate(ctx, uc.ID)
			if err != nil {
				return err
			}
		}
		return nil
	})

	return err
}

func (s *service) SignUp(signUpRequest *model.SignUpRequest) (*model.User, error) {

	log.Debug().Interface("email", signUpRequest.Email).Msg("signing up user")
	err := signUpRequest.Validate()

	if err != nil {
		return nil, err
	}

	newUser := model.User{
		AuditDetails: model.AuditDetails{CreatedBy: "SIGNUP"},
		FirstName:    signUpRequest.FirstName,
		LastName:     signUpRequest.LastName,
		Email:        strings.ToLower(signUpRequest.Email),
		Active:       true,
	}

	newUser.CreatedBy = "SIGNUP"
	newUser.UpdatedBy = newUser.CreatedBy

	passwordHash, err := crypto.HashPassword(crypto.SHA256([]byte(signUpRequest.Password)))
	if err != nil {
		return nil, error_util.NewInternal(err, "failed to hash password")
	}

	activationToken := uuid.New().String()
	activationTokenHash := fmt.Sprintf("%x", sha256.Sum256([]byte(activationToken)))
	err = s.db.Tx(context.Background(), func(ctx context.Context) error {
		err = s.checkForDuplicate(ctx, signUpRequest.Email, "email", s.userDao.ExistsByEmail)
		if err != nil {
			return err
		}

		err = s.userDao.Insert(ctx, &newUser)
		if err != nil {
			return error_util.NewInternal(err, "failed to insert user")
		}

		cred := model.UserCredential{
			ID:                     newUser.ID,
			PasswordHash:           passwordHash,
			ActivationKey:          activationTokenHash,
			ActivationKeyExpiresAt: time.Now().Add(time.Second * 1200),
		}

		err = s.credentialDao.Insert(ctx, &cred)
		if err != nil {
			return error_util.NewInternal(err, "Internal error - unable to insert user_credential")
		}
		return nil
	})

	if err != nil {
		return nil, error_util.ToError(err, "failed to complete sign-up transaction")
	}

	err = s.notifier.NotifyActivation(&newUser, activationToken)

	if err != nil {
		log.Error().Err(err).Msg("failed to send account activation email")
		return nil, error_util.NewInternal(err, "failed to send account activation email")
	}

	log.Debug().Interface("user", newUser).Msg("successfully signed up user")
	return &newUser, nil
}

func (s *service) checkForDuplicate(ctx context.Context, input string, by string, fn func(context.Context, string) (bool, error)) error {
	exists, err := fn(ctx, input)
	if err != nil {
		return error_util.NewInternal(err, "Error while checking for duplicate email")
	} else if exists {
		log.Info().Str(by, input).Msgf("found user with same %s", by)
		fieldErrors := []error_util.FieldError{{Field: "email", Message: "user already exists"}}
		return error_util.WithFieldErrors(fieldErrors)
	}
	return nil
}
