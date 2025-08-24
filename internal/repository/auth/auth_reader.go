package auth

import (
	"context"
	"database/sql"
	"errors"

	"github.com/kucingscript/go-grpc-ecommerce-be/internal/model"
)

func (r *authRepository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	query := `SELECT id, email, password, full_name, role_code FROM "user" 
			WHERE email = $1 
			AND is_deleted IS false`

	row := r.db.QueryRowContext(ctx, query, email)
	if row.Err() != nil {
		return nil, row.Err()
	}

	var user model.User
	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.FullName,
		&user.RoleCode,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
