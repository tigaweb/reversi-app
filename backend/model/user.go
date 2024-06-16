package model

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email" gorm:"unique"`
	UserName  string    `json:"user_name" gorm:"unique"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserResponse struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	UserName string `json:"user_name" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
}
