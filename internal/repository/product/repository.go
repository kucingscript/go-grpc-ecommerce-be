package product

import (
	"context"
	"database/sql"
	"time"

	"github.com/kucingscript/go-grpc-ecommerce-be/internal/model"
	"github.com/kucingscript/go-grpc-ecommerce-be/pb/common"
)

type IProductRepository interface {
	CreateNewProduct(ctx context.Context, product *model.Product) error
	GetProductByID(ctx context.Context, id string) (*model.Product, error)
	UpdateProduct(ctx context.Context, product *model.Product) error
	DeleteProduct(ctx context.Context, id string, deletedAt time.Time, deletedBy string) error
	GetProductsPagination(ctx context.Context, pagination *common.PaginationRequest) ([]*model.Product, *common.PaginationResponse, error)
	GetProductsPaginationAdmin(ctx context.Context, pagination *common.PaginationRequest) ([]*model.Product, *common.PaginationResponse, error)
	GetProductHighlight(ctx context.Context) ([]*model.Product, error)
}

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) IProductRepository {
	return &productRepository{
		db: db,
	}
}
