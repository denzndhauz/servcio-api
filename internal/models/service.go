package models

import (
	"time"

	"github.com/google/uuid"
)

type Service struct {
	ID          uuid.UUID         `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Name        string            `gorm:"not null" json:"name"`
	Description *string           `json:"description"`
	Duration    int               `gorm:"not null" json:"duration"` // in minutes
	Price       float64           `gorm:"not null" json:"price"`
	CategoryID  uuid.UUID         `gorm:"type:uuid;not null" json:"categoryId"`
	Categories  ServiceCategory   `gorm:"foreignKey:CategoryID" json:"categories"`
	Providers   []ServiceProvider `gorm:"many2many:service_provider_services;" json:"providers"`
	Bookings    []Booking         `gorm:"foreignKey:ServiceID" json:"bookings"`
	CreatedAt   time.Time         `json:"createdAt"`
	UpdatedAt   time.Time         `json:"updatedAt"`
}
