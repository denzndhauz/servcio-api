package models

import (
	"time"

	"github.com/google/uuid"
)

type Schedule struct {
	ID         uuid.UUID       `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	ProviderID uuid.UUID       `gorm:"type:uuid;not null" json:"providerId"`
	Provider   ServiceProvider `gorm:"foreignKey:ProviderID" json:"provider"`
	DayOfWeek  int             `gorm:"not null" json:"dayOfWeek"` // 0 = Sunday, 1 = Monday, etc.
	StartTime  string          `gorm:"not null" json:"startTime"` // HH:MM format
	EndTime    string          `gorm:"not null" json:"endTime"`   // HH:MM format
	IsActive   bool            `gorm:"default:true" json:"isActive"`
	CreatedAt  time.Time       `json:"createdAt"`
	UpdatedAt  time.Time       `json:"updatedAt"`
}
