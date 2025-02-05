package models

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
	"github.com/TejaswiniYammanuru/golang-restaurant-management/database"
)

type Food struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" validate:"required,min=2,max=100"`
	Price     float64   `json:"price" validate:"required"`
	FoodImage string    `json:"food_image" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	MenuID    string    `json:"menu_id" validate:"required"`
}

func GetFoodByID(FoodID int) (*Food, error) {
	query := "SELECT id,name,price,food_image,created_at,updated_at,food_id,menu_id FROM food WHERE id=$1"

	var food Food

	err := database.DB.QueryRow(query, FoodID).Scan(
		&food.ID,
		&food.Name,
		&food.Price,
		&food.FoodImage,
		&food.CreatedAt,
		&food.UpdatedAt,
		&food.MenuID,
	)
	if err != nil {
		// If the error returned by the database query is sql.ErrNoRows,
		// it means that no rows were returned by the query. In this case,
		// we return a custom error message "Food not found" to indicate that
		// the requested food item was not found in the database.
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("Food not found")
		}

		return nil, err
	}
	return &food, nil
}

// The CreateFood function is responsible for inserting a new food item into the database.
// It takes a pointer to a Food struct as input, which contains the details of the food item to be created.
// The function returns an error if any issues occur during the database operation.
func CreateFood(food *Food) error {
	// The SQL query to insert a new food item into the food table.
	// The placeholders ($1, $2, etc.) are used to prevent SQL injection attacks.
	query := `INSERT INTO food (name, price, food_image, created_at, updated_at, food_id, menu_id) 
    VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`

	// Execute the SQL query with the provided food details.
	// The Scan function is used to retrieve the generated food ID from the database.
	err := database.DB.QueryRow(query, food.Name, food.Price, food.FoodImage, food.CreatedAt, food.UpdatedAt, food.MenuID).Scan(&food.ID)

	// If any error occurs during the database operation, return an error message.
	if err != nil {
		return fmt.Errorf("unable to insert food record: %v", err)
	}

	// If the food item is successfully inserted into the database, return nil to indicate success.
	return nil
}
