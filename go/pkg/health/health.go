package health

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type Service interface {
	RegisterRoutes(r chi.Router)
}
type Checker interface {
	Name() string
	Check() bool
}

type service struct {
	checkers []Checker
}

func NewService(checkers ...Checker) Service {
	return &service{checkers}
}

func (s *service) RegisterRoutes(r chi.Router) {
	r.Get("/health", s.healthCheck)
}

func (s *service) healthCheck(w http.ResponseWriter, r *http.Request) {
	statuses := make(map[string]interface{})
	result := true

	for _, checker := range s.checkers {
		ok := checker.Check()
		if !ok {
			result = false
		} else {
			statuses[checker.Name()] = struct{ healthy bool }{true}
		}
	}

	render.JSON(w, r, map[string]interface{}{
		"healthy": result,
		"data":    statuses,
	})
}
