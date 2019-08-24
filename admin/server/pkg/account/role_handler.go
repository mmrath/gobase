package account

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"gopkg.in/go-playground/validator.v9"
	"github.com/rs/zerolog/log"
	"mmrath.com/gobase/common/errors"
	"mmrath.com/gobase/model"
	"net/http"
)

type RoleHandler struct {
	roleService RoleService
}

func NewRoleHandler(roleService RoleService) *RoleHandler {
	return &RoleHandler{roleService: roleService}
}

func (h *RoleHandler) FindRole() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		err := validator.New().Var(&id, "")
		err := h.Service.Activate(key)

		if err != nil {
			errors.RenderError(w, r, err)
			return
		} else {
			render.Status(r, http.StatusOK)
			render.PlainText(w, r, http.StatusText(http.StatusOK))
			return
		}
	}
}


func (h *RoleHandler) CreateRole() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		role := new(model.Role)

		if err := render.DecodeJSON(r.Body, role); err != nil {
			render.JSON(w, r, err)
			return
		}

		err := h.roleService.Create(r.Context(), role)

		if err != nil {
			log.Error().Err(err).Msg("error creating role")
			errors.RenderError(w, r, err)
			return
		} else {
			render.Status(r, http.StatusOK)
			render.JSON(w, r, role)
			return
		}
	}
}
