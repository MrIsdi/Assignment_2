package models

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	ID          uint
	ItemCode    int32
	Description string
	Quantity    int32
	OrderID     uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
