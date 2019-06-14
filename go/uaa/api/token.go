package main

import (
	"time"

	"gopkg.in/oauth2.v3"
)

type Token struct {
	ClientID    string
	UserID      string
	RedirectURI string
	Scope       string

	Code          string
	CodeCreatedAt time.Time
	CodeExpiresIn time.Duration

	AccessToken          string
	AccessTokenCreatedAt time.Time
	AccessTokenExpiresIn time.Duration

	RefreshToken          string
	RefreshTokenCreatedAt time.Time
	RefreshTokenExpiresIn time.Duration
}

func (t *Token) New() oauth2.TokenInfo {
	return &Token{}
}
func (t *Token) GetClientID() string {
	return t.ClientID
}
func (t *Token) SetClientID(clientID string) {
	t.ClientID = clientID
}

func (t *Token) GetUserID() string {
	return t.UserID
}
func (t *Token) SetUserID(userID string) {
	t.UserID = userID
}
func (t *Token) GetRedirectURI() string {
	return t.RedirectURI
}
func (t *Token) SetRedirectURI(redirectURI string) {
	t.RedirectURI = redirectURI
}
func (t *Token) GetScope() string {
	return t.Scope
}
func (t *Token) SetScope(scope string) {
	t.Scope = scope
}
func (t *Token) GetCode() string {
	return t.Code
}
func (t *Token) SetCode(code string) {
	t.Code = code
}
func (t *Token) GetCodeCreateAt() time.Time {
	return t.CodeCreatedAt
}
func (t *Token) SetCodeCreateAt(tm time.Time) {
	t.CodeCreatedAt = tm
}
func (t *Token) GetCodeExpiresIn() time.Duration {
	return t.CodeExpiresIn
}
func (t *Token) SetCodeExpiresIn(tm time.Duration) {
	t.CodeExpiresIn = tm
}
func (t *Token) GetAccess() string {
	return t.AccessToken
}
func (t *Token) SetAccess(token string) {
	t.AccessToken = token
}
func (t *Token) GetAccessCreateAt() time.Time {
	return t.AccessTokenCreatedAt
}
func (t *Token) SetAccessCreateAt(tm time.Time) {
	t.AccessTokenCreatedAt = tm
}
func (t *Token) GetAccessExpiresIn() time.Duration {
	return t.AccessTokenExpiresIn
}
func (t *Token) SetAccessExpiresIn(tm time.Duration) {
	t.AccessTokenExpiresIn = tm
}
func (t *Token) GetRefresh() string {
	return t.RefreshToken
}
func (t *Token) SetRefresh(token string) {
	t.RefreshToken = token
}
func (t *Token) GetRefreshCreateAt() time.Time {
	return t.RefreshTokenCreatedAt
}
func (t *Token) SetRefreshCreateAt(tm time.Time) {
	t.RefreshTokenCreatedAt = tm
}
func (t *Token) GetRefreshExpiresIn() time.Duration {
	return t.RefreshTokenExpiresIn
}
func (t *Token) SetRefreshExpiresIn(tm time.Duration) {
	t.RefreshTokenExpiresIn = tm
}
