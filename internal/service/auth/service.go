package auth

import (
	"context"

	authRepo "github.com/kucingscript/go-grpc-ecommerce-be/internal/repository/auth"
	"github.com/kucingscript/go-grpc-ecommerce-be/pb/auth"
)

type IAuthService interface {
	Register(ctx context.Context, request *auth.RegisterRequest) (*auth.RegisterResponse, error)
}

type authService struct {
	authRepository authRepo.IAuthRepository
}

func NewAuthService(authRepository authRepo.IAuthRepository) IAuthService {
	return &authService{
		authRepository: authRepository,
	}
}
