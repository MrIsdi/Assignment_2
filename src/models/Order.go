package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID           uint
	CustomerName string
	Items        []Item
	OrderedAt    time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
