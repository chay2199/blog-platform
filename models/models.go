package models

import (
	"time"
)

type Post struct {
	ID        uint      `gorm:"primarykey"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
}

type User struct {
	ID       uint   `gorm:"primarykey"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"` // i.e. "reader", "writer", "administrator"
}
