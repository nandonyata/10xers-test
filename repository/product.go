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
	query := `
		INSERT INTO product (title, price, stock, type, userId)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`

	var productId int

	err := r.Database.QueryRow(query, in.Title, in.Price, in.Stock, in.Type, in.UserId).Scan(&productId)
	if err != nil {
		return 0, err
	}

	return productId, nil
}
