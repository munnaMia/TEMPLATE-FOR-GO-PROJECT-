package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/munnaMia/ahlan/internal/domain"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) domain.UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(ctx context.Context, u *domain.User) (*domain.User, error) {
	// SQL query execution...
	query := `
		INSERT INTO users (username, email, password_hash)
		VALUES (1$, 2$, 3$)
		RETURNING id, created_at, updated_at
	`

	err := r.db.QueryRowContext(
		ctx,
		query,
		u.Name,
		u.Email,
		u.Password,
	).Scan(
		&u.Id,
		&u.CreatedAt,
		&u.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create new user in postgres: %w", err)
	}

	return u, nil
}
