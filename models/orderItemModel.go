package models

import (
	"time"

	"github.com/TejaswiniYammanuru/golang-restaurant-management/database"
)

type OrderItem struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Quantity  string    `json:"quantity" validate:"required,eq=S|eq=M|eq=L"`
	UnitPrice float64   `json:"unit_price" validate:"required"`
	FoodID    int       `json:"food_id" validate:"required" gorm:"not null;index"`
	OrderID   int       `json:"order_id" validate:"required" gorm:"not null;index"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func GetOrderItems() (orderItems []OrderItem, err error) {
	query := "SELECT id, quantity, unit_price, food_id, order_id, created_at, updated_at FROM order_items"
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var orderItem OrderItem
		err := rows.Scan(&orderItem.ID, &orderItem.Quantity, &orderItem.UnitPrice, &orderItem.FoodID, &orderItem.OrderID, &orderItem.CreatedAt, &orderItem.UpdatedAt)
		if err != nil {
			return nil, err
		}
		orderItems = append(orderItems, orderItem)
	}
	return orderItems, nil

}

func GetOrderItem(id int) (orderItem OrderItem, err error) {
	query := "SELECT id, quantity, unit_price, food_id, order_id, created_at, updated_at FROM order_items WHERE id=$1"
	row := database.DB.QueryRow(query, id)
	err = row.Scan(&orderItem.ID, &orderItem.Quantity, &orderItem.UnitPrice, &orderItem.FoodID, &orderItem.OrderID, &orderItem.CreatedAt, &orderItem.UpdatedAt)
	if err != nil {
		return orderItem, err
	}
	return orderItem, nil

}

func CreateOrderItem(orderItem *OrderItem) error {
	query := "INSERT INTO order_items (quantity, unit_price, food_id, order_id) VALUES ($1, $2, $3, $4) RETURNING id"
	err := database.DB.QueryRow(query, orderItem.Quantity, orderItem.UnitPrice, orderItem.FoodID, orderItem.OrderID).Scan(&orderItem.ID)
	if err != nil {
		return err
	}
	return nil
}

func UpdateOrderItem(orderItem *OrderItem) error {
	query := "UPDATE order_items SET quantity=$1, unit_price=$2, food_id=$3, order_id=$4, updated_at=$5 WHERE id=$6"
	_, err := database.DB.Exec(query, orderItem.Quantity, orderItem.UnitPrice, orderItem.FoodID, orderItem.OrderID, orderItem.UpdatedAt, orderItem.ID)
	if err != nil {
		return err
	}
	return nil
}

func GetOrderItemsByOrder(orderId int) ([]OrderItem, error) {
	query := "SELECT id, quantity, unit_price, food_id, order_id, created_at, updated_at FROM order_items WHERE order_id=$1"
	rows, err := database.DB.Query(query, orderId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orderItems []OrderItem
	for rows.Next() {
		var orderItem OrderItem
		err := rows.Scan(&orderItem.ID, &orderItem.Quantity, &orderItem.UnitPrice, &orderItem.FoodID, &orderItem.OrderID, &orderItem.CreatedAt, &orderItem.UpdatedAt)
		if err != nil {
			return nil, err
		}
		orderItems = append(orderItems, orderItem)
	}
	return orderItems, nil

}
