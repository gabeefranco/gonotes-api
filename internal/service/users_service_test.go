package service_test

import (
	"testing"

	"github.com/gabeefranco/gonotes-api/internal/domain"
	"github.com/gabeefranco/gonotes-api/internal/repository"
	"github.com/gabeefranco/gonotes-api/internal/service"
)

func TestCreateUser(t *testing.T) {

	tests := []struct {
		name                      string
		inputEmail, inputPassword string
		want                      *domain.User
		wantError                 bool
		previousUsers             []domain.User
	}{
		{
			name:          "invalid email",
			inputEmail:    "invalid_email",
			inputPassword: "123456",
			want:          nil,
			wantError:     true,
		},
		{
			name:          "invalid password",
			inputEmail:    "email@example.com",
			inputPassword: "12345",
			want:          nil,
			wantError:     true,
		},
		{
			previousUsers: []domain.User{
				{
					ID:       1,
					Email:    "email@example.com",
					Password: "123456",
				},
			},
			inputEmail:    "email@example.com",
			inputPassword: "123456",
			want:          nil,
			wantError:     true,
		},
		{
			name:          "valid",
			inputEmail:    "email@example.com",
			inputPassword: "123456",
			want: &domain.User{
				Email: "email@example.com",
			},
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := repository.NewTestingUsersRepository()
			repo.Users = tt.previousUsers
			usersService := service.NewUsersService(repo)

			u, err := usersService.CreateUser(tt.inputEmail, tt.inputPassword)
			if (err != nil) != tt.wantError {
				t.Fatalf("expected error=%v; got %v", tt.wantError, err)
			}
			if tt.want != nil && u != nil && tt.want.Email != u.Email {
				t.Errorf("expected created user email=%s; got %s", tt.want.Email, u.Email)
			}
			if u != nil && tt.inputPassword == u.Password {
				t.Errorf("expected password to be hashed; it was not")
			}
		})
	}
}
