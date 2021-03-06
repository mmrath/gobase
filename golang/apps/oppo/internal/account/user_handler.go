package account

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/spf13/cast"
	"gopkg.in/go-playground/validator.v9"

	"github.com/mmrath/gobase/golang/pkg/db"
	"github.com/mmrath/gobase/golang/pkg/errutil"
	"github.com/mmrath/gobase/golang/pkg/model"
)

type UserHandler struct {
	userService UserService
}

func NewUserHandler(database *db.DB) *UserHandler {
	userService := NewUserService(database)
	return &UserHandler{userService: userService}
}

func (h *UserHandler) FindUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := validator.New().Var(&id, "required,int32")

	if err != nil {
		errutil.RenderError(w, r, err)
		return
	}
	user, err := h.userService.FindUserByID(r.Context(), cast.ToInt64(id))

	if err != nil {
		errutil.RenderError(w, r, err)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, user)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	userCreateReq := model.CreateUserRequest{}

	if err := render.DecodeJSON(r.Body, &userCreateReq); err != nil {
		render.JSON(w, r, err)
		return
	}

	user, err := h.userService.CreateUser(r.Context(), &userCreateReq)

	if err != nil {
		errutil.RenderError(w, r, err)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, user)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	user := model.User{}

	if err := render.DecodeJSON(r.Body, &user); err != nil {
		render.JSON(w, r, err)
		return
	}

	err := h.userService.UpdateUser(r.Context(), &user)

	if err != nil {
		errutil.RenderError(w, r, err)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, user)
}
