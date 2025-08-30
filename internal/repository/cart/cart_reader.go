package cart

import (
	"context"
	"database/sql"
	"errors"

	"github.com/kucingscript/go-grpc-ecommerce-be/internal/model"
)

func (r *cartRepository) GetCartByProductAndUserID(ctx context.Context, productID, userID string) (*model.UserCart, error) {
	query := `SELECT id, product_id, user_id, quantity, created_at, created_by, updated_at, updated_by 
			FROM user_cart WHERE product_id = $1 AND user_id = $2`

	row := r.db.QueryRowContext(ctx, query, productID, userID)
	if row.Err() != nil {
		return nil, row.Err()
	}

	var cart model.UserCart
	err := row.Scan(
		&cart.ID,
		&cart.ProductID,
		&cart.UserID,
		&cart.Quantity,
		&cart.CreatedAt,
		&cart.CreatedBy,
		&cart.UpdatedAt,
		&cart.UpdatedBy,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &cart, nil
}

func (r *cartRepository) GetListCart(ctx context.Context, userID string) ([]*model.UserCart, error) {
	query := `SELECT uc.id, uc.product_id, uc.user_id, uc.quantity, uc.created_at, uc.created_by, uc.updated_at, uc.updated_by,
			p.id, p.name, p.price, p.image_file_name
			FROM user_cart uc 
			JOIN product p ON uc.product_id = p.id
			WHERE uc.user_id = $1
			AND p.is_deleted IS false`

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}

	var carts []*model.UserCart = make([]*model.UserCart, 0)
	for rows.Next() {
		var cart model.UserCart
		cart.Product = &model.Product{}

		err = rows.Scan(
			&cart.ID,
			&cart.ProductID,
			&cart.UserID,
			&cart.Quantity,
			&cart.CreatedAt,
			&cart.CreatedBy,
			&cart.UpdatedAt,
			&cart.UpdatedBy,
			&cart.Product.ID,
			&cart.Product.Name,
			&cart.Product.Price,
			&cart.Product.ImageFileName,
		)

		if err != nil {
			return nil, err
		}

		carts = append(carts, &cart)
	}

	return carts, nil
}

func (r *cartRepository) GetCartByID(ctx context.Context, id string) (*model.UserCart, error) {
	query := `SELECT id, product_id, user_id, quantity, created_at, created_by, updated_at, updated_by 
			FROM user_cart WHERE id = $1`

	row := r.db.QueryRowContext(ctx, query, id)
	if row.Err() != nil {
		return nil, row.Err()
	}

	var cart model.UserCart
	err := row.Scan(
		&cart.ID,
		&cart.ProductID,
		&cart.UserID,
		&cart.Quantity,
		&cart.CreatedAt,
		&cart.CreatedBy,
		&cart.UpdatedAt,
		&cart.UpdatedBy,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &cart, nil
}
