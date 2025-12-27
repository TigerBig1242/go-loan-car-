package models

import (
	"time"

	"gorm.io/gorm"
)

type Brand struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	BrandName   string         `json:"brand_name"`
	Country     string         `json:"country"`
	Description string         `json:"description"`
	CarModels   []ModelsCar    `gorm:"foreignKey:BrandID" json:"car_models,omitempty"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}
