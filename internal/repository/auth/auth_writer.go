package auth

import (
	"context"

	"github.com/kucingscript/go-grpc-ecommerce-be/internal/model"
)

func (r *authRepository) CreateUser(ctx context.Context, user *model.User) error {
	query := `INSERT INTO "user" (
			id, full_name, email, password, role_code, 
			created_at, created_by, updated_at, updated_by, 
			deleted_at, deleted_by, is_deleted
			) 
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`

	_, err := r.db.ExecContext(ctx, query,
		user.ID, user.FullName, user.Email, user.Password, user.RoleCode,
		user.CreatedAt, user.CreatedBy, user.UpdatedAt, user.UpdatedBy,
		user.DeletedAt, user.DeletedBy, user.IsDeleted)

	if err != nil {
		return err
	}

	return nil
}
