package model

import "time"

type UserCart struct {
	ID        string     `json:"id"`
	ProductID string     `json:"product_id"`
	UserID    string     `json:"user_id"`
	Quantity  int64      `json:"quantity"`
	CreatedAt time.Time  `json:"created_at"`
	CreatedBy *string    `json:"created_by,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	UpdatedBy *string    `json:"updated_by,omitempty"`

	Product *Product
}
