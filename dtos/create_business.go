package dtos

type CreateBusiness struct {
	Name         string        `json:"name"`
	Latitude     float64       `json:"latitude"`
	Longitude    float64       `json:"longitude"`
	Transactions []string      `json:"transactions"`
	Address1     string        `json:"address_1"`
	Address2     string        `json:"address_2"`
	Address3     string        `json:"address_3"`
	CityId       int           `json:"city_id"`
	ZipCode      string        `json:"zip_code"`
	CountryCode  string        `json:"country_code"`
	StateCode    string        `json:"state_code"`
	CrossStreets string        `json:"cross_streets"`
	Phone        string        `json:"phone"`
	DateOpened   string        `json:"date_opened"`
	Categories   []Category    `json:"categories"`
	OpenHours    []OpenHour    `json:"hours"`
	SpecialHour  []SpecialHour `json:"special_hour"`
}
