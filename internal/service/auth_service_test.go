package service_test

import (
	"testing"

	"github.com/gabeefranco/gonotes-api/internal/domain"
	"github.com/gabeefranco/gonotes-api/internal/repository"
	"github.com/gabeefranco/gonotes-api/internal/service"
	"github.com/go-chi/jwtauth"
	"golang.org/x/crypto/bcrypt"
)

func TestAuthenticateUser(t *testing.T) {

	testHash, _ := bcrypt.GenerateFromPassword([]byte("correctpassword"), bcrypt.DefaultCost)
	tokenAuth := jwtauth.New("HS256", []byte("secret"), nil)

	tests := []struct {
		name                      string
		inputEmail, inputPassword string
		wantError                 bool
		previousUsers             []domain.User
	}{
		{
			name:          "invalid email",
			inputEmail:    "email@example.com",
			inputPassword: "123456",
			previousUsers: []domain.User{},
			wantError:     true,
		},
		{
			name:          "invalid password",
			inputEmail:    "email@example.com",
			inputPassword: "wrongpassword",
			previousUsers: []domain.User{
				{
					Email:    "email@example.com",
					Password: "some hashed password",
				},
			},
			wantError: true,
		},
		{
			name:          "valid user",
			inputEmail:    "email@example.com",
			inputPassword: "correctpassword",
			previousUsers: []domain.User{
				{
					ID:       67,
					Email:    "email@example.com",
					Password: string(testHash),
				},
			},
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := repository.NewTestingUsersRepository()
			repo.Users = tt.previousUsers

			authService := service.NewAuthService(*tokenAuth, repo)
			token, err := authService.AuthenticateUser(tt.inputEmail, tt.inputPassword)

			if (err != nil) != tt.wantError {
				t.Errorf("expected error=%v; got %v", tt.wantError, err)
			}
			if tt.wantError {
				return
			}
			if !tt.wantError && token == "" {
				t.Fatalf("expected token not to be empty; got \"\"")
			}

			jwt, err := tokenAuth.Decode(token)
			if !tt.wantError && (err != nil) {
				t.Fatalf("expected result token to be valid; it could not be decoded")
			}
			userId, ok := jwt.Get("user_id")
			if !tt.wantError && !ok {
				t.Fatalf("expected result token to have \"user_id\" in its claims; it had not")
			}
			if !tt.wantError && userId.(float64) != 67 {
				t.Fatalf("expected user_id from token = 67; got %v", userId)
			}

		})
	}
}
