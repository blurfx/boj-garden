package models

import (
	"gorm.io/gorm"
	"time"
)


type Submission struct {
	gorm.Model
	ID			uint		`gorm:"primaryKey"`
	UserID		uint		`gorm:"index"`
	Result		string		`gorm:"index"`
	CreatedAt	time.Time
	UpdatedAt	time.Time
}