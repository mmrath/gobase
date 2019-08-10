package account

import (
	"fmt"
	"mmrath.com/gobase/common/errors"
	"mmrath.com/gobase/model"
)

type RoleService interface {
	Find(tx model.Tx, id int32) (*model.Role, error)
	Create(tx model.Tx, role *model.Role) (err error)
	Update(tx model.Tx, role *model.Role) (err error)
}

type roleService struct {
	roleDao model.RoleDao
}

func (s *roleService) Create(tx model.Tx, role *model.Role) (err error) {
	exists, err := s.roleDao.ExistsByName(tx, role.Name)
	if err != nil {
		return errors.NewInternal(err, "error while checking if role exists")
	}
	if exists {
		return errors.NewBadRequest(fmt.Sprintf("role with name %s already exists", role.Name))
	}
	return s.roleDao.Create(tx, role)
}

func (s *roleService) Update(tx model.Tx, role *model.Role) (err error) {
	panic("implement me")
}

func (s *roleService) Delete(tx model.Tx, id int32) (err error) {
	panic("implement me")
}

func NewRoleService(roleDao model.RoleDao) RoleService {
	return &roleService{roleDao: roleDao}
}

func (s *roleService) Find(tx model.Tx, id int32) (*model.Role, error) {
	return s.roleDao.Find(tx, id)
}
