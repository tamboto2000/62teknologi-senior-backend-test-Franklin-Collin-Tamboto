package models

import "time"

type BusinessHour struct {
	Id          int
	BusinessId  int
	Day         int
	StartHour   string
	EndHour     string
	IsOvernight bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
