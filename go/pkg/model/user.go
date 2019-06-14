package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type User struct {
	tableName struct{} `sql:"user_account"`

	AuditDetails

	ID          int64  `json:"id,omitempty"	`
	FirstName   string `json:"firstName,omitempty"`
	LastName    string `json:"lastName,omitempty"`
	Email       string `json:"email,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
	Active      bool   `json:"active,omitempty"`
	Roles       []Role `json:"roles" sql:"-"`
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
	tx *Tx
}

type UserDao interface {
	Get(id int64) (*User, error)
	Insert(user *User) error
	Update(user *User) error
	GetByEmail(email string) (*User, error)
	ExistsByEmail(email string) (bool, error)
}

func newUserDao(tx *Tx) UserDao {
	return &userDao{
		tx: tx,
	}
}

func (dao *userDao) Get(id int64) (*User, error) {
	user := &User{ID: id}
	err := dao.tx.Select(user)
	return user, err
}

func (dao *userDao) Insert(user *User) error {
	err := dao.tx.Insert(user)
	return err
}

func (dao *userDao) Update(user *User) error {
	err := dao.tx.Update(user)
	return err
}

func (dao *userDao) GetByEmail(email string) (*User, error) {
	user := new(User)
	err := dao.tx.Model(user).Where("email = ?", email).Select()
	if err != nil {
		return nil, err
	} else {
		return user, nil
	}
}

func (dao *userDao) ExistsByEmail(email string) (bool, error) {
	user := new(User)
	return dao.tx.Model(user).Where("email = ?", email).Exists()
}
