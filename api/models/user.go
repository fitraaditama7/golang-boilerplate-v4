package models

import "time"

type User struct {
	ID        int       `json:"id" bind:"required"`
	Nama      string    `json:"nama" bind:"required"`
	Username  string    `json:"username" bind:"required"`
	Email     string    `json:"email" bind:"required"`
	Password  string    `json:"password" bind:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
