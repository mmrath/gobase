package model

import (
	"time"
)

type AuthToken struct {
	ID         uint64    `json:"id,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
	UserID     uint64    `json:"-"`
	Token      string    `json:"-"`
	ExpiresAt  time.Time `json:"-"`
	Mobile     bool      `sql:",notnull" json:"mobile"`
	Identifier string    `json:"identifier,omitempty"`
}
