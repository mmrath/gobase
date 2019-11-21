package model

import (
	"context"
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
}

type UserCredentialDao interface {
	GetByActivationKey(ctx context.Context, key string) (*UserCredential, error)
	Activate(ctx context.Context, id int64) error
	Get(ctx context.Context, id int64) (*UserCredential, error)
	Insert(ctx context.Context, credential *UserCredential) error
	IncrementInvalidAttempts(ctx context.Context, id int64, lock bool) error
	UpdateResetKey(ctx context.Context, id int64, resetKey string, expiresAt time.Time) error
	FindByResetKey(ctx context.Context, key string) (*UserCredential, error)
	ResetPassword(ctx context.Context, id int64, newPassword string) error
	ChangePassword(ctx context.Context, id int64, newPassword string) error
	ResetInvalidAttempts(ctx context.Context, id int64) error
}

func newUserCredentialDao(tx *Tx) UserCredentialDao {
	return &userCredentialDao{}
}

func (dao *userCredentialDao) GetByActivationKey(ctx context.Context, key string) (*UserCredential, error) {
	userCred := new(UserCredential)
	err := TxFromContext(ctx).Model(userCred).Where("activation_key = ?", key).Select()
	if err != nil {
		return nil, err
	} else {
		return userCred, nil
	}
}

func (dao *userCredentialDao) Activate(ctx context.Context, id int64) error {
	userCred := UserCredential{ID: id, Activated: true}
	_, err := TxFromContext(ctx).Model(&userCred).Column("activated").WherePK().Update()
	return err
}

func (dao *userCredentialDao) Get(ctx context.Context, id int64) (*UserCredential, error) {
	userCred := &UserCredential{ID: id}
	err := TxFromContext(ctx).Select(userCred)
	return userCred, err
}

func (dao *userCredentialDao) Insert(ctx context.Context, credential *UserCredential) error {
	err := TxFromContext(ctx).Insert(credential)
	return err
}

func (dao *userCredentialDao) IncrementInvalidAttempts(ctx context.Context, id int64, lock bool) error {
	userCred := UserCredential{ID: id}
	_, err := TxFromContext(ctx).Model(&userCred).
		Set("invalid_attempts = invalid_attempts + 1").
		Set("locked = ?", lock).
		WherePK().
		Update()
	return err
}

func (dao *userCredentialDao) UpdateResetKey(ctx context.Context, id int64, resetKey string, expiresAt time.Time) error {
	userCred := UserCredential{ID: id, ResetKey: resetKey, ResetKeyExpiresAt: expiresAt}
	_, err := TxFromContext(ctx).Model(&userCred).
		Column("reset_key", "reset_key_expires_at").
		WherePK().
		Update()
	return err
}

func (dao *userCredentialDao) FindByResetKey(ctx context.Context, key string) (*UserCredential, error) {
	userCred := new(UserCredential)
	err := TxFromContext(ctx).Model(userCred).
		Where("reset_key = ?", key).
		Select()
	return userCred, err
}

func (dao *userCredentialDao) ResetPassword(ctx context.Context, id int64, newPassword string) error {
	userCred := UserCredential{ID: id, PasswordHash: newPassword}

	_, err := TxFromContext(ctx).Model(&userCred).
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

func (dao *userCredentialDao) ChangePassword(ctx context.Context, id int64, newPassword string) error {
	userCred := UserCredential{ID: id, PasswordHash: newPassword}

	_, err := TxFromContext(ctx).Model(&userCred).
		Set("password_hash = ?", newPassword).
		Set("expires_at = ?", time.Now().AddDate(1, 0, 0)).
		Set("invalid_attempts = 0").
		Set("locked = false").
		WherePK().
		Update()
	return err
}

func (dao *userCredentialDao) ResetInvalidAttempts(ctx context.Context, id int64) error {
	userCred := UserCredential{ID: id}

	_, err := TxFromContext(ctx).Model(&userCred).
		Set("invalid_attempts = 0").
		WherePK().
		Update()
	return err
}
