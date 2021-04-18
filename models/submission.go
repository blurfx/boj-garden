package models

import (
	"gorm.io/gorm"
	"time"
)

type Submission struct {
	gorm.Model
	ID			uint		`gorm:"primaryKey"`
	ProblemID	uint		`gorm:"index;not null"`
	UserID		uint		`gorm:"index;not null"`
	Language	string		`gorm:"not null"`
	Result		string		`gorm:"index;not null"`
	Date		time.Time	`gorm:"not null"`
	CreatedAt	time.Time
	UpdatedAt	time.Time
}
