package auth

import (
	"context"

	jwtModel "github.com/kucingscript/go-grpc-ecommerce-be/internal/model/jwt"
	"github.com/kucingscript/go-grpc-ecommerce-be/internal/utils"
	"github.com/kucingscript/go-grpc-ecommerce-be/pb/auth"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (as *authService) GetProfile(ctx context.Context, request *auth.GetProfileRequest) (*auth.GetProfileResponse, error) {
	claims, err := jwtModel.GetClaimsFromContext(ctx)
	if err != nil {
		return nil, err
	}

	user, err := as.authRepository.GetUserByEmail(ctx, claims.Email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return &auth.GetProfileResponse{
			BaseResponse: utils.BadRequestResponse("User not found"),
		}, nil
	}

	return &auth.GetProfileResponse{
		BaseResponse: utils.SuccessResponse("Get profile successfully"),
		UserId:       claims.Subject,
		FullName:     claims.FullName,
		Email:        claims.Email,
		RoleCode:     claims.Role,
		MemberSince:  timestamppb.New(user.CreatedAt),
	}, nil
}
