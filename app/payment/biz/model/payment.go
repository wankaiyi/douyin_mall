package model

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type Payment struct {
	ID        int64   `gorm:"primary_key" json:"id"`
	OrderID   int64   `json:"order_id"`
	Amount    float32 `json:"amount"`
	CreatedAt string  `json:"created_at"`
}

func Create(db *gorm.DB, ctx context.Context, paymentLog *Payment) error {
	paymentLog.CreatedAt = time.Now().Format(time.RFC3339)
	result := db.WithContext(ctx).Create(paymentLog)
	return result.Error
}
