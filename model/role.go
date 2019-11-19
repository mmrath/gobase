package model

import (
	"context"
	"strings"

	"github.com/go-pg/pg"
	"github.com/rs/zerolog/log"

	"github.com/mmrath/gobase/common/error_util"
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
	Role        *Role
	Permissions []int32
}

type RoleDao interface {
	Find(ctx context.Context, id int32) (*Role, error)
	FindPermissionsByRoleId(ctx context.Context, id int32) ([]int32, error)
	ExistsByName(ctx context.Context, name string) (bool, error)
	Create(ctx context.Context, role *Role, permissions []int32) error
	Update(ctx context.Context, role *Role, permissions []int32) error
}

func NewRoleDao() *roleDao {
	return &roleDao{}
}

type roleDao struct {
}

func (d *roleDao) Find(ctx context.Context,id int32) (*Role, error) {
	role := Role{ID: id}
	err := TxFromContext(ctx).Select(&role)
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (d *roleDao) FindPermissionsByRoleId(ctx context.Context, id int32) ([]int32, error) {
	var permissions []int32
	err := TxFromContext(ctx).Model(&RolePermission{}).Column("permission_id").Where("role_id = ?", id).Select(&permissions)
	if err != nil {
		return nil, err
	}
	return permissions, nil
}

func (d *roleDao) ExistsByName(ctx context.Context, name string) (bool, error) {
	role := new(Role)
	return TxFromContext(ctx).Model(role).Where(" LOWER(name) = ?", strings.ToLower(name)).Exists()
}

func (d *roleDao) Create(ctx context.Context, role *Role, permissions []int32) error {
	err := TxFromContext(ctx).Insert(role)
	if err != nil {
		log.Error().
			Int32("roleId", role.ID).
			Err(err).
			Msg("failed to update role")
		return error_util.NewInternal(err, "failed to insert role")
	}
	return createRolePermissions(TxFromContext(ctx), role.ID, permissions)
}

func (d *roleDao) Update(ctx context.Context, role *Role, permissions []int32) error {
	err := TxFromContext(ctx).Update(role)
	if err != nil {
		log.Error().
			Int32("roleId", role.ID).
			Err(err).
			Msg("failed to update role")

		return error_util.NewInternal(err, "failed to update role")
	}

	return createRolePermissions(TxFromContext(ctx), role.ID, permissions)
}

func createRolePermissions(tx *pg.Tx, roleId int32, permissions []int32) error {
	_, err := tx.Model(&RolePermission{}).Where("role_id = ?", roleId).Delete()
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
	}
	return tx.Insert(rolePermissions)
}
