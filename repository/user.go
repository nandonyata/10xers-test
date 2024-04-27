package repository

import (
	"context"
	"database/sql"

	"github.com/nandonyata/10xers-test/entity"
)

type UserRepository struct {
	Database *sql.DB
}

func (r *UserRepository) Register(ctx context.Context, in entity.User) (int, error) {
	query := `
		INSERT INTO "user" (name, role, email, password)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`
	var userID int

	err := r.Database.QueryRowContext(ctx, query, in.Name, in.Role, in.Email, in.Password).Scan(&userID)
	if err != nil {
		return 0, err
	}

	return userID, nil
}
