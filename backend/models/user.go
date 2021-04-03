package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID			uint			`gorm:"primaryKey"`
	Username	string			`gorm:"unique"`
	CreatedAt	time.Time
	UpdatedAt	time.Time
	Submission	[]Submission
}