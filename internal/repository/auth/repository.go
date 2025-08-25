package auth

import (
	"context"
	"database/sql"

	"github.com/kucingscript/go-grpc-ecommerce-be/internal/model"
)

type IAuthRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	CreateUser(ctx context.Context, user *model.User) error
	UpdateUserPassword(ctx context.Context, userID, hashedNewPassword, updatedBy string) error
}

type authRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) IAuthRepository {
	return &authRepository{
		db: db,
	}
}
