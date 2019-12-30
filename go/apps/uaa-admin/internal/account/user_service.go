package account

import (
	"context"
	"github.com/mmrath/gobase/go/pkg/model"
)

type UserService interface {
	Find(ctx context.Context, id int32) (*model.User, error)
	Create(ctx context.Context, role *model.User) (err error)
	Update(ctx context.Context, role *model.User) (err error)
}
