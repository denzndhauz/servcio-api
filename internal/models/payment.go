package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Payment struct {
	ID            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	BookingID     uuid.UUID `json:"bookingId"`
	Amount        float64   `json:"amount"`
	TransactionID *string   `json:"transactionId"`
}

// BeforeCreate hook to generate UUID before saving to the database
func (s *Payment) BeforeCreate(tx *gorm.DB) (err error) {
	s.ID = uuid.New()
	return
}
