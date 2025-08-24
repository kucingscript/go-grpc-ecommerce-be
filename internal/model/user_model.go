package model

import "time"

const (
	USER_ROLE_ADMIN    = "admin"
	USER_ROLE_CUSTOMER = "customer"
)

type UserRole struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Code      string     `json:"code"`
	CreatedAt time.Time  `json:"created_at"`
	CreatedBy *string    `json:"created_by,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	UpdatedBy *string    `json:"updated_by,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	DeletedBy *string    `json:"deleted_by,omitempty"`
	IsDeleted bool       `json:"is_deleted"`
}

type User struct {
	ID        string     `json:"id"`
	FullName  string     `json:"full_name"`
	Email     string     `json:"email"`
	Password  string     `json:"-"`
	RoleCode  string     `json:"role_code"`
	CreatedAt time.Time  `json:"created_at"`
	CreatedBy *string    `json:"created_by,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	UpdatedBy *string    `json:"updated_by,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	DeletedBy *string    `json:"deleted_by,omitempty"`
	IsDeleted *bool      `json:"is_deleted,omitempty"`
}
