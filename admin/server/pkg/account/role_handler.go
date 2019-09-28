package account

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cast"
	"gopkg.in/go-playground/validator.v9"
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
		err := validator.New().Var(&id, "required")

		if err != nil {
			errors.RenderError(w, r, err)
			return
		}

		role, err := h.roleService.Find(r.Context(), cast.ToInt32(id))

		if err != nil {
			errors.RenderError(w, r, err)
			return
		} else {
			render.Status(r, http.StatusOK)
			render.JSON(w, r, role)
			return
		}
	}
}

func (h *RoleHandler) CreateRole() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		role := new(model.RoleAndPermission)

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

func (h *RoleHandler) UpdateRole() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		role := new(model.RoleAndPermission)

		if err := render.DecodeJSON(r.Body, role); err != nil {
			render.JSON(w, r, err)
			return
		}

		err := h.roleService.Update(r.Context(), role)

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