package models

import (
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	FirstName string    `gorm:"not null" json:"firstName"`
	LastName  string    `gorm:"not null" json:"lastName"`
	Email     string    `gorm:"unique;not null" json:"email"`
	Phone     *string   `json:"phone"`
	Bookings  []Booking `gorm:"foreignKey:CustomerID" json:"bookings"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
