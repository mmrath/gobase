package model

import (
	"mmrath.com/gobase/common/errors"
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

type RoleDao interface {
	Find(id int32) (*Role, error)
	ExistsByName(name string) (bool, error)
	Create(role *Role) error
	Update(role *Role) error
}

func NewRoleDao(tx *Tx) RoleDao {
	return &roleDao{tx}
}

type roleDao struct {
	tx *Tx
}

func (d *roleDao) Find(id int32) (*Role, error) {
	role := &Role{ID: id}
	err := d.tx.Select(role)
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (d *roleDao) ExistsByName(name string) (bool, error) {
	role := new(Role)
	return d.tx.Model(role).Where(" LOWER(name) = ?", strings.ToLower(name)).Exists()
}

func (d *roleDao) Create(role *Role) error {
	err := d.tx.Insert(role)
	if err != nil {
		return errors.NewInternal(err, "failed to insert role")
	}
	return createRolePermissions(d.tx, role)
}

func (d *roleDao) Update(role *Role) error {
	err := d.tx.Update(role)
	if err != nil {
		return errors.NewInternal(err, "failed to update role")
	}
	_, err = d.tx.Model(&RolePermission{}).Where("role_id = ?", role.ID).Delete()
	if err != nil {
		return errors.NewInternal(err, "failed to delete existing permissions of role")
	}
	return createRolePermissions(d.tx, role)
}

func createRolePermissions(tx Tx, role *Role) error {
	if role.Permissions == nil || len(role.Permissions) == 0 {
		return nil
	}
	rolePermissions := make([]RolePermission, len(role.Permissions))
	for i, perm := range role.Permissions {
		rolePermissions[i].RoleID = role.ID
		rolePermissions[i].PermissionID = perm.ID
	}
	return tx.Insert(rolePermissions)
}
