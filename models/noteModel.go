package models

import "time"

type Note struct {
	ID        int64     `json:"id" gorm:"primary_key"`
	Text      string    `json:"text"`
	Title     string    `json:"title" `
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
         
                