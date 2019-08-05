package main

import (
	"strconv"

	"gopkg.in/oauth2.v3"
)

type OauthClientInfo struct {
	ID                   int32
	ClientID             string
	Secret               string
	Domain               string
	RedirectURI          string
	Scope                string
	GrantTypes           string
	AccessTokenValidity  int32
	RefreshTokenValidity int32
	AutoApprove          bool
}

type clientStore struct {
	data map[string]OauthClientInfo
}

// NewClientStore create client store
func NewClientStore() oauth2.ClientStore {
	data := make(map[string]OauthClientInfo)
	data["Ara Client Web"] = OauthClientInfo{
		ID:                   1,
		ClientID:             "Ara Client Web",
		Secret:               "",
		Domain:               "",
		RedirectURI:          "http://localhost:9094",
		Scope:                "",
		GrantTypes:           "",
		AccessTokenValidity:  60,
		RefreshTokenValidity: 120,
		AutoApprove:          true,
	}
	return &clientStore{
		data: data,
	}
}

func (s *clientStore) GetByID(id string) (client oauth2.ClientInfo, err error) {

	if c, ok := s.data[id]; ok {
		client = &c
		return
	}
	err = errors2.NewBadRequest("invalid client id:" + id)
	return

}

func (c *OauthClientInfo) GetID() string {
	return c.ClientID
}
func (c *OauthClientInfo) GetSecret() string {
	return c.Secret
}
func (c *OauthClientInfo) GetDomain() string {
	return c.Domain
}
func (c *OauthClientInfo) GetUserID() string {
	return strconv.Itoa(int(c.ID))
}
