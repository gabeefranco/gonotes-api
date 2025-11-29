package routes

import (
	"github.com/gabeefranco/gonotes-api/internal/http/handlers"
	"github.com/go-chi/chi/v5"
)

type AuthRoutes struct {
	Handler handlers.AuthHandler
}

func NewAuthRoutes(h handlers.AuthHandler) *AuthRoutes {
	return &AuthRoutes{
		Handler: h,
	}
}

func (a AuthRoutes) Setup(r *chi.Mux) {
	r.Post("/auth", a.Handler.HandleAuth)
}
