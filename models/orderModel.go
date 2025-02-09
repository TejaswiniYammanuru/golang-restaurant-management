package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/TejaswiniYammanuru/golang-restaurant-management/database"
)

type Order struct {
	ID         int         `json:"id" gorm:"primaryKey"`
	OrderDate  time.Time   `json:"order_date"`
	TableID    int         `json:"table_id" gorm:"not null;index"`
	UserID     int         `json:"user_id" gorm:"not null;index"`
	CreatedAt  time.Time   `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time   `json:"updated_at" gorm:"autoUpdateTime"`
	OrderItems []OrderItem `json:"order_items" gorm:"foreignKey:OrderID"`
	Invoice    Invoice     `json:"invoice" gorm:"foreignKey:OrderID"`
}

func GetAllOrders() ([]Order, error) {
	var orders []Order

	query := "SELECT id, order_date, table_id, user_id, created_at, updated_at FROM orders"
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var order Order
		err := rows.Scan(&order.ID, &order.OrderDate, &order.TableID, &order.UserID, &order.CreatedAt, &order.UpdatedAt)
		if err != nil {
			return nil, err
		}

		// Fetch OrderItems manually
		order.OrderItems, err = GetOrderItemsByOrder(order.ID)
		if err != nil {
			return nil, err
		}

		order.Invoice, _ = GetInvoiceByOrderID(order.ID)

		orders = append(orders, order)
	}

	return orders, nil
}

func GetOrderByID(id int) (*Order, error) {
	var order Order

	query := "SELECT id, order_date, table_id, user_id, created_at, updated_at FROM orders WHERE id=$1"
	err := database.DB.QueryRow(query, id).Scan(
		&order.ID, &order.OrderDate, &order.TableID, &order.UserID, &order.CreatedAt, &order.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("order not found")
		}
		return nil, err
	}

	// Fetch OrderItems manually
	order.OrderItems, err = GetOrderItemsByOrder(order.ID)
	if err != nil {
		return nil, err
	}

	order.Invoice, _ = GetInvoiceByOrderID(order.ID)

	return &order, nil
}

func UpdateOrder(order *Order) error {
	query := `UPDATE orders 
              SET order_date=$1, table_id=$2, user_id=$3, updated_at=$4 
              WHERE id=$5 RETURNING id, order_date, table_id, user_id, created_at, updated_at`

	// Execute the query
	err := database.DB.QueryRow(query, order.OrderDate, order.TableID, order.UserID, time.Now(), order.ID).
		Scan(&order.ID, &order.OrderDate, &order.TableID, &order.UserID, &order.CreatedAt, &order.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("order with ID %d not found", order.ID)
		}
		return fmt.Errorf("unable to update order: %v", err)
	}

	return nil
}

func CreateOrder(order *Order) (*Order, error) {
	query := `INSERT INTO orders (order_date, table_id, user_id, created_at, updated_at) 
	          VALUES ($1, $2, $3, $4, $5) RETURNING id`

	err := database.DB.QueryRow(query, order.OrderDate, order.TableID, order.UserID, time.Now(), time.Now()).
		Scan(&order.ID)
	if err != nil {
		return nil, fmt.Errorf("unable to insert order: %v", err)
	}

	return order, nil
}

func GetOrdersByUserID(userID int) ([]Order, error) {
	var orders []Order

	query := "SELECT id, order_date, table_id, user_id, created_at, updated_at FROM orders WHERE user_id=$1"
	rows, err := database.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var order Order
		err := rows.Scan(&order.ID, &order.OrderDate, &order.TableID, &order.UserID, &order.CreatedAt, &order.UpdatedAt)
		if err != nil {
			return nil, err
		}

		order.OrderItems, err = GetOrderItemsByOrder(order.ID)
		if err != nil {
			return nil, err
		}

		order.Invoice, _ = GetInvoiceByOrderID(order.ID)

		orders = append(orders, order)
	}

	if len(orders) == 0 {
		return nil, fmt.Errorf("no orders found for user ID %d", userID)
	}

	return orders, nil
}
