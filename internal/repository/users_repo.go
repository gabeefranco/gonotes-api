package repository

import (
	"context"

	"github.com/gabeefranco/gonotes-api/internal/domain"
)

type UsersRepository interface {
	Create(ctx context.Context, u *domain.User) error
	GetByID(ctx context.Context, id int64) (*domain.User, error)
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
	// Update(ctx context.Context, u *domain.User) error
	// Delete(ctx context.Context, id int64) error
}
