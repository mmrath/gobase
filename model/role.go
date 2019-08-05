package model

type Role struct {
	tableName struct{} `sql:"role"`

	AuditDetails
	ID          uint32       `json:"id,omitempty"`
	Name        string       `json:"name,omitempty" validate:"required"`
	Description string       `json:"description,omitempty" validate:"required"`
	Permissions []Permission `json:"permissions" sql:"-"`
}
