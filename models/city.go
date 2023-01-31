package models

import "database/sql"

type City struct {
	Id         int
	Name       string
	CountryId  int
	StateId    sql.NullInt64
	Latitude   float64
	Longtitude float64
}
