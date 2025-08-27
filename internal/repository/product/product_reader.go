package product

import (
	"context"
	"database/sql"
	"errors"

	"github.com/kucingscript/go-grpc-ecommerce-be/internal/model"
)

func (r *productRepository) GetProductByID(ctx context.Context, id string) (*model.Product, error) {
	query := `SELECT id, name, description, price, image_file_name
			FROM "product" WHERE id = $1
			AND is_deleted IS false`

	row := r.db.QueryRowContext(ctx, query, id)
	if row.Err() != nil {
		return nil, row.Err()
	}

	var product model.Product
	err := row.Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.ImageFileName,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &product, nil
}
