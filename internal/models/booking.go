package models

import (
	"time"

	"github.com/google/uuid"
)

type BookingStatus string

const (
	BookingStatusPending    BookingStatus = "PENDING"
	BookingStatusConfirmed  BookingStatus = "CONFIRMED"
	BookingStatusInProgress BookingStatus = "IN_PROGRESS"
	BookingStatusCompleted  BookingStatus = "COMPLETED"
	BookingStatusCancelled  BookingStatus = "CANCELLED"
	BookingStatusNoShow     BookingStatus = "NO_SHOW"
)

type Booking struct {
	ID           uuid.UUID       `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	CustomerID   uuid.UUID       `gorm:"type:uuid;not null" json:"customerId"`
	Customer     Customer        `gorm:"foreignKey:CustomerID" json:"customer"`
	ServiceID    uuid.UUID       `gorm:"type:uuid;not null" json:"serviceId"`
	Service      Service         `gorm:"foreignKey:ServiceID" json:"service"`
	ProviderID   uuid.UUID       `gorm:"type:uuid;not null" json:"providerId"`
	Provider     ServiceProvider `gorm:"foreignKey:ProviderID" json:"provider"`
	ScheduledAt  time.Time       `gorm:"not null" json:"scheduledAt"`
	Status       BookingStatus   `gorm:"type:varchar(20);default:'PENDING'" json:"status"`
	TotalAmount  float64         `gorm:"not null" json:"totalAmount"`
	Transactions []Transaction   `gorm:"foreignKey:BookingID" json:"transactions"`
	Notes        *string         `json:"notes"`
	CreatedAt    time.Time       `json:"createdAt"`
	UpdatedAt    time.Time       `json:"updatedAt"`
}
