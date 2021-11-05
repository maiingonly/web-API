package entities

import "time"

type User_account struct {
	User_ID     uint
	Name        string
	Email       string
	PhoneNumber string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Person struct {
}
