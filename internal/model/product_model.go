package model

import "time"

type Product struct {
	ID            string     `json:"id"`
	Name          string     `json:"name"`
	Description   string     `json:"description"`
	Price         float64    `json:"price"`
	ImageFileName string     `json:"image_file_name"`
	CreatedAt     time.Time  `json:"created_at"`
	CreatedBy     *string    `json:"created_by,omitempty"`
	UpdatedAt     *time.Time `json:"updated_at,omitempty"`
	UpdatedBy     *string    `json:"updated_by,omitempty"`
	DeletedAt     *time.Time `json:"deleted_at,omitempty"`
	DeletedBy     *string    `json:"deleted_by,omitempty"`
	IsDeleted     bool       `json:"is_deleted"`
}
