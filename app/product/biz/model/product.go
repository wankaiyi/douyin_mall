package model

import "time"

type Product struct {
	ID          int64     `gorm:"primary_key" json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Picture     string    `json:"picture"`
	Price       string    `json:"price"`
	Stock       int64     `json:"stock"`
	Sale        int64     `json:"sale"`
	PublicState int64     `json:"public_state"`
	LockStock   int64     `json:"lock_stock"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
