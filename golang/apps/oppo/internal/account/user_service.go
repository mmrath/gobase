package account

import (
	"context"

	"github.com/mmrath/gobase/golang/pkg/db"
	"github.com/mmrath/gobase/golang/pkg/model"
)

type UserService interface {
	FindUserByID(ctx context.Context, id int64) (model.User, error)
	CreateUser(ctx context.Context, role *model.CreateUserRequest) (model.User, error)
	UpdateUser(ctx context.Context, role *model.User) error
}

type userService struct {
	db      *db.DB
	userDao model.UserDao
}

func (u userService) FindUserByID(ctx context.Context, id int64) (model.User, error) {
	var user model.User
	err := u.db.RunInTx(ctx, func(tx *db.Tx) error {
		var err error
		user, err = u.userDao.Find(tx, id)
		return err
	})
	return user, err
}

func (u userService) CreateUser(ctx context.Context, user *model.CreateUserRequest) (model.User, error) {
	panic("implement me")
}

func (u userService) UpdateUser(ctx context.Context, user *model.User) (err error) {
	panic("implement me")
}

func NewUserService(database *db.DB) UserService {
	return &userService{
		db:      database,
		userDao: model.NewUserDao(),
	}
}
