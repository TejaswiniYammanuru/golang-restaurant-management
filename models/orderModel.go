package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/TejaswiniYammanuru/golang-restaurant-management/database"
)

type Order struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	OrderDate time.Time `json:"order_date"`
	TableID   int       `json:"table_id" gorm:"not null;index"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	
	// Table Table `json:"table" gorm:"foreignKey:TableID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}


func GetAllOrders() ([]Order, error) {
	var orders []Order

	query := "select * from orders"

	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var order Order
		err := rows.Scan(&order.ID, &order.OrderDate, &order.TableID, &order.CreatedAt, &order.UpdatedAt)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

func GetOrderByID(id int) (*Order, error) {
	var order Order

	query := "select * from orders where id=$1"

	err := database.DB.QueryRow(query, id).Scan(&order.ID, &order.OrderDate, &order.TableID, &order.CreatedAt, &order.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("order not found")
		}
		return nil, err
	}

	return &order, nil
}

func UpdateOrder(order *Order) error {
	query := `UPDATE orders 
              SET order_date=$1, table_id=$2, updated_at=$3 
              WHERE id=$4 RETURNING id, order_date, table_id, created_at, updated_at`

	// Execute the query
	err := database.DB.QueryRow(query, order.OrderDate, order.TableID, time.Now(), order.ID).
		Scan(&order.ID, &order.OrderDate, &order.TableID, &order.CreatedAt, &order.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("order with ID %d not found", order.ID)
		}
		return fmt.Errorf("unable to update order: %v", err)
	}

	return nil
}



func CreateOrder(order *Order) (*Order, error) {
	query := `INSERT INTO orders (order_date, table_id, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id`

    err := database.DB.QueryRow(query, order.OrderDate, order.TableID, time.Now(), time.Now()).Scan(&order.ID)
    if err!= nil {
        return nil, fmt.Errorf("unable to insert order: %v", err)
    }

    return order, nil
}
