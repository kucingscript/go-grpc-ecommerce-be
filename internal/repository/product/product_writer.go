package product

import (
	"context"

	"github.com/kucingscript/go-grpc-ecommerce-be/internal/model"
)

func (r *productRepository) CreateNewProduct(ctx context.Context, product *model.Product) error {
	query := `INSERT INTO "product" 
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
