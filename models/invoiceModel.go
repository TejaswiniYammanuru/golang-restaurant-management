package models

import (
	"database/sql"
	"errors"
	"time"

	"github.com/TejaswiniYammanuru/golang-restaurant-management/database"
)


type Invoice struct {
	ID             int       `json:"id" gorm:"primaryKey"`
	OrderID        int    `json:"order_id" validate:"required" gorm:"not null;index"`
	PaymentMethod  string    `json:"payment_method" validate:"eq=CARD|eq=CASH|"`
	PaymentStatus  string    `json:"payment_status" validate:"required,eq=PENDING|eq=PAID"`
	PaymentDueDate time.Time `json:"payment_due_date"`
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	
	// Order Order `json:"order" gorm:"foreignKey:OrderID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}


func GetInvoices() ([]Invoice, error) {
	query := "SELECT id,order_id,payment_method,payment_status,payment_due_date,created_at,updated_at FROM invoice"

	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var invoices []Invoice
	for rows.Next() {
		var invoice Invoice
		err := rows.Scan(&invoice.ID, &invoice.OrderID, &invoice.PaymentMethod, &invoice.PaymentStatus, &invoice.PaymentDueDate, &invoice.CreatedAt, &invoice.UpdatedAt)
		if err != nil {
			return nil, err
		}
		invoices = append(invoices, invoice)
	}
	return invoices, nil
}

func GetInvoiceByID(invoiceID int) (*Invoice, error) {
	query := "SELECT id,order_id,payment_method,payment_status,payment_due_date,created_at,updated_at FROM invoice WHERE id=$1"

	var invoice Invoice

	err := database.DB.QueryRow(query, invoiceID).Scan(
		&invoice.ID,
		&invoice.OrderID,
		&invoice.PaymentMethod,
		&invoice.PaymentStatus,
		&invoice.PaymentDueDate,
		&invoice.CreatedAt,
		&invoice.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("Invoice not found")
		}
		return nil, err
	}
	return &invoice, nil
}

func CreateInvoice(invoice *Invoice) error {
	query := "INSERT INTO invoice (order_id, payment_method, payment_status, payment_due_date, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"

	err := database.DB.QueryRow(query, invoice.OrderID, invoice.PaymentMethod, invoice.PaymentStatus, invoice.PaymentDueDate, invoice.CreatedAt, invoice.UpdatedAt).Scan(&invoice.ID)
	if err != nil {
		return err
	}
	return nil
}

func UpdateInvoice(invoice *Invoice) error {
    query := "UPDATE invoice SET order_id=$1, payment_method=$2, payment_status=$3, payment_due_date=$4, updated_at=$5 WHERE id=$6"

    _, err := database.DB.Exec(query, invoice.OrderID, invoice.PaymentMethod, invoice.PaymentStatus, invoice.PaymentDueDate, invoice.UpdatedAt, invoice.ID)
    return err
}



