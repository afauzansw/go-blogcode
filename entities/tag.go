package entities

import "time"

type Tag struct {
	Id        uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
