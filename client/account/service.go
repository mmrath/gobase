package account

import (
	"context"
	"crypto/sha256"
	"fmt"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/mmrath/gobase/common/auth"
	"github.com/mmrath/gobase/common/crypto"
	"github.com/mmrath/gobase/common/error_util"
	"github.com/mmrath/gobase/model"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
)

type Service struct {
	notifier Notifier
	db       *model.DB
}

func NewService(notifier Notifier, db *model.DB) *Service {
	return &Service{notifier: notifier, db: db}
}

func (s *Service) Activate(token string) error {
	err := validation.Validate(token,
		validation.Required,
		validation.Length(4, 128),
	)
	if err != nil {
		return err
	}

	tokenHash := fmt.Sprintf("%x", sha256.Sum256([]byte(token)))

	err = s.db.RunTx(func(tx *model.Tx) error {

		uc, err := tx.UserCredentialDao().GetByActivationKey(tokenHash)
		if err != nil {
			if model.IsNoDataFound(err) {
				return error_helper.NewBadRequest("invalid activation token")
			}
			return err
		}
		if !uc.Activated {
			if uc.ActivationKeyExpiresAt.Before(time.Now()) {
				return error_helper.NewBadRequest("activation token is expired, sign up again")
			}
			err = tx.UserCredentialDao().Activate(uc.ID)
			if err != nil {
				return err
			}
		}
		return nil
	})

	return err
}

