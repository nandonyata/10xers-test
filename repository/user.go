package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/nandonyata/10xers-test/entity"
)

type UserRepository struct {
	Database *sql.DB
}

func (r *UserRepository) Insert(ctx context.Context, in entity.User) (int, error) {
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

func (r *UserRepository) FindOne(ctx context.Context, in entity.User) (entity.User, error) {
	query := `
		SELECT id, name, role, email, password
		FROM "user"
		WHERE email = $1
	`

	var user entity.User

	err := r.Database.QueryRowContext(ctx, query, in.Email).Scan(&user.Id, &user.Name, &user.Role, &user.Email, &user.Password)
	if err != nil {
		return user, errors.New("invalid email/password")
	}

	return user, nil
}

func (r *UserRepository) FindById(ctx context.Context, in entity.User) (entity.User, error) {
	query := `
		SELECT id, name, role, email, password
		FROM "user"
		WHERE id = $1
	`

	var user entity.User

	err := r.Database.QueryRowContext(ctx, query, in.Id).Scan(&user.Id, &user.Name, &user.Role, &user.Email, &user.Password)
	if err != nil {
		return user, err
	}

	return user, nil
}
