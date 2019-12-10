package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/mmrath/gobase/pkg/db"
)

type User struct {
	AuditDetails

	ID          int64  `json:"id,omitempty"`
	FirstName   string `json:"firstName,omitempty"`
	LastName    string `json:"lastName,omitempty"`
	Email       string `json:"email,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
	Active      bool   `json:"active,omitempty"`
}

func (User) TableName() string {
	return "app_user"
}

func (u *User) GetName() string {
	return u.FirstName
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetId() int64 {
	return u.ID
}

type ChangePasswordRequest struct {
	CurrentPassword string `json:"currentPassword,omitempty" validate:"required"`
	NewPassword     string `json:"newPassword,omitempty" validate:"required"`
}

type ResetPasswordRequest struct {
	ResetToken  string `json:"resetToken,omitempty" validate:"required"`
	NewPassword string `json:"newPassword,omitempty" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type CreateUserRequest struct {
	FirstName   string  `json:"firstName,omitempty" validate:"required,aplha"`
	LastName    string  `json:"lastName,omitempty" validate:"required,alpha"`
	Email       string  `json:"email,omitempty" validate:"required,email"`
	PhoneNumber string  `json:"phoneNumber,omitempty" validate:"required"`
	Active      bool    `json:"active,omitempty"`
	Roles       []int32 `json:"roles,omitempty"`
}

func (login *LoginRequest) Validate() error {
	err := validation.ValidateStruct(login,
		validation.Field(&login.Email, validation.Required, validation.Length(6, 32)),
		validation.Field(&login.Password, validation.Required, validation.Length(6, 32)),
	)
	if err != nil {
		return err
	}
	return nil
}

type SignUpRequest struct {
	Password    string `json:"password"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
}

func (s *SignUpRequest) Validate() error {
	err := validation.ValidateStruct(s,
		validation.Field(&s.FirstName, validation.Required, validation.Length(2, 32)),
		validation.Field(&s.LastName, validation.Required, validation.Length(1, 32)),
		validation.Field(&s.Email, validation.Required, validation.Length(6, 32), is.Email),
		validation.Field(&s.Password, validation.Required, validation.Length(6, 32)),
		validation.Field(&s.PhoneNumber, is.E164), // International phone
	)
	if err != nil {
		return err
	}
	return nil
}

type userDao struct {
}

type UserDao interface {
	Find(tx *db.Tx, id int64) (User, error)
	Insert(tx *db.Tx, user *User) error
	Update(tx *db.Tx, user *User) error
	FindByEmail(tx *db.Tx, email string) (User, error)
	ExistsByEmail(tx *db.Tx, email string) (bool, error)
}

func (dao *userDao) Find(tx *db.Tx, id int64) (User, error) {
	user := User{}
	err := tx.First(&user, id).Error
	return user, err
}

func (dao *userDao) Insert(tx *db.Tx, user *User) error {
	err := tx.Model(user).Create(user).Error
	return err
}

func (dao *userDao) Update(tx *db.Tx, user *User) error {
	return tx.Save(user).Error
}

func (dao *userDao) FindByEmail(tx *db.Tx, email string) (User, error) {
	user := User{}
	err := tx.Where("email = ?", email).Find(&user).Error
	return user, err
}

func (dao *userDao) ExistsByEmail(tx *db.Tx, email string) (bool, error) {
	count := 0
	err := tx.Model(&User{}).Where("email = ?", email).Count(&count).Error
	return count != 0, err
}
