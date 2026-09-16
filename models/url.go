package models

import (
	"gorm.io/gorm"
)

type Url struct {
	gorm.Model
	Code        string `gorm:"uniqueIndex;not null"`
	OriginalURL string `gorm:"nut null"`
	ClickCount  int    `gorm:"default:0"`
}
