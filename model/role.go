package model

import (
	"github.com/mmrath/gobase/common/error_util"
	"github.com/mmrath/gobase/pkg/db"
	"github.com/rs/zerolog/log"
)

type Role struct {
	AuditDetails
	ID          int32        `json:"id,omitempty"`
	Name        string       `json:"name,omitempty" validate:"required"`
	Description string       `json:"description,omitempty" validate:"required"`
	Permissions []Permission `json:"permissions" sql:"-"`
}

type RolePermission struct {
	RoleID       int32 `json:"roleId,omitempty" validate:"required"`
	PermissionID int32 `json:"permissionId,omitempty" validate:"required"`
}

type RoleAndPermission struct {
	Role        Role
	Permissions []int32
}

type RoleDao interface {
	Find(tx *db.Tx, id int32) (Role, error)
	FindPermissionsByRoleId(tx *db.Tx, id int32) ([]int32, error)
	ExistsByName(tx *db.Tx, name string) (bool, error)
	Create(tx *db.Tx, role *Role, permissions []int32) error
	Update(tx *db.Tx, role *Role, permissions []int32) error
}

type roleDao struct {
}

func (dao *roleDao) Find(tx *db.Tx, id int32) (Role, error) {
	role := Role{}
	err := tx.First(&role, id).Error
	return role, err
}

func (dao *roleDao) FindPermissionsByRoleId(tx *db.Tx, id int32) ([]int32, error) {
	var permissions []int32
	err := tx.Model(&RolePermission{}).
		Where("role_id = ?", id).
		Pluck("permission_id", &permissions).Error
	return permissions, err
}

func (dao *roleDao) ExistsByName(tx *db.Tx, name string) (bool, error) {
	count := 0
	err := tx.Find(&Role{}, " LOWER(name) = LOWER(?)", name).Count(&count).Error
	return count != 0, err
}

func (dao *roleDao) Create(tx *db.Tx, role *Role, permissions []int32) error {
	err := tx.Model(role).Create(role).Error
	if err != nil {
		return err
	}

	return dao.createRolePermissions(tx, role.ID, permissions)
}

func (dao *roleDao) Update(tx *db.Tx, role *Role, permissions []int32) error {
	err := tx.Save(role).Error
	if err != nil {
		log.Error().
			Int32("roleId", role.ID).
			Err(err).
			Msg("failed to update role")

		return error_util.NewInternal(err, "failed to update role")
	}

	return dao.createRolePermissions(tx, role.ID, permissions)
}

func (dao *roleDao) createRolePermissions(tx *db.Tx, roleId int32, permissions []int32) error {
	err := tx.Delete(&RolePermission{}, "role_id = ?", roleId).Error
	if err != nil {
		log.Error().
			Int32("roleId", roleId).
			Err(err).
			Msg("failed to delete existing permissions of role")
		return error_util.NewInternal(err, "failed to delete existing permissions of role")
	}
	if permissions == nil || len(permissions) == 0 {
		return nil
	}
	rolePermissions := make([]RolePermission, len(permissions))
	for i, perm := range permissions {
		rolePermissions[i].RoleID = roleId
		rolePermissions[i].PermissionID = perm
		err := tx.Create(rolePermissions).Error
		if err != nil {
			return err
		}
	}
	return nil
}
