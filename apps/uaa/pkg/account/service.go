package account

import (
	"context"
	"crypto/sha256"
	"fmt"
	"github.com/mmrath/gobase/pkg/db"
	"github.com/mmrath/gobase/pkg/error_util"
	"strings"
	"time"

	"github.com/mmrath/gobase/pkg/auth"
	"github.com/mmrath/gobase/pkg/crypto"
	"github.com/mmrath/gobase/pkg/model"
	"github.com/rs/zerolog/log"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
)

type Service struct {
	notifier          Notifier
	db                *db.DB
	userCredentialDao model.UserCredentialDao
	userDao           model.UserDao
}

func NewService(notifier Notifier, db *db.DB) *Service {
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

	err = s.db.Tx(context.Background(), func(tx *db.Tx) error {

		uc, err := s.userCredentialDao.GetByActivationKey(tx, tokenHash)
		if err != nil {
			if db.IsNoDataFound(err) {
				return error_util.NewBadRequest("invalid activation token")
			}
			return err
		}
		if !uc.Activated {
			if uc.ActivationKeyExpiresAt.Before(time.Now()) {
				return error_util.NewBadRequest("activation token is expired, sign up again")
			}
			err = s.userCredentialDao.Activate(tx, uc.ID)
			if err != nil {
				return err
			}
		}
		return nil
	})

	return err
}

func (s *Service) Login(login model.LoginRequest) (user model.User, err error) {
	err = login.Validate()
	if err != nil {
		return
	}

	err = s.db.Tx(context.Background(), func(tx *db.Tx) error {
		user, err = s.userDao.FindByEmail(tx, login.Email)
		if err != nil {
			if db.IsNoDataFound(err) {
				return error_util.NewUnauthorized("invalid email or password")
			}
			return err
		}

		if !user.Active {
			return error_util.NewUnauthorized("user is not active")
		}

		uc, err := s.userCredentialDao.Get(tx, user.ID)
		if err != nil {
			if db.IsNoDataFound(err) {
				return error_util.NewUnauthorized("invalid email or password")
			}
			return err
		}

		if !uc.Activated {
			return error_util.NewUnauthorized("user is not activated")
		} else if !uc.ExpiresAt.IsZero() && uc.ExpiresAt.Before(time.Now()) {
			return error_util.NewUnauthorized("password expired")
		} else if uc.Locked {
			return error_util.NewUnauthorized("account is locked")
		}

		passwordSha := crypto.SHA256([]byte(login.Password))
		matched, err := crypto.CheckPassword(passwordSha, uc.PasswordHash)

		if err != nil {
			return err
		}

		if !matched {
			if uc.InvalidAttempts >= 3 {
				err := s.userCredentialDao.IncrementInvalidAttempts(tx, user.ID, true)
				if err != nil {
					return err
				}
			} else {
				err := s.userCredentialDao.IncrementInvalidAttempts(tx, user.ID, false)
				if err != nil {
					return err
				}
			}
			return error_util.NewUnauthorized("invalid email or password")
		}

		if uc.InvalidAttempts > 0 {
			err := s.userCredentialDao.ResetInvalidAttempts(tx, user.ID)
			if err != nil {
				log.Error().Err(err).Msg("failed resetting invalid attempts")
			}
		}
		return nil
	})

	if err != nil {
		return
	}
	return
}

