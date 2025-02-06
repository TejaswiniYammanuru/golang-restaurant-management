package database

import (
	
    "database/sql"
    _ "github.com/lib/pq"
    "fmt"
    "log"
    
)

const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    password = "abc123"
    dbname   = "restaurant_management"
)

var DB *sql.DB // Global DB variable

// InitDB initializes the database connection and assigns it to DB
func InitDB() {
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
        "password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)

    var err error
    DB, err = sql.Open("postgres", psqlInfo)
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    err = DB.Ping()
    if err != nil {
        log.Fatalf("Failed to ping database: %v", err)
    }

    log.Println("Successfully connected to the database")
}


