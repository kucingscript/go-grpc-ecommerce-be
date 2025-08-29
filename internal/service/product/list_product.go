package product

import (
	"context"
	"fmt"

	"github.com/kucingscript/go-grpc-ecommerce-be/internal/utils"
	"github.com/kucingscript/go-grpc-ecommerce-be/pb/product"
)

func (s *productService) ListProduct(ctx context.Context, request *product.ListProductRequest) (*product.ListProductResponse, error) {
	products, paginationResponse, err := s.productRepository.GetProductsPagination(ctx, request.Pagination)
	if err != nil {
		return nil, err
	}

	var data []*product.ListProductResponseItem = make([]*product.ListProductResponseItem, 0)
	for _, prod := range products {
		data = append(data, &product.ListProductResponseItem{
			Id:          prod.ID,
			Name:        prod.Name,
			Description: prod.Description,
			Price:       prod.Price,
			ImageUrl:    fmt.Sprintf("%s/product/%s", s.StorageServiceUrl, prod.ImageFileName),
		})
	}

	return &product.ListProductResponse{
		BaseResponse: utils.SuccessResponse("Get product successfully"),
		Pagination:   paginationResponse,
		Data:         data,
	}, nil
}
