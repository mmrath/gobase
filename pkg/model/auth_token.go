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

type AuthTokenDao struct {
	tx *Tx
}

func NewAuthTokenDao(tx *Tx) *AuthTokenDao {
	return &AuthTokenDao{
		tx,
	}
}

func (s *AuthTokenDao) GetToken(t string) (*AuthToken, error) {
	token := AuthToken{Token: t}
	err := s.tx.Model(&token).
		Where("token = ?token").
		First()

	return &token, err
}

// CreateOrUpdateToken creates or updates an existing refresh token.
func (s *AuthTokenDao) CreateOrUpdateToken(t *AuthToken) error {
	var err error
	if t.ID == 0 {
		err = s.tx.Insert(t)
	} else {
		err = s.tx.Update(t)
	}
	return err
}

// DeleteToken deletes a refresh token.
func (s *AuthTokenDao) DeleteToken(t *AuthToken) error {
	err := s.tx.Delete(t)
	return err
}

// PurgeExpiredToken deletes expired refresh token.
func (s *AuthTokenDao) PurgeExpiredToken() error {
	_, err := s.tx.Model(&AuthToken{}).
		Where("expires_at < ?", time.Now()).
		Delete()

	return err
}
