package model

import (
	"time"
)

type UserCredential struct {
	tableName struct{} `sql:"user_credential"`

	ID                     int64     `json:"id,omitempty"`
	PasswordHash           string    `json:"-"`
	ExpiresAt              time.Time `json:"expiresAt,omitempty"`
	InvalidAttempts        uint16    `json:"invalidAttempts,omitempty"`
	Locked                 bool      `json:"locked,omitempty"`
	ActivationKey          string    `json:"activationKey,omitempty"`
	ActivationKeyExpiresAt time.Time `json:"activationKeyExpiresAt,omitempty"`
	Activated              bool      `json:"activated,omitempty"`
	ResetKey               string    `json:"resetKey,omitempty"`
	ResetKeyExpiresAt      time.Time `json:"resetKeyExpiresAt,omitempty"`
	ResetAt                time.Time `json:"resetAt,omitempty"`
	UpdatedAt              time.Time `json:"updatedAt,omitempty"`
	Version                uint16    `json:"version,omitempty"`
}

type userCredentialDao struct {
	tx *Tx
}

type UserCredentialDao interface {
	GetByActivationKey(key string) (*UserCredential, error)
	Activate(id int64) error
	Get(id int64) (*UserCredential, error)
	Insert(credential *UserCredential) error
	IncrementInvalidAttempts(id int64, lock bool) error
	UpdateResetKey(id int64, resetKey string, expiresAt time.Time) error
	FindByResetKey(key string) (*UserCredential, error)
	ResetPassword(id int64, newPassword string) error
	ChangePassword(id int64, newPassword string) error
	ResetInvalidAttempts(id int64) error
}

func newUserCredentialDao(tx *Tx) UserCredentialDao {
	return &userCredentialDao{tx}
}

func (dao *userCredentialDao) GetByActivationKey(key string) (*UserCredential, error) {
	userCred := new(UserCredential)
	err := dao.tx.Model(userCred).Where("activation_key = ?", key).Select()
	if err != nil {
		return nil, err
	} else {
		return userCred, nil
	}
}

func (dao *userCredentialDao) Activate(id int64) error {
	userCred := UserCredential{ID: id, Activated: true}
	_, err := dao.tx.Model(&userCred).Column("activated").WherePK().Update()
	return err
}

func (dao *userCredentialDao) Get(id int64) (*UserCredential, error) {
	userCred := &UserCredential{ID: id}
	err := dao.tx.Select(userCred)
	return userCred, err
}

func (dao *userCredentialDao) Insert(credential *UserCredential) error {
	err := dao.tx.Insert(credential)
	return err
}

func (dao *userCredentialDao) IncrementInvalidAttempts(id int64, lock bool) error {
	userCred := UserCredential{ID: id}
	_, err := dao.tx.Model(&userCred).
		Set("invalid_attempts = invalid_attempts + 1").
		Set("locked = ?", lock).
		WherePK().
		Update()
	return err
}

func (dao *userCredentialDao) UpdateResetKey(id int64, resetKey string, expiresAt time.Time) error {
	userCred := UserCredential{ID: id, ResetKey: resetKey, ResetKeyExpiresAt: expiresAt}
	_, err := dao.tx.Model(&userCred).
		Column("reset_key", "reset_key_expires_at").
		WherePK().
		Update()
	return err
}

func (dao *userCredentialDao) FindByResetKey(key string) (*UserCredential, error) {
	userCred := new(UserCredential)
	err := dao.tx.Model(userCred).
		Where("reset_key = ?", key).
		Select()
	return userCred, err
}

func (dao *userCredentialDao) ResetPassword(id int64, newPassword string) error {
	userCred := UserCredential{ID: id, PasswordHash: newPassword}

	_, err := dao.tx.Model(&userCred).
		Set("password_hash = ?", newPassword).
		Set("expires_at = ?", time.Now().AddDate(1, 0, 0)).
		Set("activated = true").
		Set("invalid_attempts = 0").
		Set("locked = false").
		Set("reset_key = null").
		Set("reset_key_expires_at = null").
		Set("reset_at = ?", time.Now()).
		WherePK().
		Update()
	return err
}

func (dao *userCredentialDao) ChangePassword(id int64, newPassword string) error {
	userCred := UserCredential{ID: id, PasswordHash: newPassword}

	_, err := dao.tx.Model(&userCred).
		Set("password_hash = ?", newPassword).
		Set("expires_at = ?", time.Now().AddDate(1, 0, 0)).
		Set("invalid_attempts = 0").
		Set("locked = false").
		WherePK().
		Update()
	return err
}

func (dao *userCredentialDao) ResetInvalidAttempts(id int64) error {
	userCred := UserCredential{ID: id}

	_, err := dao.tx.Model(&userCred).
		Set("invalid_attempts = 0").
		WherePK().
		Update()
	return err
}
