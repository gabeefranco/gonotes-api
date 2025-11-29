package service

import (
	"context"
	"errors"
	"time"

	"github.com/gabeefranco/gonotes-api/internal/repository"
	"github.com/go-chi/jwtauth"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	Provider jwtauth.JWTAuth
	Users    repository.UsersRepository
}

func NewAuthService(p jwtauth.JWTAuth, r repository.UsersRepository) *AuthService {
	return &AuthService{
		Provider: p,
		Users:    r,
	}
}

var ErrInvalidCredentials = errors.New("invalid credentials")

func (s AuthService) AuthenticateUser(email string, password string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	user, err := s.Users.GetByEmail(ctx, email)
	if err != nil {
		return "", ErrInternal
	}

	if user == nil {
		return "", ErrInvalidCredentials
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", ErrInvalidCredentials
	}

	_, tokenString, err := s.Provider.Encode(map[string]interface{}{"user_id": user.ID})
	if err != nil {
		return "", ErrInternal
	}

	return tokenString, nil
}
