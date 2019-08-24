package auth

import "net/http"

type Handler interface {
	Login() http.HandlerFunc
	Auth() http.HandlerFunc
	Token() http.HandlerFunc
	Authorize() http.HandlerFunc
}

func NewHandler() Handler{
	return &handler{}
}

type handler struct {

}

func (h *handler) Login() http.HandlerFunc {
	panic("implement me")
}

func (h *handler) Auth() http.HandlerFunc {
	panic("implement me")
}

func (h *handler) Token() http.HandlerFunc {
	panic("implement me")
}

func (h *handler) Authorize() http.HandlerFunc {
	panic("implement me")
}


