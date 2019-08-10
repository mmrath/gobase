package model

import "strings"

type Role struct {
	tableName struct{} `sql:"role"`

	AuditDetails
	ID          int32        `json:"id,omitempty"`
	Name        string       `json:"name,omitempty" validate:"required"`
	Description string       `json:"description,omitempty" validate:"required"`
	Permissions []Permission `json:"permissions" sql:"-"`
}

type RoleDao interface {
	Find(tx Tx, id int32) (*Role, error)
	ExistsByName(tx Tx, name string) (bool, error)
	Create(tx Tx, role *Role) error
	Update(tx Tx, role *Role) error
}

func NewRoleDao() RoleDao {
	return &roleDao{}
}

type roleDao struct {
}

func (d *roleDao) Find(tx Tx, id int32) (*Role, error) {
	role := &Role{ID: id}
	err := tx.Select(role)
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (d *roleDao) ExistsByName(tx Tx, name string) (bool, error) {
	role := new(Role)
	return tx.Model(role).Where(" LOWER(name) = ?", strings.ToLower(name)).Exists()
}

func (d *roleDao) Create(tx Tx, role *Role) error {
	return tx.Insert(role)
}

func (d *roleDao) Update(tx Tx, role *Role) error {
	return tx.Update(role)
}
