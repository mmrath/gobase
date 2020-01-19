package account

import (
	"context"
	"crypto/sha256"
	"fmt"
	"github.com/mmrath/gobase/go/pkg/db"
	"github.com/mmrath/gobase/go/pkg/errutil"
	"github.com/mmrath/gobase/go/pkg/validate"
	"strings"
	"time"

	"github.com/mmrath/gobase/go/pkg/auth"
	"github.com/mmrath/gobase/go/pkg/crypto"
	"github.com/mmrath/gobase/go/pkg/model"
	"github.com/rs/zerolog/log"

	"github.com/google/uuid"
)

type Service struct {
	notifier          Notifier
	db                *db.DB
	userCredentialDao model.UserCredentialDao
	userDao           model.UserDao
}

func NewService(notifier Notifier, db *db.DB) *Service {
	return &Service{
		notifier:          notifier,
		db:                db,
		userCredentialDao: model.NewUserCredentialDao(),
		userDao:           model.NewUserDao(),
	}
}

func (s *Service) Activate(token string) error {
	err := validate.Field(token, "required,min=4,max=128")
	if err != nil {
		return err
	}

	tokenHash := fmt.Sprintf("%x", sha256.Sum256([]byte(token)))

	err = s.db.RunInTx(context.Background(), func(tx *db.Tx) error {

		uc, err := s.userCredentialDao.GetByActivationKey(tx, tokenHash)
		if err != nil {
			if db.IsNoDataFound(err) {
				return errutil.NewBadRequest("invalid activation token")
			}
			return errutil.Wrap(err, "error while trying get activation")
		}
		if !uc.Activated {
			if uc.ActivationKeyExpiresAt.Before(time.Now()) {
				return errutil.New("activation token is expired, sign up again")
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
	err = validate.Struct(login)
	if err != nil {
		return user, errutil.Wrap(err, "failed validation")
	}

	invalidCredentialMsg := "invalid email or password"
	err = s.db.RunInTx(context.Background(), func(tx *db.Tx) error {
		user, err = s.userDao.FindByEmail(tx, login.Email)
		if err != nil {
			if db.IsNoDataFound(err) {
				return errutil.NewUnauthorized(invalidCredentialMsg)
			}
			return err
		}

		if !user.Active {
			return errutil.NewUnauthorized("user is not active")
		}

		uc, err := s.userCredentialDao.Get(tx, user.ID)
		if err != nil {
			if db.IsNoDataFound(err) {
				return errutil.NewUnauthorized(invalidCredentialMsg)
			}
			return err
		}

		if !uc.Activated {
			return errutil.NewUnauthorized("user is not activated")
		} else if !uc.ExpiresAt.IsZero() && uc.ExpiresAt.Before(time.Now()) {
			return errutil.NewUnauthorized("password expired")
		} else if uc.Locked {
			return errutil.NewUnauthorized("account is locked")
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
					return errutil.Wrapf(err, "failed to lock and increment invalid attempts")
				}
			} else {
				err := s.userCredentialDao.IncrementInvalidAttempts(tx, user.ID, false)
				if err != nil {
					return errutil.Wrapf(err, "failed to increment invalid attempts")
				}
			}
			return errutil.NewUnauthorized(invalidCredentialMsg)
		}

		if uc.InvalidAttempts > 0 {
			err := s.userCredentialDao.ResetInvalidAttempts(tx, user.ID)
			if err != nil {
				// allow user to login
				log.Error().Err(err).Msg("failed resetting invalid attempts")
			}
		}
		return nil
	})
	if err != nil {
		log.Error().Err(err).Send()
	}
	return
}

func (s *Service) ChangePassword(ctx context.Context, data model.ChangePasswordRequest) error {

	id := auth.UserIdFromContext(ctx)

	err := s.db.RunInTx(context.Background(), func(tx *db.Tx) error {

		uc, err := s.userCredentialDao.Get(tx, id)
		log.Info().Interface("user_credential", uc).Msg("credential found")

		if err != nil {
			if db.IsNoDataFound(err) {
				return errutil.NewBadRequest("password cannot be changed as user does not exist")
			}
			return errutil.Wrapf(err, "failed to retrieve user credential for %d", id)
		}

		currentPasswordSha := crypto.SHA256([]byte(data.CurrentPassword))
		matched, err := crypto.CheckPassword(currentPasswordSha, uc.PasswordHash)

		if err != nil {
			return errutil.Wrapf(err, "failed to validate password")
		}

		if !matched {
			if uc.InvalidAttempts >= 3 {
				err := s.userCredentialDao.IncrementInvalidAttempts(tx, id, true)
				if err != nil {
					return errutil.Wrapf(err, "failed to lock user")
				}
			} else {
				err := s.userCredentialDao.IncrementInvalidAttempts(tx, id, false)
				if err != nil {
					return errutil.Wrapf(err, "failed to update invalid attempts")
				}
			}
			return errutil.NewUnauthorized("invalid current password")
		}

		newPasswordHash, err := crypto.HashPassword(crypto.SHA256([]byte(data.NewPassword)))
		if err != nil {
			return errutil.Wrap(err, "failed to hash password")
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

	err = s.db.RunInTx(context.Background(), func(tx *db.Tx) error {
		user, err = s.userDao.FindByEmail(tx, email)
		if err != nil {
			if db.IsNoDataFound(err) {
				return errutil.NewBadRequest("user not found")
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
		return errutil.Wrap(err, "Failed to send password reset email")
	}

	log.Info().Int64("id", user.ID).Msg("successfully sent password reset email")
	return nil
}

func (s *Service) ResetPassword(passwordResetRequest model.ResetPasswordRequest) error {
	resetTokenSha := fmt.Sprintf("%x", sha256.Sum256([]byte(passwordResetRequest.ResetToken)))

	err := s.db.RunInTx(context.Background(), func(tx *db.Tx) error {
		uc, err := s.userCredentialDao.FindByResetKey(tx, resetTokenSha)

		if err != nil {
			if db.IsNoDataFound(err) {
				return errutil.NewBadRequest("reset key is invalid")
			}
			return err
		}

		if uc.ResetKeyExpiresAt.Before(time.Now()) {
			return errutil.NewBadRequest("reset key is expired")
		}

		passwordHash, err := crypto.HashPassword(crypto.SHA256([]byte(passwordResetRequest.NewPassword)))
		if err != nil {
			return errutil.Wrap(err, "failed to hash password")
		}

		err = s.userCredentialDao.ResetPassword(tx, uc.ID, passwordHash)
		return err
	})

	return err
}

func (s *Service) Register(signUpRequest model.RegisterAccountRequest) (*model.User, error) {

	log.Debug().Interface("email", signUpRequest.Email).Msg("signing up user")
	err := validate.Struct(signUpRequest)

	if err != nil {
		return nil, err
	}

	newUser := model.User{
		AuditDetails: model.AuditDetails{UpdatedBy: "SIGNUP"},
		FirstName:    signUpRequest.FirstName,
		LastName:     signUpRequest.LastName,
		Email:        strings.ToLower(signUpRequest.Email),
		Active:       true,
	}

	newUser.UpdatedBy = "SIGNUP"

	passwordHash, err := crypto.HashPassword(crypto.SHA256([]byte(signUpRequest.Password)))
	if err != nil {
		return nil, errutil.Wrap(err, "failed to hash password")
	}

	activationToken := uuid.New().String()
	activationTokenHash := fmt.Sprintf("%x", sha256.Sum256([]byte(activationToken)))
	err = s.db.RunInTx(context.Background(), func(tx *db.Tx) error {
		err = s.checkForDuplicate(tx, signUpRequest.Email, "email", s.userDao.ExistsByEmail)
		if err != nil {
			return err
		}

		err = s.userDao.Insert(tx, &newUser)
		if err != nil {
			return errutil.Wrap(err, "failed to insert user")
		}

		cred := model.UserCredential{
			ID:                     newUser.ID,
			PasswordHash:           passwordHash,
			ActivationKey:          activationTokenHash,
			ActivationKeyExpiresAt: time.Now().Add(time.Second * 1200),
		}

		err = s.userCredentialDao.Insert(tx, &cred)
		if err != nil {
			return errutil.Wrap(err, "Internal error - unable to insert user_credential")
		}
		return nil
	})

	if err != nil {
		return nil, errutil.Wrap(err, "failed to complete sign-up transaction")
	}

	err = s.notifier.NotifyActivation(newUser, activationToken)

	if err != nil {
		return nil, errutil.Wrap(err, "failed to send account activation email")
	}

	log.Debug().Interface("user", newUser).Msg("successfully signed up user")
	return &newUser, nil
}

func (s *Service) GetProfile(ctx context.Context) (user model.User, err error) {
	id := auth.UserIdFromContext(ctx)

	err = s.db.RunInTx(context.Background(), func(tx *db.Tx) error {
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
		return errutil.Wrap(err, "defaultError while checking for duplicate email")
	} else if exists {
		log.Info().Str(by, input).Msgf("found user with same %s", by)
		return errutil.NewFieldError("email", "email already registered")
	}
	return nil
}
