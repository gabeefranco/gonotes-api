package repository

import (
	"context"
	"database/sql"

	"github.com/gabeefranco/gonotes-api/internal/domain"
)

type SqlNotesRepository struct {
	DB *sql.DB
}

func NewSqlNotesRepository(db *sql.DB) *SqlNotesRepository {
	return &SqlNotesRepository{
		DB: db,
	}
}

func (r SqlNotesRepository) Create(ctx context.Context, n *domain.Note) error {
	const query = `
		INSERT INTO notes (user_id, title, content) 
		VALUES ($1, $2, $3)
		RETURNING id
	`

	err := r.DB.QueryRowContext(ctx, query, n.UserID, n.Title, n.Content).Scan(&n.ID)

	return err

}

func (r SqlNotesRepository) GetByID(ctx context.Context, id int64, userID int64) (*domain.Note, error) {
	const query = `
		SELECT id, user_id, title, content, created_at, updated_at
		FROM notes
		WHERE id = $1 AND user_id = $2
		LIMIT 1
	`

	row := r.DB.QueryRowContext(ctx, query, id, userID)

	var n domain.Note
	err := row.Scan(
		&n.ID,
		&n.UserID,
		&n.Title,
		&n.Content,
		&n.CreatedAt,
		&n.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &n, nil

}

func (r SqlNotesRepository) List(ctx context.Context, userID int64) ([]domain.Note, error) {
	const query = `
		SELECT id, user_id, title, content, created_at, updated_at
		FROM notes
		WHERE user_id = $1
		ORDER BY created_at DESC
	`

	rows, err := r.DB.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	notes := make([]domain.Note, 0)

	for rows.Next() {
		var n domain.Note
		err := rows.Scan(
			&n.ID,
			&n.UserID,
			&n.Title,
			&n.Content,
			&n.CreatedAt,
			&n.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		notes = append(notes, n)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return notes, nil
}
