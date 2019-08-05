package model

import "time"

type AuditDetails struct {
	CreatedAt time.Time `json:"createdAt,omitempty" db:"created_at"`
	CreatedBy string    `json:"createdBy,omitempty" db:"created_by"`
	UpdatedAt time.Time `json:"updatedAt,omitempty" db:"updated_at"`
	UpdatedBy string    `json:"updatedBy,omitempty" db:"updated_by"`
	Version   uint32    `json:"version,omitempty" db:"version"`
}
