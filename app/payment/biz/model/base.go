package model

import "time"

type Base struct {
	ID        int64 `gorm:"primarykey"`
	CreatedAt time.Time
}
