package repository

import (
	"context"

	"github.com/gabeefranco/gonotes-api/internal/domain"
)

type NotesRepository interface {
	Create(ctx context.Context, n *domain.Note) error
	GetByID(ctx context.Context, id int64, userID int64) (*domain.Note, error)
	List(ctx context.Context, userID int64) ([]domain.Note, error)
	// Update(ctx context.Context, n *domain.Note) error
	// Delete(ctx context.Context, id int64, userID int64) error
}
