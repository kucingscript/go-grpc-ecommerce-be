package cart

import (
	"context"
	"database/sql"

	"github.com/kucingscript/go-grpc-ecommerce-be/internal/model"
)

type ICartRepository interface {
	GetCartByProductAndUserID(ctx context.Context, productID, userID string) (*model.UserCart, error)
	CreateNewCart(ctx context.Context, cart *model.UserCart) error
	UpdateCart(ctx context.Context, cart *model.UserCart) error
	GetListCart(ctx context.Context, userID string) ([]*model.UserCart, error)
	GetCartByID(ctx context.Context, cartID string) (*model.UserCart, error)
	DeleteCart(ctx context.Context, cartID string) error
}

type cartRepository struct {
	db *sql.DB
}

func NewCartRepository(db *sql.DB) ICartRepository {
	return &cartRepository{
		db: db,
	}
}
