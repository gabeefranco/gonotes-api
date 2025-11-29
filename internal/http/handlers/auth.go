package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gabeefranco/gonotes-api/internal/service"
	"github.com/go-chi/render"
)

type AuthHandler struct {
	Service service.AuthService
}

func NewAuthHandler(s service.AuthService) *AuthHandler {
	return &AuthHandler{
		Service: s,
	}
}

func (h AuthHandler) HandleAuth(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	} // TODO: move typedef to a dto package

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, render.M{
			"error": "invalid request body",
		})
	}

	token, err := h.Service.AuthenticateUser(req.Email, req.Password)
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			render.Status(r, http.StatusUnauthorized)
			render.JSON(w, r, render.M{
				"error": err.Error(),
			})
			return
		} else {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, render.M{
				"error": err.Error(),
			})
			return
		}
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, render.M{
		"token": token,
	})

}
