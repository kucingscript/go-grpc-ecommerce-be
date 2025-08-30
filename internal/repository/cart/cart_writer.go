package cart

import (
	"context"

	"github.com/kucingscript/go-grpc-ecommerce-be/internal/model"
)

func (r *cartRepository) CreateNewCart(ctx context.Context, cart *model.UserCart) error {
	query := `INSERT INTO user_cart(id, product_id, user_id, quantity, created_at, created_by, updated_at, updated_by)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err := r.db.ExecContext(ctx, query,
		cart.ID, cart.ProductID, cart.UserID, cart.Quantity, cart.CreatedAt, cart.CreatedBy, cart.UpdatedAt, cart.UpdatedBy)

	if err != nil {
		return err
	}

	return nil
}

func (r *cartRepository) UpdateCart(ctx context.Context, cart *model.UserCart) error {
	query := `UPDATE user_cart SET product_id = $1, user_id = $2, quantity = $3, updated_at = $4, updated_by = $5 WHERE id = $6`

	_, err := r.db.ExecContext(ctx, query,
		cart.ProductID, cart.UserID, cart.Quantity, cart.UpdatedAt, cart.UpdatedBy, cart.ID)

	if err != nil {
		return err
	}

	return nil
}

func (r *cartRepository) DeleteCart(ctx context.Context, cartID string) error {
	query := `DELETE FROM user_cart WHERE id = $1`

	_, err := r.db.ExecContext(ctx, query, cartID)

	if err != nil {
		return err
	}

	return nil
}
