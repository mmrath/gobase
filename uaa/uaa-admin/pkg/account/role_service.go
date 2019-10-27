package account

import (
	"context"
	"fmt"
	"github.com/mmrath/gobase/common/errors"
	"github.com/mmrath/gobase/model"
)

type RoleService interface {
	Find(ctx context.Context, id int32) (*model.RoleAndPermission, error)
	Create(ctx context.Context, role *model.RoleAndPermission) (err error)
	Update(ctx context.Context, role *model.RoleAndPermission) (err error)
}

type roleService struct {
	db *model.DB
}

func (s *roleService) Find(ctx context.Context, id int32) (role *model.RoleAndPermission, err error) {
	err = s.db.RunTx(func(tx *model.Tx) error {
		role.Role, err = find(tx, id)
		return err
	})
	return role, err
}

func find(tx *model.Tx, id int32) (*model.Role, error) {
	roleDao := model.NewRoleDao(tx)
	return roleDao.Find(id)
}

func (s *roleService) Create(ctx context.Context, role *model.RoleAndPermission) (err error) {
	err = s.db.RunTx(func(tx *model.Tx) error {
		err = create(tx, role)
		return err
	})
	return err
}

func create(tx *model.Tx, roleAndPermission *model.RoleAndPermission) (err error) {
	roleDao := model.NewRoleDao(tx)

	exists, err := roleDao.ExistsByName(roleAndPermission.Role.Name)
	if err != nil {
		return errors.NewInternal(err, "error while checking if role exists")
	}
	if exists {
		return errors.NewBadRequest(fmt.Sprintf("role with name %s already exists", roleAndPermission.Role.Name))
	}
	return roleDao.Create(roleAndPermission.Role, roleAndPermission.Permissions)
}

func (s *roleService) Update(ctx context.Context, roleAndPermission *model.RoleAndPermission) (err error) {
	err = s.db.RunTx(func(tx *model.Tx) error {
		roleDao := model.NewRoleDao(tx)
		return roleDao.Update(roleAndPermission.Role, roleAndPermission.Permissions)
	})
	return err
}

func NewRoleService(db *model.DB) RoleService {
	return &roleService{db}
}
