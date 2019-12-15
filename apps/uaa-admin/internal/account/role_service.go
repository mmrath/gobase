package account

import (
	"context"
	"fmt"
	"github.com/mmrath/gobase/pkg/error_util"
	"github.com/mmrath/gobase/pkg/model"
	"github.com/mmrath/gobase/pkg/db"
)

type RoleService interface {
	Find(ctx context.Context, id int32) (*model.RoleAndPermission, error)
	Create(ctx context.Context, role *model.RoleAndPermission) (err error)
	Update(ctx context.Context, role *model.RoleAndPermission) (err error)
}

type roleService struct {
	db      *db.DB
	roleDao model.RoleDao
}

func (s *roleService) Find(ctx context.Context, id int32) (role *model.RoleAndPermission, err error) {
	err = s.db.Tx(ctx, func(tx *db.Tx) error {
		role.Role, err = s.roleDao.Find(tx, id)
		return err
	})
	return role, err
}

func (s *roleService) Create(ctx context.Context, roleAndPermission *model.RoleAndPermission) (err error) {
	err = s.db.Tx(ctx, func(tx *db.Tx) error {
		exists, err := s.roleDao.ExistsByName(tx, roleAndPermission.Role.Name)
		if err != nil {
			return error_util.NewInternal(err, "error while checking if roleAndPermission exists")
		}
		if exists {
			return error_util.NewBadRequest(
				fmt.Sprintf("role with name %s already exists", roleAndPermission.Role.Name))
		}
		err = s.roleDao.Create(tx, &roleAndPermission.Role, roleAndPermission.Permissions)
		return err
	})
	return err
}

func (s *roleService) Update(ctx context.Context, roleAndPermission *model.RoleAndPermission) (err error) {
	err = s.db.Tx(ctx, func(tx *db.Tx) error {
		return s.roleDao.Update(tx, &roleAndPermission.Role, roleAndPermission.Permissions)
	})
	return err
}

func NewRoleService(db *db.DB) RoleService {
	return &roleService{
		db:      db,
	}
}
