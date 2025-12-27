package models

import (
	"time"

	"gorm.io/gorm"
)

type Detail struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Brand_id    uint           `gorm:"foreignKey" json:"brand_id"`
	Model_id    uint           `gorm:"foreignKey" json:"model_id"`
	Detail_code string         `json:"detail_code"`
	Price       int            `json:"price"`
	Created_at  time.Time      `json:"created_at"`
	Updated_at  time.Time      `json:"updated_at"`
	Deleted_at  gorm.DeletedAt `json:"deleted_at"`
}
