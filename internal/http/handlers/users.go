package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gabeefranco/gonotes-api/internal/service"
	"github.com/go-chi/render"
)

type UsersHandler struct {
	Service service.UsersService
}

func NewUsersHandler(s service.UsersService) *UsersHandler {
	return &UsersHandler{
		Service: s,
	}
}

func (h UsersHandler) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, render.M{
			"error": "invalid request body",
		})
		return
	}

	user, err := h.Service.CreateUser(req.Email, req.Password)

	if err != nil {
		if errors.Is(err, service.ErrInvalidEmail) || errors.Is(err, service.ErrPasswordTooShort) {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, render.M{
				"error": err.Error(),
			})
			return
		} else if errors.Is(err, service.ErrUserAlreadyExists) {
			render.Status(r, http.StatusConflict)
			render.JSON(w, r, render.M{
				"error": err.Error(),
			})
			return
		} else {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, render.M{
				"error": "internal server error",
			})
			return
		}
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, render.M{
		"user": user,
	})
}
