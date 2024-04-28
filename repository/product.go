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

func (r *ProductRepository) FindAll(ctx context.Context) ([]entity.Product, error) {
	query := `
		SELECT id, userId, title, type, price, stock
		FROM product
		ORDER BY id DESC
	`

	var products []entity.Product

	rows, err := r.Database.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var product entity.Product

		if err := rows.Scan(&product.Id, &product.UserId, &product.Title, &product.Type, &product.Price, &product.Stock); err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

func (r *ProductRepository) FindById(ctx context.Context, id int) (entity.Product, error) {
	query := `
		SELECT id, userId, title, type, price, stock
		FROM product
		WHERE id = $1
	`
	var product entity.Product

	err := r.Database.QueryRowContext(ctx, query, id).Scan(&product.Id, &product.UserId, &product.Title, &product.Type, &product.Price, &product.Stock)
	if err != nil {
		return product, err
	}

	return product, nil

}

func (r *ProductRepository) UpdateById(ctx context.Context, in entity.Product) (interface{}, error) {
	query := `
		UPDATE product
		SET title = $2, type = $3, price = $4, stock = $5
		WHERE id = $1
	`

	result, err := r.Database.ExecContext(ctx, query, in.Id, in.Title, in.Type, in.Price, in.Stock)
	if err != nil {
		return nil, err
	}

	return result, nil
}
