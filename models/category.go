package models

import "time"

type Category struct {
	Id        int
	Alias     string
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
