package product

import (
	"context"
	"database/sql"

	"github.com/kucingscript/go-grpc-ecommerce-be/internal/model"
)

type IProductRepository interface {
	CreateNewProduct(ctx context.Context, product *model.Product) error
	GetProductByID(ctx context.Context, id string) (*model.Product, error)
}

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) IProductRepository {
	return &productRepository{
		db: db,
	}
}
