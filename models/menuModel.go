package models

import (
	"fmt"
	"time"

	"github.com/TejaswiniYammanuru/golang-restaurant-management/database"
)

type Menu struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Category  string    `json:"category"`
	StartDate time.Time `json:"start-date"`
	EndDate   time.Time `json:"end-date"`
	CreatedAt time.Time `json:"created-at"`
	UpdatedAt time.Time `json:"updated-at"`
	Foods     []Food    `json:"foods" gorm:"foreignKey:MenuID"`

	
}

func GetMenus() ([]Menu, error) {
	query := "SELECT id,name,category,start_date,end_date,created_at,updated_at FROM menu"
	rows, err := database.DB.Query(query)//Query executes a query that returns rows, typically a SELECT. The args are for any placeholder parameters in the query.
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var menus []Menu
	for rows.Next() {
		var menu Menu
		err := rows.Scan(&menu.ID, &menu.Name, &menu.Category, &menu.StartDate, &menu.EndDate, &menu.CreatedAt, &menu.UpdatedAt)
		if err != nil {
			return nil, err
		}		
		menu.Foods,_=GetFoodsByMenuID(menu.ID);



		menus = append(menus, menu)
	}
	return menus, nil
}

func GetMenuByID(menuId int) (*Menu, error) {
	query := "SELECT id,name,category,start_date,end_date,created_at,updated_at FROM menu WHERE id=$1"

	var menu Menu

	err := database.DB.QueryRow(query, menuId).Scan(
		&menu.ID,
		&menu.Name,
		&menu.Category,
		&menu.StartDate,
		&menu.EndDate,
		&menu.CreatedAt,
		&menu.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	menu.Foods,_=GetFoodsByMenuID(menu.ID);
	return &menu, nil
}

func CreateMenu(menu *Menu) error {

    fmt.Printf("Inserting values:\nName: %s\nCategory: %s\nStartDate: %v\nEndDate: %v\nCreatedAt: %v\nUpdatedAt: %v\n",
        menu.Name, menu.Category, menu.StartDate, menu.EndDate, menu.CreatedAt, menu.UpdatedAt)

	query := `INSERT INTO menu (name, category, start_date,end_date, created_at, updated_at) 
    VALUES ($1, $2, $3, $4, $5,$6) RETURNING id`

	err := database.DB.QueryRow(query, menu.Name, menu.Category, menu.StartDate, menu.EndDate, menu.CreatedAt, menu.UpdatedAt).Scan(&menu.ID)

	if err != nil {
		return fmt.Errorf("unable to insert menu record: %v", err)
	}

	return nil
}



func UpdateMenu(menu Menu)error{
    query := `UPDATE menu SET name=$1, category=$2, start_date=$3, end_date=$4, updated_at=$5 WHERE id=$6 RETURNING id`
    _, err := database.DB.Exec(query, menu.Name, menu.Category, menu.StartDate, menu.EndDate, menu.UpdatedAt, menu.ID)
    if err!= nil{
        return fmt.Errorf("unable to update menu record: %v",err)
    }
    return nil
}


func DeleteFoodItemsByMenuID(menuid int) error{
    query := `DELETE FROM food WHERE menu_id=$1`
    _, err := database.DB.Exec(query, menuid)
    if err!= nil{
        return fmt.Errorf("unable to delete food items by menu id: %v",err)
    }
    return nil
}

func DeleteMenu(menuid int) error{
    query := `DELETE FROM menu WHERE id=$1`
    _, err := database.DB.Exec(query, menuid)
    if err!= nil{
        return fmt.Errorf("unable to delete menu: %v",err)
    }
    return nil
}

