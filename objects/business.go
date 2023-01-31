package objects

import "time"

type Business struct {
	Id             string
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
	StateCode      string
	CrossStreets   string
	Phone          string
	IsClaimed      string
	DateOpened     time.Time
	DateClosed     *time.Time
	MessageUrl     string
	MessageUseCase string
	IsDeleted      bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
}

type BusinessHour struct {
	Id          int
	Day         int
	StartHour   string
	EndHour     string
	IsOvernight bool
	CreatedAt   time.Time
	UpdateAt    time.Time
}


