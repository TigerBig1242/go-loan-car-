package models

import (
	"time"

	"gorm.io/gorm"
)

type ModelsCar struct {
	ID              uint           `gorm:"primaryKey" json:"id"`
	BrandID         uint           `json:"brand_id"`
	Brand           Brand          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"brand"` // ยังไม่ได้ข้อสรุปที่แน่ชัด
	ModelCode       string         `json:"model_code"`
	ModelName       string         `json:"model_name"`
	YearStart       int32          `json:"year_start"`
	YearEnd         int32          `json:"year_end"`
	BodyType        string         `json:"body_type"`
	EngineType      string         `json:"engine_type"`
	EngineSize      int32          `json:"engine_size"`
	FuelConsumption string         `json:"fuel_consumption"`
	Transmission    string         `json:"transmission"`
	Generation      string         `json:"generation"`
	Created_at      time.Time      `json:"created_at"`
	Updated_at      time.Time      `json:"updated_at"`
	Deleted_at      gorm.DeletedAt `json:"deleted_at"`
}
