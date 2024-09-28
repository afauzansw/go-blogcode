package entities

import "time"

type Publisher struct {
	Id        uint
	Name      string
	Email     string
	JobTitle  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
