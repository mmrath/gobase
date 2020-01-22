package account

import (
	"context"
	"fmt"

	"github.com/mmrath/gobase/go/pkg/db"
	"github.com/mmrath/gobase/go/pkg/errutil"
	"github.com/mmrath/gobase/go/pkg/model"
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
	err = s.db.RunInTx(ctx, func(tx *db.Tx) error {
		role.Role, err = s.roleDao.Find(tx, id)
		return err
	})
	return role, err
}

func (s *roleService) Create(ctx context.Context, roleAndPermission *model.RoleAndPermission) (err error) {
	err = s.db.RunInTx(ctx, func(tx *db.Tx) error {
		exists, err := s.roleDao.ExistsByName(tx, roleAndPermission.Role.Name)
		if err != nil {
			return errutil.Wrap(err, "error while checking if roleAndPermission exists")
		}
		if exists {
			return errutil.NewBadRequest(
				fmt.Sprintf("role with name %s already exists", roleAndPermission.Role.Name))
		}
		err = s.roleDao.Create(tx, &roleAndPermission.Role, roleAndPermission.Permissions)
		return err
	})
	return err
}

func (s *roleService) Update(ctx context.Context, roleAndPermission *model.RoleAndPermission) (err error) {
	err = s.db.RunInTx(ctx, func(tx *db.Tx) error {
		return s.roleDao.Update(tx, &roleAndPermission.Role, roleAndPermission.Permissions)
	})
	return err
}

func NewRoleService(db *db.DB) RoleService {
	return &roleService{
		db:      db,
		roleDao: model.NewRoleDao(),
	}
}
