package account

import (
	"context"
	"fmt"
	"mmrath.com/gobase/common/errors"
	"mmrath.com/gobase/model"
)

type RoleService interface {
	Find(ctx context.Context, id int32) (*model.Role, error)
	Create(ctx context.Context, role *model.Role) (err error)
	Update(ctx context.Context, role *model.Role) (err error)
}

type roleService struct {
	db model.DB
}

func (s *roleService) Find(ctx context.Context, id int32) (role *model.Role, err error) {
	err = s.db.RunTx(func(tx *model.Tx) error {
		role, err = find(tx, id)
		return err
	})
	return role, err
}

func (s *roleService) Create(ctx context.Context, role *model.Role) (err error) {
	err = s.db.RunTx(func(tx *model.Tx) error {
		err = create(tx, role)
		return err
	})
	return err
}

func create(tx *model.Tx, role *model.Role) (err error) {
	roleDao := model.NewRoleDao(tx)

	exists, err := roleDao.ExistsByName(role.Name)
	if err != nil {
		return errors.NewInternal(err, "error while checking if role exists")
	}
	if exists {
		return errors.NewBadRequest(fmt.Sprintf("role with name %s already exists", role.Name))
	}
	return roleDao.Create(role)
}

func (s *roleService) Update(ctx context.Context, role *model.Role) (err error) {
	err = s.db.RunTx(func(tx *model.Tx) error {
		roleDao := model.NewRoleDao(tx)
		return roleDao.Update(role)
	})
	return err
}


func find(tx *model.Tx, id int32) (*model.Role, error) {
	roleDao := model.NewRoleDao(tx)
	return roleDao.Find(id)
}

func NewRoleService(db model.DB) RoleService {
	return &roleService{db}
}


