package repository

import (
	"context"

	"github.com/gabeefranco/gonotes-api/internal/domain"
)

type TestingUsersRepository struct {
	Users []domain.User
}

func NewTestingUsersRepository() *TestingUsersRepository {
	return &TestingUsersRepository{
		Users: make([]domain.User, 0),
	}
}

func (r TestingUsersRepository) Create(ctx context.Context, u *domain.User) error {
	r.Users = append(r.Users, *u)

	return nil
}

func (r TestingUsersRepository) GetByID(ctx context.Context, id int64) (*domain.User, error) {
	for _, u := range r.Users {
		if u.ID == id {
			return &u, nil
		}
	}

	return nil, nil
}

func (r TestingUsersRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	for _, u := range r.Users {
		if u.Email == email {
			return &u, nil
		}
	}

	return nil, nil
}
