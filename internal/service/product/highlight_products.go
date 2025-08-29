package product

import (
	"context"
	"fmt"

	"github.com/kucingscript/go-grpc-ecommerce-be/internal/utils"
	"github.com/kucingscript/go-grpc-ecommerce-be/pb/product"
)

func (s *productService) HighlightProducts(ctx context.Context, request *product.HighlightProductsRequest) (*product.HighlightProductsResponse, error) {
	products, err := s.productRepository.GetProductHighlight(ctx)
	if err != nil {
		return nil, err
	}

	var data []*product.HighlightProductsResponseItem = make([]*product.HighlightProductsResponseItem, 0)
	for _, prod := range products {
		data = append(data, &product.HighlightProductsResponseItem{
			Id:          prod.ID,
			Name:        prod.Name,
			Description: prod.Description,
			Price:       prod.Price,
			ImageUrl:    fmt.Sprintf("%s/product/%s", s.StorageServiceUrl, prod.ImageFileName),
		})
	}

	return &product.HighlightProductsResponse{
		BaseResponse: utils.SuccessResponse("Get highlight products successfully"),
		Data:         data,
	}, nil
}
