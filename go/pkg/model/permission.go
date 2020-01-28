package model

import (
	"strings"

	"github.com/mmrath/gobase/go/pkg/db"
)

type Permission struct {
	ID          int32  `json:"id,omitempty"`
	Application string `json:"application,omitempty" sql:"default:null"`
	Authority   string `json:"authority,omitempty" sql:"default:null"`
	Description string `json:"description,omitempty" sql:"default:null"`
}

type PermissionDao interface {
	FindByID(tx *db.Tx, id int32) (Permission, error)
	FindAllByApplication(tx *db.Tx, app string) ([]Permission, error)
	FindAll(tx *db.Tx) ([]Permission, error)
}

type permissionDao struct {
}

func (p permissionDao) FindByID(tx *db.Tx, id int32) (perm Permission, err error) {
	err = tx.Find(&perm, id).Error
	return
}

func (p permissionDao) FindAllByApplication(tx *db.Tx, app string) ([]Permission, error) {
	var perms []Permission
	err := tx.Where("UPPER(application) = ", strings.ToUpper(app)).Find(perms).Error
	if err != nil {
		return nil, err
	}
	return perms, nil
}

func (p permissionDao) FindAll(tx *db.Tx) ([]Permission, error) {
	var perms []Permission
	err := tx.Find(perms).Error
	if err != nil {
		return nil, err
	}
	return perms, nil
}

func NewPermissionDao() PermissionDao {
	return &permissionDao{}
}
