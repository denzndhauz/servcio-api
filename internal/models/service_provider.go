package models

import (
	"time"

	"github.com/google/uuid"
)

type ServiceProvider struct {
	ID          uuid.UUID  `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	FirstName   string     `gorm:"not null" json:"firstName"`
	LastName    string     `gorm:"not null" json:"lastName"`
	Email       string     `gorm:"unique;not null" json:"email"`
	Phone       *string    `json:"phone"`
	Specialties []Service  `gorm:"many2many:service_provider_services;" json:"specialties"`
	Schedules   []Schedule `gorm:"foreignKey:ProviderID" json:"schedules"`
	Bookings    []Booking  `gorm:"foreignKey:ProviderID" json:"bookings"`
	IsActive    bool       `gorm:"default:true" json:"isActive"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
}