func (s *Service) Login(login *model.LoginRequest) (*model.User, error) {
	err := login.Validate()
	if err != nil {
		return nil, err
	}
	var user *model.User

	err = s.db.RunTx(func(tx *model.Tx) error {
		user, err = tx.UserDao().GetByEmail(login.Email)
		if err != nil {
			if model.IsNoDataFound(err) {
				return error_helper.NewUnauthorized("invalid email or password")
			}
			return err
		}

		if !user.Active {
			return error_helper.NewUnauthorized("user is not active")
		}

		uc, err := tx.UserCredentialDao().Get(user.ID)
		if err != nil {
			if model.IsNoDataFound(err) {
				return error_helper.NewUnauthorized("invalid email or password")
			}
			return err
		}

		if !uc.Activated {
			return error_helper.NewUnauthorized("user is not activated")
		} else if !uc.ExpiresAt.IsZero() && uc.ExpiresAt.Before(time.Now()) {
			return error_helper.NewUnauthorized("password expired")
		} else if uc.Locked {
			return error_helper.NewUnauthorized("account is locked")
		}

		passwordSha := crypto.SHA256([]byte(login.Password))
		matched, err := crypto.CheckPassword(passwordSha, uc.PasswordHash)

		if err != nil {
			return err
		}

		if !matched {
			if uc.InvalidAttempts >= 3 {
				err := tx.UserCredentialDao().IncrementInvalidAttempts(user.ID, true)
				if err != nil {
					return err
				}
			} else {
				err := tx.UserCredentialDao().IncrementInvalidAttempts(user.ID, false)
				if err != nil {
					return err
				}
			}
			return error_helper.NewUnauthorized("invalid email or password")
		}

		if uc.InvalidAttempts > 0 {
			err := tx.UserCredentialDao().ResetInvalidAttempts(user.ID)
			if err != nil {
				log.Error().Err(err).Msg("failed resetting invalid attempts")
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Service) ChangePassword(ctx context.Context, data *model.ChangePasswordRequest) error {

	id := auth.UserIdFromContext(ctx)

	err := s.db.RunTx(func(tx *model.Tx) error {
		uc, err := tx.UserCredentialDao().Get(id)

		if err != nil {
			if model.IsNoDataFound(err) {
				return error_helper.NewBadRequest("password cannot be changed as user does not exist")
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
				err := tx.UserCredentialDao().IncrementInvalidAttempts(id, true)
				if err != nil {
					return err
				}
			} else {
				err := tx.UserCredentialDao().IncrementInvalidAttempts(id, false)
				if err != nil {
					return err
				}
			}
			return error_helper.NewUnauthorized("invalid current password")
		}

		newPasswordHash, err := crypto.HashPassword(crypto.SHA256([]byte(data.NewPassword)))
		if err != nil {
			return error_helper.NewInternal(err, "failed to hash password")
		}

		err = tx.UserCredentialDao().ChangePassword(uc.ID, newPasswordHash)
		return err
	})

	return err
}

func (s *Service) InitiatePasswordReset(email string) error {

	var user *model.User
	var err error
	resetToken := uuid.New().String()
	resetTokenSha := fmt.Sprintf("%x", sha256.Sum256([]byte(resetToken)))
	expiresAt := time.Now().Add(20 * time.Minute)

	err = s.db.RunTx(func(tx *model.Tx) error {
		user, err = tx.UserDao().GetByEmail(email)
		if err != nil {
			if model.IsNoDataFound(err) {
				return error_helper.NewBadRequest("user not found")
			}
			return err
		}
		_, err = tx.UserCredentialDao().Get(user.ID)

		if err != nil {
			if model.IsNoDataFound(err) {
				cred := model.UserCredential{
					ID:                user.ID,
					ResetKey:          resetTokenSha,
					ResetKeyExpiresAt: expiresAt,
				}
				err = tx.UserCredentialDao().Insert(&cred)
				return err
			} else {
				return err
			}
		} else {
			err = tx.UserCredentialDao().UpdateResetKey(user.ID, resetTokenSha, expiresAt)
			return err
		}
	})

	if err != nil {
		return err
	}
	err = s.notifier.NotifyPasswordResetInit(user, resetToken)

	if err != nil {
		log.Error().Err(err).Int64("id", user.ID).Msg("failed to send password reset email")
		return error_helper.NewInternal(err, "Failed to send password reset email")
	}

	log.Info().Int64("id", user.ID).Msg("successfully sent password reset email")
	return nil
}

func (s *Service) ResetPassword(passwordResetRequest *model.ResetPasswordRequest) error {
	resetTokenSha := fmt.Sprintf("%x", sha256.Sum256([]byte(passwordResetRequest.ResetToken)))

	err := s.db.RunTx(func(tx *model.Tx) error {
		uc, err := tx.UserCredentialDao().FindByResetKey(resetTokenSha)

		if err != nil {
			if model.IsNoDataFound(err) {
				return error_helper.NewBadRequest("reset key is invalid")
			}
			return err
		}

		if uc.ResetKeyExpiresAt.Before(time.Now()) {
			return error_helper.NewBadRequest("reset key is expired")
		}

		passwordHash, err := crypto.HashPassword(crypto.SHA256([]byte(passwordResetRequest.NewPassword)))
		if err != nil {
			return error_helper.NewInternal(err, "failed to hash password")
		}

		err = tx.UserCredentialDao().ResetPassword(uc.ID, passwordHash)
		return err
	})

	return err
}

func (s *Service) SignUp(signUpRequest *model.SignUpRequest) (*model.User, error) {

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
		return nil, error_helper.NewInternal(err, "failed to hash password")
	}

	activationToken := uuid.New().String()
	activationTokenHash := fmt.Sprintf("%x", sha256.Sum256([]byte(activationToken)))
	err = s.db.RunTx(func(tx *model.Tx) error {
		err = s.checkForDuplicate(signUpRequest.Email, "email", tx.UserDao().ExistsByEmail)
		if err != nil {
			return err
		}

		err = tx.UserDao().Insert(&newUser)
		if err != nil {
			return error_helper.NewInternal(err, "failed to insert user")
		}

		cred := model.UserCredential{
			ID:                     newUser.ID,
			PasswordHash:           passwordHash,
			ActivationKey:          activationTokenHash,
			ActivationKeyExpiresAt: time.Now().Add(time.Second * 1200),
		}

		err = tx.UserCredentialDao().Insert(&cred)
		if err != nil {
			return error_helper.NewInternal(err, "Internal error - unable to insert user_credential")
		}
		return nil
	})

	if err != nil {
		return nil, error_helper.ToError(err, "failed to complete sign-up transaction")
	}

	err = s.notifier.NotifyActivation(&newUser, activationToken)

	if err != nil {
		log.Error().Err(err).Msg("failed to send account activation email")
		return nil, error_helper.NewInternal(err, "failed to send account activation email")
	}

	log.Debug().Interface("user", newUser).Msg("successfully signed up user")
	return &newUser, nil
}

func (s *Service) GetProfile(ctx context.Context) (*model.User, error) {
	id := auth.UserIdFromContext(ctx)

	var user *model.User
	var err error
	err = s.db.RunTx(func(tx *model.Tx) error {
		user, err = tx.UserDao().Get(id)
		return err
	})
	return user, err
}

func (s *Service) UpdateProfile(user *model.User) error {
	return nil
}

func (s *Service) checkForDuplicate(input string, by string, fn func(string) (bool, error)) error {
	exists, err := fn(input)
	if err != nil {
		return error_helper.NewInternal(err, "Error while checking for duplicate email")
	} else if exists {
		log.Info().Str(by, input).Msgf("found user with same %s", by)
		fieldErrors := []error_helper.FieldError{{Field: "email", Message: "user already exists"}}
		return error_helper.WithFieldErrors(fieldErrors)
	}
	return nil
}