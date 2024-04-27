package repository

import (
	"context"
	"database/sql"

	"github.com/nandonyata/10xers-test/entity"
)

type ProductRepository struct {
	Database *sql.DB
}

func (r *ProductRepository) Insert(ctx context.Context, in entity.Product) (int, error) {
	return 1, nil
}
