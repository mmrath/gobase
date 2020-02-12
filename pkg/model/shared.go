package model

import "time"

type AuditDetails struct {
	UpdatedAt time.Time `json:"updatedAt,omitempty" db:"updated_at"`
	UpdatedBy string    `json:"updatedBy,omitempty" db:"updated_by"`
	Version   uint32    `json:"version,omitempty" db:"version"`
}
