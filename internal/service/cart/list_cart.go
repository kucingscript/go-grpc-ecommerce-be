package cart

import (
	"context"
	"fmt"

	"github.com/kucingscript/go-grpc-ecommerce-be/internal/model"
	jwtModel "github.com/kucingscript/go-grpc-ecommerce-be/internal/model/jwt"
	"github.com/kucingscript/go-grpc-ecommerce-be/internal/utils"
	"github.com/kucingscript/go-grpc-ecommerce-be/pb/cart"
)

func (s *cartService) ListCart(ctx context.Context, request *cart.ListCartRequest) (*cart.ListCartResponse, error) {
	claims, err := jwtModel.GetClaimsFromContext(ctx)
	if err != nil {
		return nil, err
	}

	if claims.Role != model.USER_ROLE_CUSTOMER && claims.Role != model.USER_ROLE_ADMIN {
		return nil, utils.UnauthenticatedResponse()
	}

	carts, err := s.cartRepository.GetListCart(ctx, claims.Subject)
	if err != nil {
		return nil, err
	}

	var items []*cart.ListCartResponseItem = make([]*cart.ListCartResponseItem, 0)
	for _, ct := range carts {
		item := cart.ListCartResponseItem{
			CartId:          ct.ID,
			ProductId:       ct.ProductID,
			ProductName:     ct.Product.Name,
			ProductPrice:    ct.Product.Price,
			Quantity:        ct.Quantity,
			ProductImageUrl: fmt.Sprintf("%s/product/%s", s.StorageServiceUrl, ct.Product.ImageFileName),
		}

		items = append(items, &item)
	}

	return &cart.ListCartResponse{
		BaseResponse: utils.SuccessResponse(),
		Items:        items,
	}, nil
}
