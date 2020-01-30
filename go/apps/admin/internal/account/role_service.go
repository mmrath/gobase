package account

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"

	"github.com/mmrath/gobase/go/pkg/db"
	"github.com/mmrath/gobase/go/pkg/errutil"
	"github.com/mmrath/gobase/go/pkg/model"
)

type RoleService interface {
	FindRoleByID(ctx context.Context, id int32) (model.RoleAndPermission, error)
	CreateRole(ctx context.Context, role *model.RoleAndPermission) error
	UpdateRole(ctx context.Context, role *model.RoleAndPermission) error
}

type roleService struct {
	db      *db.DB
	roleDao model.RoleDao
}

func (s *roleService) FindRoleByID(ctx context.Context, id int32) (model.RoleAndPermission, error) {
	var role model.RoleAndPermission
	var err error
	err = s.db.RunInTx(ctx, func(tx *db.Tx) error {
		role, err = s.findRoleTx(tx, id)
		return err
	})
	return role, err
}

func (s *roleService) findRoleTx(tx *db.Tx, id int32) (role model.RoleAndPermission, err error) {
	role.Role, err = s.roleDao.Find(tx, id)
	if err != nil {
		if !db.IsNoDataFound(err) {
			return role, errutil.NewBadRequest("no data found for role")
		}
		return role, errutil.Wrap(err, "failed while fetching role")
	}
	role.Permissions, err = s.roleDao.FindPermissionsByRoleID(tx, role.Role.ID)
	if err != nil {
		// ignore anything but no data found
		if !db.IsNoDataFound(err) {
			return role, errutil.Wrap(err, "failed while fetching role permissions")
		}
	}
	return role, nil
}

func (s *roleService) CreateRole(ctx context.Context, roleAndPermission *model.RoleAndPermission) error {
	err := s.db.RunInTx(ctx, func(tx *db.Tx) error {
		return s.createRoleTx(tx, roleAndPermission)
	})
	return err
}

func (s *roleService) createRoleTx(tx *db.Tx, roleAndPermission *model.RoleAndPermission) error {
	exists, err := s.roleDao.ExistsByName(tx, roleAndPermission.Role.Name)
	if err != nil {
		return errutil.Wrap(err, "error while checking if role already exists")
	}
	if exists {
		return errutil.NewBadRequest(
			fmt.Sprintf("role with name %s already exists", roleAndPermission.Role.Name))
	}
	err = s.roleDao.Create(tx, &roleAndPermission.Role, roleAndPermission.Permissions)
	if err != nil {
		dbErr := tx.Rollback().Error
		if dbErr != nil {
			// we return the original error
			log.Error().Err(dbErr).Msg("error while rolling back transaction")
		}
		return errutil.Wrap(err, "failed to create role")
	}
	return nil
}

func (s *roleService) UpdateRole(ctx context.Context, roleAndPermission *model.RoleAndPermission) error {
	err := s.db.RunInTx(ctx, func(tx *db.Tx) error {
		return s.roleDao.Update(tx, &roleAndPermission.Role, roleAndPermission.Permissions)
	})
	return err
}

func NewRoleService(
	database *db.DB) RoleService {
	return &roleService{
		db:      database,
		roleDao: model.NewRoleDao(),
	}
}
