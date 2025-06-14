package models

import (
	"time"

	"github.com/google/uuid"
)

type ServiceCategory struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Name        string    `gorm:"not null" json:"name"`
	Description *string   `json:"description"`
	Services    []Service `gorm:"foreignKey:CategoryID" json:"services"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
