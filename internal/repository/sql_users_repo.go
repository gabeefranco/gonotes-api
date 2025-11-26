package repository

import (
	"context"
	"database/sql"

	"github.com/gabeefranco/gonotes-api/internal/domain"
)

type SqlUsersRepository struct {
	DB *sql.DB
}

func NewSqlUsersRepository(db *sql.DB) *SqlUsersRepository {
	return &SqlUsersRepository{
		DB: db,
	}
}

func (r SqlUsersRepository) GetByID(ctx context.Context, id int64) (*domain.User, error) {
	const query = `
		SELECT id, email, password_hash
		FROM users
		WHERE id = $1
		LIMIT 1
	`

	row := r.DB.QueryRowContext(ctx, query, id)

	var u domain.User
	err := row.Scan(&u.ID, &u.Email, &u.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &u, nil

}

func (r SqlUsersRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	const query = `
		SELECT id, email, password_hash
		FROM users
		WHERE email = $1
		LIMIT 1
	`

	row := r.DB.QueryRowContext(ctx, query, email)

	var u domain.User
	err := row.Scan(&u.ID, &u.Email, &u.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &u, err
}

func (r SqlUsersRepository) Create(ctx context.Context, u *domain.User) error {
	const query = `
		INSERT INTO users (email, password_hash)
		VALUES ($1, $2)
		RETURNING id
	`

	err := r.DB.QueryRowContext(ctx, query,
		u.Email,
		u.Password,
	).Scan(&u.ID)

	if err != nil {
		return err
	}

	return nil
}
