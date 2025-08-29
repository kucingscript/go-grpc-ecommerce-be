package product

import (
	"context"
	"time"

	"github.com/kucingscript/go-grpc-ecommerce-be/internal/model"
)

func (r *productRepository) CreateNewProduct(ctx context.Context, product *model.Product) error {
	query := `INSERT INTO product 
			(id, name, description, price, image_file_name, 
			created_at, created_by, updated_at, updated_by, deleted_at, deleted_by, is_deleted) 
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`

	_, err := r.db.ExecContext(ctx, query,
		product.ID, product.Name, product.Description, product.Price, product.ImageFileName,
		product.CreatedAt, product.CreatedBy, product.UpdatedAt, product.UpdatedBy, product.DeletedAt, product.DeletedBy, product.IsDeleted)

	if err != nil {
		return err
	}

	return nil
}

func (r *productRepository) UpdateProduct(ctx context.Context, product *model.Product) error {
	query := `UPDATE product SET name = $1, description = $2, price = $3, 
			image_file_name = $4, updated_at = $5, updated_by = $6 WHERE id = $7`

	_, err := r.db.ExecContext(ctx, query, product.Name, product.Description, product.Price, product.ImageFileName, product.UpdatedAt, product.UpdatedBy, product.ID)

	if err != nil {
		return err
	}

	return nil
}

func (r *productRepository) DeleteProduct(ctx context.Context, id string, deletedAt time.Time, deletedBy string) error {
	query := `UPDATE product SET is_deleted = true, deleted_at = $1, deleted_by = $2 WHERE id = $3`

	_, err := r.db.ExecContext(ctx, query, deletedAt, deletedBy, id)

	if err != nil {
		return err
	}

	return nil
}
