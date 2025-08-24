package auth

import (
	"context"
	"time"

	jwtModel "github.com/kucingscript/go-grpc-ecommerce-be/internal/model/jwt"
	"github.com/kucingscript/go-grpc-ecommerce-be/internal/utils"
	"github.com/kucingscript/go-grpc-ecommerce-be/pb/auth"
)

func (as *authService) Logout(ctx context.Context, request *auth.LogoutRequest) (*auth.LogoutResponse, error) {
	jwtToken, err := jwtModel.ParseTokenFromContext(ctx)
	if err != nil {
		return nil, err
	}

	tokenClaims, err := jwtModel.GetClaimsFromToken(jwtToken)
	if err != nil {
		return nil, err
	}

	as.cacheService.Set(jwtToken, "", time.Duration(tokenClaims.ExpiresAt.Time.Unix()-time.Now().Unix())*time.Second)

	return &auth.LogoutResponse{
		BaseResponse: utils.SuccessResponse("Logout successfully"),
	}, nil
}
