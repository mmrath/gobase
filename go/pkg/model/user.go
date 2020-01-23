package model

import (
	"github.com/mmrath/gobase/go/pkg/db"
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
	return "user_account"
}

func (u User) GetName() string {
	return u.FirstName
}

func (u User) GetEmail() string {
	return u.Email
}

func (u User) GetID() int64 {
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
	Email    string `json:"email,omitempty" validate:"required,email"`
	Password string `json:"password,omitempty" validate:"required,min=6,max=20"`
}

type CreateUserRequest struct {
	FirstName   string  `json:"firstName,omitempty" validate:"required,aplha"`
	LastName    string  `json:"lastName,omitempty" validate:"required,alpha"`
	Email       string  `json:"email,omitempty" validate:"required,email"`
	PhoneNumber string  `json:"phoneNumber,omitempty" validate:"required"`
	Active      bool    `json:"active,omitempty"`
	Roles       []int32 `json:"roles,omitempty"`
}

type RegisterAccountRequest struct {
	Password    string `json:"password" validate:"required,min=6,max=20" valid:"length(6|20)"`
	FirstName   string `json:"firstName" validate:"required,alpha,min=2,max=32"  valid:"alpha,length(2|32)"`
	LastName    string `json:"lastName" validate:"required,alpha,max=32" valid:"alpha,length(1|32)"`
	Email       string `json:"email" validate:"required,email,min=6,max=32" valid:"email,length(6|32)"`
	PhoneNumber string `json:"phoneNumber"`
}

type userDao struct {
}

func NewUserDao() UserDao {
	return &userDao{}
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
