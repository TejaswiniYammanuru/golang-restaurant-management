package models

import "time"

type Invoice struct {
	ID             int    `json:"id"`
	InvoiceID      string `json:"invoice_id"`
	OrderID        string    `json:"order_id"`
	PaymentMethod  string `json:"payment_method" validate:"eq=CARD|eq=CASH|eq="`
	PaymentStatus  string `json:"payment_status" validate:"required,eq=PENDING|eq=PAID"`
	PaymentDueDate time.Time `json:"payment_due_date"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`

}
