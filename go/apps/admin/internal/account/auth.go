package account

import (
	"encoding/gob"
	"net/http"

	"github.com/go-chi/render"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"

	"github.com/mmrath/gobase/go/pkg/errutil"
	"github.com/mmrath/gobase/go/pkg/model"
)

var store *sessions.CookieStore
var sessionCookieName = "SESSION_ID"

func init() {
	authKeyOne := securecookie.GenerateRandomKey(64)
	encryptionKeyOne := securecookie.GenerateRandomKey(32)

	store = sessions.NewCookieStore(
		authKeyOne,
		encryptionKeyOne,
	)

	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   60 * 15,
		HttpOnly: true,
	}

	gob.Register(model.Staff{})
}

type authHandler struct {
}

type AuthHandler interface {
	Login(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
}

func NewAuthHandler() *authHandler {
	return &authHandler{}
}

func (h *authHandler) Login(w http.ResponseWriter, r *http.Request) {
	data := model.LoginRequest{}
	if err := render.DecodeJSON(r.Body, &data); err != nil {
		render.JSON(w, r, err)
		return
	}

	session, err := store.Get(r, sessionCookieName)
	if err != nil {
		errutil.RenderError(w, r, err)
		return
	}

	if data.Email == "test@test.com" && data.Password == "password" {
		staff := model.Staff{
			ID:        1,
			Email:     data.Email,
			FirstName: "John",
			LastName:  "Doe",
		}

		session.Values["user"] = staff

		err = session.Save(r, w)
		if err != nil {
			errutil.RenderError(w, r, err)
			return
		}
		render.Status(r, http.StatusOK)
		render.JSON(w, r, staff)
		return
	}
	errutil.RenderError(w, r, errutil.NewUnauthorized("invalid username or password"))
}

func Logout(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, sessionCookieName)
	if err != nil {
		errutil.RenderError(w, r, err)
		return
	}

	session.Values["user"] = model.Staff{}
	session.Options.MaxAge = -1

	err = session.Save(r, w)
	if err != nil {
		errutil.RenderError(w, r, err)
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, model.Staff{})
}
