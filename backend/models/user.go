package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Username	string		`gorm:"primaryKey"`
	CreatedAt	time.Time
	UpdatedAt	time.Time
	Submission	[]Submission
}