package entities

import (
	"time"
)

type Post struct {
	Id        uint
	Title     string
	Desc      string
	Tags      string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
