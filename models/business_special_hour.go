package models

import (
	"database/sql"
	"time"
)

type BusinessSpecialHour struct {
	Id          int
	BusinessId  int
	Day         int
	IsClosed    bool
	Start       sql.NullString
	End         sql.NullString
	IsOvernight bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
