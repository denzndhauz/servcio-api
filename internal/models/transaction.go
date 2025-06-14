package models

import (
	"time"

	"github.com/google/uuid"
)

type PaymentMethod string

const (
	PaymentMethodCash       PaymentMethod = "CASH"
	PaymentMethodCreditCard PaymentMethod = "CREDIT_CARD"
	PaymentMethodDebitCard  PaymentMethod = "DEBIT_CARD"
	PaymentMethodPaypal     PaymentMethod = "PAYPAL"
	PaymentMethodStripe     PaymentMethod = "STRIPE"
	PaymentMethodGcash      PaymentMethod = "GCASH"
	PaymentMethodPaymaya    PaymentMethod = "PAYMAYA"
)

type TransactionStatus string

const (
	TransactionStatusPending   TransactionStatus = "PENDING"
	TransactionStatusCompleted TransactionStatus = "COMPLETED"
	TransactionStatusFailed    TransactionStatus = "FAILED"
	TransactionStatusRefunded  TransactionStatus = "REFUNDED"
)

type Transaction struct {
	ID            uuid.UUID         `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	BookingID     uuid.UUID         `gorm:"type:uuid;not null" json:"bookingId"`
	Booking       Booking           `gorm:"foreignKey:BookingID" json:"booking"`
	Amount        float64           `gorm:"not null" json:"amount"`
	PaymentMethod PaymentMethod     `gorm:"type:varchar(20);not null" json:"paymentMethod"`
	Status        TransactionStatus `gorm:"type:varchar(20);default:'PENDING'" json:"status"`
	TransactionID *string           `json:"transactionId"` // External payment processor ID
	ProcessedAt   *time.Time        `json:"processedAt"`
	CreatedAt     time.Time         `json:"createdAt"`
	UpdatedAt     time.Time         `json:"updatedAt"`
}
