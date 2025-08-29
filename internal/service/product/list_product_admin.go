package product

import (
	"context"
	"fmt"

	"github.com/kucingscript/go-grpc-ecommerce-be/internal/model"
	jwtModel "github.com/kucingscript/go-grpc-ecommerce-be/internal/model/jwt"
	"github.com/kucingscript/go-grpc-ecommerce-be/internal/utils"
	"github.com/kucingscript/go-grpc-ecommerce-be/pb/product"
)

func (s *productService) ListProductAdmin(ctx context.Context, request *product.ListProductAdminRequest) (*product.ListProductAdminResponse, error) {
	claims, err := jwtModel.GetClaimsFromContext(ctx)
	if err != nil {
		return nil, err
	}

	if claims.Role != model.USER_ROLE_ADMIN {
		return nil, utils.UnauthenticatedResponse()
	}

	products, paginationResponse, err := s.productRepository.GetProductsPaginationAdmin(ctx, request.Pagination)
	if err != nil {
		return nil, err
	}

	var data []*product.ListProductAdminResponseItem = make([]*product.ListProductAdminResponseItem, 0)
	for _, prod := range products {
		data = append(data, &product.ListProductAdminResponseItem{
			Id:          prod.ID,
			Name:        prod.Name,
			Description: prod.Description,
			Price:       prod.Price,
			ImageUrl:    fmt.Sprintf("%s/product/%s", s.StorageServiceUrl, prod.ImageFileName),
		})
	}

	return &product.ListProductAdminResponse{
		BaseResponse: utils.SuccessResponse("Get product admin successfully"),
		Pagination:   paginationResponse,
		Data:         data,
	}, nil
}
