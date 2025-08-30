package cart

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/kucingscript/go-grpc-ecommerce-be/internal/model"
	jwtModel "github.com/kucingscript/go-grpc-ecommerce-be/internal/model/jwt"
	"github.com/kucingscript/go-grpc-ecommerce-be/internal/utils"
	"github.com/kucingscript/go-grpc-ecommerce-be/pb/cart"
)

func (s *cartService) AddProductToCart(ctx context.Context, request *cart.AddProductToCartRequest) (*cart.AddProductToCartResponse, error) {
	claims, err := jwtModel.GetClaimsFromContext(ctx)
	if err != nil {
		return nil, err
	}

	if claims.Role != model.USER_ROLE_CUSTOMER && claims.Role != model.USER_ROLE_ADMIN {
		return nil, utils.UnauthenticatedResponse()
	}

	productExits, err := s.productRepository.GetProductByID(ctx, request.ProductId)
	if err != nil {
		return nil, err
	}

	if productExits == nil {
		return &cart.AddProductToCartResponse{
			BaseResponse: utils.NotFoundResponse("Product not found"),
		}, nil
	}

	cartEntity, err := s.cartRepository.GetCartByProductAndUserID(ctx, request.ProductId, claims.Subject)
	if err != nil {
		return nil, err
	}

	if cartEntity != nil {
		now := time.Now()
		cartEntity.Quantity += 1
		cartEntity.UpdatedAt = &now
		cartEntity.UpdatedBy = &claims.FullName

		err = s.cartRepository.UpdateCart(ctx, cartEntity)
		if err != nil {
			return nil, err
		}

		return &cart.AddProductToCartResponse{
			BaseResponse: utils.SuccessResponse("Product added to cart successfully"),
			Id:           cartEntity.ID,
		}, nil
	}

	newCartEntity := model.UserCart{
		ID:        uuid.NewString(),
		ProductID: request.ProductId,
		UserID:    claims.Subject,
		Quantity:  1,
		CreatedAt: time.Now(),
		CreatedBy: &claims.FullName,
	}

	err = s.cartRepository.CreateNewCart(ctx, &newCartEntity)
	if err != nil {
		return nil, err
	}

	return &cart.AddProductToCartResponse{
		BaseResponse: utils.SuccessResponse("Product added to cart successfully"),
		Id:           newCartEntity.ID,
	}, nil
}
