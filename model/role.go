package model

import (
	"github.com/rs/zerolog/log"
	"github.com/mmrath/gobase/common/errors"
	"strings"
)

type Role struct {
	tableName struct{} `sql:"role"`

	AuditDetails
	ID          int32        `json:"id,omitempty"`
	Name        string       `json:"name,omitempty" validate:"required"`
	Description string       `json:"description,omitempty" validate:"required"`
	Permissions []Permission `json:"permissions" sql:"-"`
}

type RolePermission struct {
	tableName struct{} `sql:"role_permission"`

	RoleID       int32 `json:"roleId,omitempty" validate:"required"`
	PermissionID int32 `json:"permissionId,omitempty" validate:"required"`
}

type RoleAndPermission struct {
	Role *Role
	Permissions []int32
}

type RoleDao interface {
	Find(id int32) (*Role, error)
	FindPermissionsByRoleId(id int32) ([]int32, error)
	ExistsByName(name string) (bool, error)
	Create(role *Role, permissions []int32) error
	Update(role *Role, permissions []int32) error
}

func NewRoleDao(tx *Tx) RoleDao {
	return &roleDao{tx}
}

type roleDao struct {
	tx *Tx
}

func (d *roleDao) Find(id int32) (*Role, error) {
	role := Role{ID: id}
	err := d.tx.Select(&role)
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (d *roleDao) FindPermissionsByRoleId(id int32) ([]int32, error) {
	var permissions []int32
	err := d.tx.Model(&RolePermission{}).Column("permission_id").Where("role_id = ?", id).Select(&permissions)
	if err != nil {
		return nil, err
	}
	return permissions, nil
}

func (d *roleDao) ExistsByName(name string) (bool, error) {
	role := new(Role)
	return d.tx.Model(role).Where(" LOWER(name) = ?", strings.ToLower(name)).Exists()
}



func (d *roleDao) Create(role *Role, permissions []int32) error {
	err := d.tx.Insert(role)
	if err != nil {
		log.Error().
			Int32("roleId", role.ID).
			Err(err).
			Msg("failed to update role")
		return errors.NewInternal(err, "failed to insert role")
	}
	return createRolePermissions(d.tx, role.ID, permissions)
}

func (d *roleDao) Update(role *Role, permissions []int32) error {
	err := d.tx.Update(role)
	if err != nil {
		log.Error().
			Int32("roleId", role.ID).
			Err(err).
			Msg("failed to update role")

		return errors.NewInternal(err, "failed to update role")
	}

	return createRolePermissions(d.tx, role.ID, permissions)
}

func createRolePermissions(tx *Tx, roleId int32, permissions []int32) error {
	_, err := tx.Model(&RolePermission{}).Where("role_id = ?", roleId).Delete()
	if err != nil {
		log.Error().
			Int32("roleId", roleId).
			Err(err).
			Msg("failed to delete existing permissions of role")
		return errors.NewInternal(err, "failed to delete existing permissions of role")
	}
	if permissions == nil || len(permissions) == 0 {
		return nil
	}
	rolePermissions := make([]RolePermission, len(permissions))
	for i, perm := range permissions {
		rolePermissions[i].RoleID = roleId
		rolePermissions[i].PermissionID =perm
	}
	return tx.Insert(rolePermissions)
}
