package model

import "time"

type Base struct {
	ID        int32 `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
