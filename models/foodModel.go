package models

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
	"github.com/TejaswiniYammanuru/golang-restaurant-management/database"
)

type Food struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" validate:"required,min=2,max=100"`
	Price     float64   `json:"price" validate:"required"`
	FoodImage string    `json:"food_image" validate:"required"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	MenuID    int    `json:"menu_id" validate:"required" gorm:"not null;index"`

	
	
}


func GetFoodByID(FoodID int) (*Food, error) {
	query := "SELECT id,name,price,food_image,created_at,updated_at,menu_id FROM food WHERE id=$1"

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
	query := `INSERT INTO food (name, price, food_image, created_at, updated_at,menu_id) 
    VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

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


func GetAllFoods()([]Food,error){
	query := "SELECT id,name,price,food_image,created_at,updated_at,menu_id FROM food"

    rows,err := database.DB.Query(query)
    if err!= nil{
        return nil,err
    }

    defer rows.Close()

    var foods []Food

    for rows.Next(){
        var food Food
        err := rows.Scan(&food.ID, &food.Name, &food.Price, &food.FoodImage, &food.CreatedAt, &food.UpdatedAt, &food.MenuID)
        if err!= nil{
            return nil,err
        }

        foods = append(foods, food)
    }

    return foods,nil

}


func UpdateFood(food *Food) error {
	query := `UPDATE food SET name=$1, price=$2, food_image=$3, updated_at=$4, menu_id=$5 WHERE id=$6`

    result, err := database.DB.Exec(query, food.Name, food.Price, food.FoodImage, food.UpdatedAt, food.MenuID, food.ID)
    if err!= nil {
        return fmt.Errorf("unable to update food record: %v", err)
    }

    rowsAffected, err := result.RowsAffected()
    if err!= nil {
        return fmt.Errorf("failed to get rows affected: %v", err)
    }

    if rowsAffected == 0 {
        return errors.New("food not found")
    }

    return nil
	
}


func DeleteFood(foodID int) error {
	query := "DELETE FROM food WHERE id=$1"

    result, err := database.DB.Exec(query, foodID)
    if err!= nil {
        return fmt.Errorf("unable to delete food record: %v", err)
    }

    rowsAffected, err := result.RowsAffected()
    if err!= nil {
        return fmt.Errorf("failed to get rows affected: %v", err)
    }

    if rowsAffected == 0 {
        return errors.New("food not found")
    }

    return nil
}


func GetFoodsByMenuID(menuID int)([]Food, error){
	query := "SELECT id,name,price,food_image,created_at,updated_at,menu_id FROM food where menu_id=$1"

    rows,err := database.DB.Query(query,menuID)
    if err!= nil{
        return nil,err
    }

    defer rows.Close()

    var foods []Food

    for rows.Next(){
        var food Food
        err := rows.Scan(&food.ID, &food.Name, &food.Price, &food.FoodImage, &food.CreatedAt, &food.UpdatedAt, &food.MenuID)
        if err!= nil{
            return nil,err
        }

        foods = append(foods, food)
    }

    return foods,nil


}
