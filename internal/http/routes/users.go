package routes

import (
	"github.com/gabeefranco/gonotes-api/internal/http/handlers"
	"github.com/go-chi/chi/v5"
)

type UsersRoutes struct {
	Handler handlers.UsersHandler
}

func NewUsersRoutes(h handlers.UsersHandler) *UsersRoutes {
	return &UsersRoutes{
		Handler: h,
	}
}

func (u UsersRoutes) Setup(r *chi.Mux) {
	r.Post("/users", u.Handler.HandleCreateUser)
}
