package entities

import (
	"time"
)

type Post struct {
	Id          uint
	Title       string
	Description string
	Tags        string
	Status      string
	Slug        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
