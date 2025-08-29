package product

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/kucingscript/go-grpc-ecommerce-be/internal/model"
	"github.com/kucingscript/go-grpc-ecommerce-be/pb/common"
)

func (r *productRepository) GetProductByID(ctx context.Context, id string) (*model.Product, error) {
	query := `SELECT id, name, description, price, image_file_name
			FROM product WHERE id = $1
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

func (r *productRepository) GetProductsPagination(ctx context.Context, pagination *common.PaginationRequest) ([]*model.Product, *common.PaginationResponse, error) {
	row := r.db.QueryRowContext(ctx, `SELECT COUNT(*) FROM product WHERE is_deleted IS false`)
	if row.Err() != nil {
		return nil, nil, row.Err()
	}

	var totalCount int
	err := row.Scan(&totalCount)
	if err != nil {
		return nil, nil, err
	}

	query := `SELECT id, name, description, price, image_file_name
			FROM product WHERE is_deleted IS false
			ORDER BY created_at DESC
			LIMIT $1 OFFSET $2`

	offset := (pagination.CurrentPage - 1) * pagination.ItemPerPage
	totalPages := (totalCount + int(pagination.ItemPerPage)) / int(pagination.ItemPerPage)

	rows, err := r.db.QueryContext(ctx, query, pagination.ItemPerPage, offset)
	if err != nil {
		return nil, nil, err
	}

	var products []*model.Product = make([]*model.Product, 0)
	for rows.Next() {
		var product model.Product
		err = rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.ImageFileName,
		)
		if err != nil {
			return nil, nil, err
		}

		products = append(products, &product)
	}

	paginationResponse := &common.PaginationResponse{
		CurrentPage:    pagination.CurrentPage,
		ItemPerPage:    pagination.ItemPerPage,
		TotalItemCount: int32(totalCount),
		TotalPageCount: int32(totalPages),
	}
	return products, paginationResponse, nil
}

func (r *productRepository) GetProductsPaginationAdmin(ctx context.Context, pagination *common.PaginationRequest) ([]*model.Product, *common.PaginationResponse, error) {
	row := r.db.QueryRowContext(ctx, `SELECT COUNT(*) FROM product WHERE is_deleted IS false`)
	if row.Err() != nil {
		return nil, nil, row.Err()
	}

	var totalCount int
	err := row.Scan(&totalCount)
	if err != nil {
		return nil, nil, err
	}

	allowedSorts := map[string]bool{
		"name":        true,
		"description": true,
		"price":       true,
	}

	orderQuery := "ORDER BY created_at DESC"
	if pagination.Sort != nil && allowedSorts[pagination.Sort.Field] {
		direction := "asc"
		if pagination.Sort.Direction == "desc" {
			direction = "desc"
		}

		orderQuery = fmt.Sprintf("ORDER BY %s %s", pagination.Sort.Field, direction)
	}

	query := fmt.Sprintf(`SELECT id, name, description, price, image_file_name
			FROM product WHERE is_deleted IS false
			%s
			LIMIT $1 OFFSET $2`, orderQuery)

	offset := (pagination.CurrentPage - 1) * pagination.ItemPerPage
	totalPages := (totalCount + int(pagination.ItemPerPage)) / int(pagination.ItemPerPage)

	rows, err := r.db.QueryContext(ctx, query, pagination.ItemPerPage, offset)
	if err != nil {
		return nil, nil, err
	}

	var products []*model.Product = make([]*model.Product, 0)
	for rows.Next() {
		var product model.Product
		err = rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.ImageFileName,
		)
		if err != nil {
			return nil, nil, err
		}

		products = append(products, &product)
	}

	paginationResponse := &common.PaginationResponse{
		CurrentPage:    pagination.CurrentPage,
		ItemPerPage:    pagination.ItemPerPage,
		TotalItemCount: int32(totalCount),
		TotalPageCount: int32(totalPages),
	}
	return products, paginationResponse, nil
}

func (r *productRepository) GetProductHighlight(ctx context.Context) ([]*model.Product, error) {
	query := `SELECT id, name, description, price, image_file_name FROM product 
			WHERE is_deleted IS false 
			ORDER BY created_at DESC
			LIMIT 3`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var products []*model.Product = make([]*model.Product, 0)
	for rows.Next() {
		var product model.Product
		err = rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.ImageFileName,
		)
		if err != nil {
			return nil, err
		}

		products = append(products, &product)
	}

	return products, nil

}