func (s *Service) ChangePassword(ctx context.Context, data model.ChangePasswordRequest) error {

	id := auth.UserIdFromContext(ctx)

	err := s.db.Tx(context.Background(), func(tx *db.Tx) error {
		uc, err := s.userCredentialDao.Get(tx, id)

		if err != nil {
			if db.IsNoDataFound(err) {
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
				err := s.userCredentialDao.IncrementInvalidAttempts(tx, id, true)
				if err != nil {
					return err
				}
			} else {
				err := s.userCredentialDao.IncrementInvalidAttempts(tx, id, false)
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

		err = s.userCredentialDao.ChangePassword(tx, uc.ID, newPasswordHash)
		return err
	})

	return err
}

func (s *Service) InitiatePasswordReset(email string) error {

	var user model.User
	var err error
	resetToken := uuid.New().String()
	resetTokenSha := fmt.Sprintf("%x", sha256.Sum256([]byte(resetToken)))
	expiresAt := time.Now().Add(20 * time.Minute)

	err = s.db.Tx(context.Background(), func(tx *db.Tx) error {
		user, err = s.userDao.FindByEmail(tx, email)
		if err != nil {
			if db.IsNoDataFound(err) {
				return error_util.NewBadRequest("user not found")
			}
			return err
		}
		_, err = s.userCredentialDao.Get(tx, user.ID)

		if err != nil {
			if db.IsNoDataFound(err) {
				cred := model.UserCredential{
					ID:                user.ID,
					ResetKey:          resetTokenSha,
					ResetKeyExpiresAt: expiresAt,
				}
				err = s.userCredentialDao.Insert(tx, &cred)
				return err
			} else {
				return err
			}
		} else {
			err = s.userCredentialDao.UpdateResetKey(tx, user.ID, resetTokenSha, expiresAt)
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

func (s *Service) ResetPassword(passwordResetRequest model.ResetPasswordRequest) error {
	resetTokenSha := fmt.Sprintf("%x", sha256.Sum256([]byte(passwordResetRequest.ResetToken)))

	err := s.db.Tx(context.Background(), func(tx *db.Tx) error {
		uc, err := s.userCredentialDao.FindByResetKey(tx, resetTokenSha)

		if err != nil {
			if db.IsNoDataFound(err) {
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

		err = s.userCredentialDao.ResetPassword(tx, uc.ID, passwordHash)
		return err
	})

	return err
}

func (s *Service) SignUp(signUpRequest model.SignUpRequest) (*model.User, error) {

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
	err = s.db.Tx(context.Background(), func(tx *db.Tx) error {
		err = s.checkForDuplicate(tx, signUpRequest.Email, "email", s.userDao.ExistsByEmail)
		if err != nil {
			return err
		}

		err = s.userDao.Insert(tx, &newUser)
		if err != nil {
			return error_util.NewInternal(err, "failed to insert user")
		}

		cred := model.UserCredential{
			ID:                     newUser.ID,
			PasswordHash:           passwordHash,
			ActivationKey:          activationTokenHash,
			ActivationKeyExpiresAt: time.Now().Add(time.Second * 1200),
		}

		err = s.userCredentialDao.Insert(tx, &cred)
		if err != nil {
			return error_util.NewInternal(err, "Internal error - unable to insert user_credential")
		}
		return nil
	})

	if err != nil {
		return nil, error_util.ToError(err, "failed to complete sign-up transaction")
	}

	err = s.notifier.NotifyActivation(newUser, activationToken)

	if err != nil {
		log.Error().Err(err).Msg("failed to send account activation email")
		return nil, error_util.NewInternal(err, "failed to send account activation email")
	}

	log.Debug().Interface("user", newUser).Msg("successfully signed up user")
	return &newUser, nil
}

func (s *Service) GetProfile(ctx context.Context) (user model.User, err error) {
	id := auth.UserIdFromContext(ctx)

	err = s.db.Tx(context.Background(), func(tx *db.Tx) error {
		user, err = s.userDao.Find(tx, id)
		return err
	})
	return user, err
}

func (s *Service) UpdateProfile(user *model.User) error {
	return nil
}

func (s *Service) checkForDuplicate(tx *db.Tx, input string, by string, fn func(*db.Tx, string) (bool, error)) error {
	exists, err := fn(tx, input)
	if err != nil {
		return error_util.NewInternal(err, "Error while checking for duplicate email")
	} else if exists {
		log.Info().Str(by, input).Msgf("found user with same %s", by)
		fieldErrors := []error_util.FieldError{{Field: "email", Message: "user already exists"}}
		return error_util.WithFieldErrors(fieldErrors)
	}
	return nil
}
