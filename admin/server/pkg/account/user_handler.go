package account

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/spf13/cast"
	"gopkg.in/go-playground/validator.v9"
	"github.com/mmrath/gobase/common/errors"
	"github.com/mmrath/gobase/model"
	"net/http"
)

type UserHandler struct {
	userService UserService
}


func NewUserHandler(service UserService) *UserHandler{
	return &UserHandler{userService:service}
}

func (h *UserHandler) FindUser(id int64) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		err := validator.New().Var(&id, "required,int32")

		if err != nil {
			errors.RenderError(w, r, err)
			return
		}
		user, err := h.userService.Find(r.Context(), cast.ToInt32(id))

		if err != nil {
			errors.RenderError(w, r, err)
			return
		} else {
			render.Status(r, http.StatusOK)
			render.JSON(w, r, user)
			return
		}
	}
}

func (h *UserHandler) CreateUser(user *model.CreateUserRequest) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {

		if err != nil {
			errors.RenderError(w, r, err)
			return
		}
		user, err := h.userService.Find(r.Context(), cast.ToInt32(id))

		if err != nil {
			errors.RenderError(w, r, err)
			return
		} else {
			render.Status(r, http.StatusOK)
			render.JSON(w, r, user)
			return
		}
	}
}