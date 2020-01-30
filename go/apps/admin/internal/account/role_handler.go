package account

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cast"
	"gopkg.in/go-playground/validator.v9"

	"github.com/mmrath/gobase/go/pkg/db"
	"github.com/mmrath/gobase/go/pkg/errutil"
	"github.com/mmrath/gobase/go/pkg/model"
)

type RoleHandler struct {
	roleService RoleService
}

func NewRoleHandler(database *db.DB) *RoleHandler {
	roleService := NewRoleService(database)
	return &RoleHandler{roleService: roleService}
}

func (h *RoleHandler) FindRole() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		err := validator.New().Var(&id, "required")

		if err != nil {
			errutil.RenderError(w, r, err)
			return
		}

		role, err := h.roleService.FindRoleByID(r.Context(), cast.ToInt32(id))

		if err != nil {
			errutil.RenderError(w, r, err)
			return
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, role)
	}
}

func (h *RoleHandler) CreateRole() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		role := new(model.RoleAndPermission)

		if err := render.DecodeJSON(r.Body, role); err != nil {
			render.JSON(w, r, err)
			return
		}

		err := h.roleService.CreateRole(r.Context(), role)

		if err != nil {
			log.Error().Err(err).Msg("error creating role")
			errutil.RenderError(w, r, err)
			return
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, role)
	}
}

func (h *RoleHandler) UpdateRole() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		role := new(model.RoleAndPermission)

		if err := render.DecodeJSON(r.Body, role); err != nil {
			render.JSON(w, r, err)
			return
		}

		err := h.roleService.UpdateRole(r.Context(), role)

		if err != nil {
			log.Error().Err(err).Msg("error creating role")
			errutil.RenderError(w, r, err)
			return
		}
		render.Status(r, http.StatusOK)
		render.JSON(w, r, role)
	}
}
