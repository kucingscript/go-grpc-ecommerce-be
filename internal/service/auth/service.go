package auth

import (
	"context"

	authRepo "github.com/kucingscript/go-grpc-ecommerce-be/internal/repository/auth"
	"github.com/kucingscript/go-grpc-ecommerce-be/pb/auth"
	"github.com/microcosm-cc/bluemonday"
	gocache "github.com/patrickmn/go-cache"
)

type IAuthService interface {
	Register(ctx context.Context, request *auth.RegisterRequest) (*auth.RegisterResponse, error)
	Login(ctx context.Context, request *auth.LoginRequest) (*auth.LoginResponse, error)
	Logout(ctx context.Context, request *auth.LogoutRequest) (*auth.LogoutResponse, error)
	ChangePassword(ctx context.Context, request *auth.ChangePasswordRequest) (*auth.ChangePasswordResponse, error)
	GetProfile(ctx context.Context, request *auth.GetProfileRequest) (*auth.GetProfileResponse, error)
}

type authService struct {
	authRepository authRepo.IAuthRepository
	htmlSanitizer  *bluemonday.Policy
	jwtSecret      string
	cacheService   *gocache.Cache
}

func NewAuthService(authRepository authRepo.IAuthRepository, jwtSecret string, cacheService *gocache.Cache) IAuthService {
	sanitizer := bluemonday.UGCPolicy()

	return &authService{
		authRepository: authRepository,
		htmlSanitizer:  sanitizer,
		jwtSecret:      jwtSecret,
		cacheService:   cacheService,
	}
}
