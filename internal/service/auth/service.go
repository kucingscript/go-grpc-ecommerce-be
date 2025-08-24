package auth

import (
	"context"

	authRepo "github.com/kucingscript/go-grpc-ecommerce-be/internal/repository/auth"
	"github.com/kucingscript/go-grpc-ecommerce-be/pb/auth"
	"github.com/microcosm-cc/bluemonday"
)

type IAuthService interface {
	Register(ctx context.Context, request *auth.RegisterRequest) (*auth.RegisterResponse, error)
}

type authService struct {
	authRepository authRepo.IAuthRepository
	htmlSanitizer  *bluemonday.Policy
}

func NewAuthService(authRepository authRepo.IAuthRepository) IAuthService {
	sanitizer := bluemonday.UGCPolicy()

	return &authService{
		authRepository: authRepository,
		htmlSanitizer:  sanitizer,
	}
}
