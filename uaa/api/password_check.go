package main

import (
	"strconv"
	"time"


	log "github.com/sirupsen/logrus"
	"mmrath.com/gobase/common/crypto"
	"mmrath.com/gobase/model"
)

type Service struct {
	db *model.DB
}

func (s *Service) CheckPassword(username string, password string) (string, error) {
	user, err := s.Login(&model.LoginRequest{Email: username, Password: password})
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(user.ID, 10), nil
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
				return errors2.NewUnauthorized("invalid email or password")
			}
			return err
		}

		if !user.Active {
			return errors2.NewUnauthorized("user is not active")
		}

		uc, err := tx.UserCredentialDao().Get(user.ID)
		if err != nil {
			if model.IsNoDataFound(err) {
				return errors2.NewUnauthorized("invalid email or password")
			}
			return err
		}

		if !uc.Activated {
			return errors2.NewUnauthorized("user is not activated")
		} else if !uc.ExpiresAt.IsZero() && uc.ExpiresAt.Before(time.Now()) {
			return errors2.NewUnauthorized("password expired")
		} else if uc.Locked {
			return errors2.NewUnauthorized("account is locked")
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
			return errors2.NewUnauthorized("invalid email or password")
		}

		if uc.InvalidAttempts > 0 {
			err := tx.UserCredentialDao().ResetInvalidAttempts(user.ID)
			if err != nil {
				log.Errorf("Error while resetting invalid attempts %+v", err)
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return user, nil
}
