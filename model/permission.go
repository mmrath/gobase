package model

import (
	"strings"
)

type Permission struct {
	tableName struct{} `sql:"permission"`

	ID          int32  `json:"id,omitempty"`
	Application string `json:"application,omitempty"`
	Authority   string `json:"authority,omitempty"`
	Description string `json:"description,omitempty"`
}

type PermissionDao interface {
	FindById(id int32) (Permission, error)
	FindAllByApplication(app string) ([]Permission, error)
	FindAll() ([]Permission, error)
}

type permissionDao struct {
	tx Tx
}

func (p permissionDao) FindById(id int32) (perm Permission, err error) {
	perm = Permission{ID: id}
	err = p.tx.Select(&perm)
	return
}

func (p permissionDao) FindAllByApplication(app string) ([]Permission, error) {
	var perms []Permission
	err := p.tx.Model(&perms).Where("UPPER(application) = ", strings.ToUpper(app)).Select()
	if err != nil {
		return nil, err
	}
	return perms, nil
}

func (p permissionDao) FindAll() ([]Permission, error) {
	var perms []Permission
	err := p.tx.Model(&perms).Select()
	if err != nil {
		return nil, err
	}
	return perms, nil
}

func NewPermissionDao(tx Tx) PermissionDao {
	return &permissionDao{tx}
}
