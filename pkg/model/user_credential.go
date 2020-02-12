package model

import (
	"time"

	"github.com/jinzhu/gorm"

	"github.com/mmrath/gobase/pkg/db"
)

type UserCredential struct {
	ID                     int64     `json:"id,omitempty"`
	PasswordHash           string    `json:"-" sql:"default:null"`
	ExpiresAt              time.Time `json:"expiresAt,omitempty"`
	InvalidAttempts        uint16    `json:"invalidAttempts,omitempty"`
	Locked                 bool      `json:"locked,omitempty"`
	ActivationKey          string    `json:"activationKey,omitempty" sql:"default:null"`
	ActivationKeyExpiresAt time.Time `json:"activationKeyExpiresAt,omitempty"`
	Activated              bool      `json:"activated,omitempty"`
	ResetKey               string    `json:"resetKey,omitempty" sql:"default:null"`
	ResetKeyExpiresAt      time.Time `json:"resetKeyExpiresAt,omitempty"`
	ResetAt                time.Time `json:"resetAt,omitempty"`
	UpdatedAt              time.Time `json:"updatedAt,omitempty"`
	Version                uint16    `json:"version,omitempty"`
}

type userCredentialDao struct {
}

type UserCredentialDao interface {
	GetByActivationKey(tx *db.Tx, key string) (UserCredential, error)
	Activate(tx *db.Tx, id int64) error
	Get(tx *db.Tx, id int64) (UserCredential, error)
	Insert(tx *db.Tx, credential *UserCredential) error
	IncrementInvalidAttempts(tx *db.Tx, id int64, lock bool) error
	UpdateResetKey(tx *db.Tx, id int64, resetKey string, expiresAt time.Time) error
	FindByResetKey(tx *db.Tx, key string) (UserCredential, error)
	ResetPassword(tx *db.Tx, id int64, newPassword string) error
	ChangePassword(tx *db.Tx, id int64, newPassword string) error
	ResetInvalidAttempts(tx *db.Tx, id int64) error
}

func NewUserCredentialDao() UserCredentialDao {
	return &userCredentialDao{}
}

func (dao *userCredentialDao) GetByActivationKey(tx *db.Tx, key string) (UserCredential, error) {
	userCred := UserCredential{}
	err := tx.Where("activation_key = ?", key).First(&userCred).Error
	return userCred, err
}

func (dao *userCredentialDao) Activate(tx *db.Tx, id int64) error {
	userCred := UserCredential{ID: id, Activated: true}
	return tx.Model(&userCred).UpdateColumns(map[string]interface{}{
		"activated": true,
	}).Error
}

func (dao *userCredentialDao) Get(tx *db.Tx, id int64) (UserCredential, error) {
	userCred := UserCredential{ID: id}
	err := tx.First(&userCred).Error
	return userCred, err
}

func (dao *userCredentialDao) Insert(tx *db.Tx, credential *UserCredential) error {
	err := tx.Create(credential).Error
	return err
}

func (dao *userCredentialDao) IncrementInvalidAttempts(tx *db.Tx, id int64, lock bool) error {
	userCred := UserCredential{ID: id}
	err := tx.Model(&userCred).
		UpdateColumns(map[string]interface{}{
			"invalid_attempts": gorm.Expr("invalid_attempts + ?", 1),
			"locked":           lock,
		}).Error
	return err
}

func (dao *userCredentialDao) UpdateResetKey(tx *db.Tx, id int64, resetKey string, expiresAt time.Time) error {
	userCred := UserCredential{ID: id, ResetKey: resetKey, ResetKeyExpiresAt: expiresAt}
	err := tx.Model(&userCred).Updates(userCred).Error
	return err
}

func (dao *userCredentialDao) FindByResetKey(tx *db.Tx, key string) (UserCredential, error) {
	userCred := UserCredential{}
	err := tx.Where("reset_key = ?", key).
		First(&userCred).Error
	return userCred, err
}

func (dao *userCredentialDao) ResetPassword(tx *db.Tx, id int64, newPassword string) error {
	userCred := UserCredential{ID: id}
	err := tx.Model(&userCred).
		Updates(map[string]interface{}{
			"password_hash":        newPassword,
			"expires_at":           time.Now().AddDate(1, 0, 0),
			"activated":            true,
			"invalid_attempts":     0,
			"locked":               false,
			"reset_key":            nil,
			"reset_key_expires_at": nil,
			"reset_at":             time.Now(),
		}).Error
	return err
}

func (dao *userCredentialDao) ChangePassword(tx *db.Tx, id int64, newPassword string) error {
	userCred := UserCredential{ID: id}
	err := tx.Model(&userCred).
		Updates(map[string]interface{}{
			"password_hash":    newPassword,
			"expires_at":       time.Now().AddDate(1, 0, 0),
			"invalid_attempts": 0,
			"locked":           false,
		}).Error
	return err
}

func (dao *userCredentialDao) ResetInvalidAttempts(tx *db.Tx, id int64) error {
	userCred := UserCredential{ID: id}
	err := tx.Model(&userCred).
		UpdateColumn("invalid_attempts", 0).Error
	return err
}
