package models

import (
	"fmt"
	"golang-restaurant-management/database"
	"time"
)

type Menu struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Category  string    `json:"category"`
	StartDate time.Time `json:"start-date"`
	EndDate   time.Time `json:"end-date"`
	CreatedAt time.Time `json:"created-at"`
	UpdatedAt time.Time `json:"updated-at"`
	
}


func GetMenus()([]Menu,error){
	query := "SELECT id,name,category,start_date,end_date,created_at,updated_at FROM menu"
    rows,err := database.DB.Query(query)
    if err!= nil{
        return nil,err
    }
    defer rows.Close()

    var menus []Menu
    for rows.Next(){
        var menu Menu
        err := rows.Scan(&menu.ID,&menu.Name,&menu.Category,&menu.StartDate,&menu.EndDate,&menu.CreatedAt,&menu.UpdatedAt)
        if err!= nil{
            return nil,err
        }
        menus = append(menus,menu)
    }
    return menus,nil
}



func GetMenuByID(menuId int)(*Menu,error){
	query := "SELECT id,name,category,start_date,end_date,created_at,updated_at FROM menu WHERE id=$1"

    var menu Menu

    err := database.DB.QueryRow(query,menuId).Scan(
        &menu.ID,
        &menu.Name,
        &menu.Category,
        &menu.StartDate,
        &menu.EndDate,
        &menu.CreatedAt,
        &menu.UpdatedAt,
       
    )
    if err!= nil{
        return nil,err
    }
    return &menu,nil
}

func CreateMenu(menu *Menu)(error){
	query := `INSERT INTO menu (name, category, start_date,end_date, created_at, updated_at) 
    VALUES ($1, $2, $3, $4, $5) RETURNING id`

    
    err := database.DB.QueryRow(query, menu.Name, menu.Category, menu.StartDate,menu.EndDate, menu.CreatedAt, menu.UpdatedAt).Scan(&menu.ID)

    
    if err != nil {
        return fmt.Errorf("unable to insert menu record: %v", err)
    }

	

    
    return nil
}
