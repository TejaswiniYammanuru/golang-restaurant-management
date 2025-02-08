package models

import (
	"database/sql"
	"time"

	"github.com/TejaswiniYammanuru/golang-restaurant-management/database"
)

type Table struct {
    ID             int       `json:"id"`
    NumberOfGuests int       `json:"number" validate:"required"`
    TableNumber    int       `json:"table_number" validate:"required"`
    CreatedAt      time.Time `json:"created_at"`
    UpdatedAt      time.Time `json:"updated_at"`
}


func GetTables() ([]Table, error) {
	query := "SELECT id,number_of_guests,table_number,created_at,updated_at FROM tables"

	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var tables []Table
	for rows.Next() {
		var table Table
		err := rows.Scan(&table.ID, &table.NumberOfGuests, &table.TableNumber, &table.CreatedAt, &table.UpdatedAt)
		if err != nil {
			return nil, err
		}
		tables = append(tables, table)
	}

	return tables, nil
}

func GetTableByID(id int) (*Table, error) {
	query := "SELECT id,number_of_guests,table_number,created_at,updated_at FROM tables WHERE id=$1"

	row := database.DB.QueryRow(query, id)

	var table Table
	err := row.Scan(&table.ID, &table.NumberOfGuests, &table.TableNumber, &table.CreatedAt, &table.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &table, nil
}

func CreateTable(table *Table) error {
	query := "INSERT INTO tables (number_of_guests, table_number, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id"

	err := database.DB.QueryRow(query, table.NumberOfGuests, table.TableNumber, table.CreatedAt, table.UpdatedAt).Scan(&table.ID)
	if err != nil {
		return err
	}

	return nil
}

func UpdateTable(table *Table) error {
	query := "UPDATE tables SET number_of_guests=$1, table_number=$2, updated_at=NOW() WHERE id=$3"

	_, err := database.DB.Exec(query, table.NumberOfGuests, table.TableNumber, table.ID)
	if err != nil {
		return err
	}

	return nil
}
