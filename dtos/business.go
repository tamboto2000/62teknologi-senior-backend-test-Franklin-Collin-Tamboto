package dtos

type Business struct {
	ID           string        `json:"id"`
	Alias        string        `json:"alias"`
	Name         string        `json:"name"`
	ImageURL     string        `json:"image_url"`
	IsClaimed    bool          `json:"is_claimed"`
	IsClosed     bool          `json:"is_closed"`
	URL          string        `json:"url"`
	Phone        string        `json:"phone"`
	DisplayPhone string        `json:"display_phone"`
	ReviewCount  int64         `json:"review_count"`
	Categories   []Category    `json:"categories"`
	Rating       float64       `json:"rating"`
	Location     Location      `json:"location"`
	Coordinates  Coordinates   `json:"coordinates"`
	Photos       []string      `json:"photos"`
	Price        string        `json:"price"`
	Hours        []Hour        `json:"hours"`
	Transactions []string      `json:"transactions"`
	Messaging    Messaging     `json:"messaging"`
	SpecialHours []SpecialHour `json:"special_hours"`
}

type Hour struct {
	Open      []OpenHour `json:"open"`
	HoursType string     `json:"hours_type"`
	IsOpenNow bool       `json:"is_open_now"`
}

type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Location struct {
	Address1       string   `json:"address1"`
	Address2       string   `json:"address2"`
	Address3       string   `json:"address3"`
	City           string   `json:"city"`
	ZipCode        string   `json:"zip_code"`
	Country        string   `json:"country"`
	State          string   `json:"state"`
	DisplayAddress []string `json:"display_address"`
	CrossStreets   string   `json:"cross_streets"`
}

type Messaging struct {
	URL         string `json:"url"`
	UseCaseText string `json:"use_case_text"`
}
