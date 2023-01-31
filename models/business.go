package models

import (
	"database/sql"
	"time"
)

type Business struct {
	ID             string
	Alias          string
	Name           string
	ImageUrl       string
	IsClosed       bool
	Latitude       float64
	Longtitude     float64
	Transactions   []string
	Address1       string
	Address2       string
	Address3       string
	CityId         int
	ZipCode        string
	CountryCode    string
	StateCode      sql.NullString
	CrossStreets   sql.NullString
	Phone          string
	IsClaimed      bool
	DateOpened     time.Time
	DateClosed     sql.NullTime
	MessageUrl     sql.NullString
	MessageUseCase sql.NullString
	IsDeleted      bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      sql.NullTime
	City           *City
	Country        *Country
	State          *State
}
