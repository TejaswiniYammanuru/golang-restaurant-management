package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"` 
	Password  string    `json:"password"`
	Avatar    string    `json:"avatar"`
	Phone     string    `json:"phone"`
	Role      string    `json:"role"`  
	CreatedAt time.Time `json:"created_at"`
}

func CreateUser(db *sql.DB, user *User) error {
	query := `
		INSERT INTO users (first_name, last_name, email, password, avatar, phone, role, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id
	`
	return db.QueryRow(
		query,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
		user.Avatar,
		user.Phone,
		user.Role,    
		time.Now(),
	).Scan(&user.ID)
}

func GetUserByEmail(db *sql.DB, email string) (*User, error) {
	user := &User{}
	query := `
		SELECT id, first_name, last_name, email, password, avatar, phone, role, created_at
		FROM users
		WHERE email = $1
	`
	err := db.QueryRow(query, email).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.Avatar,
		&user.Phone,
		&user.Role,  
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}
