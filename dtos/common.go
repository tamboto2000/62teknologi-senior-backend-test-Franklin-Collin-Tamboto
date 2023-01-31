package dtos

type Category struct {
	Id    int    `json:"id,omitempty"`
	Alias string `json:"alias"`
	Title string `json:"title"`
}

type OpenHour struct {
	IsOvernight bool   `json:"is_overnight"`
	Start       string `json:"start"`
	End         string `json:"end"`
	Day         int64  `json:"day"`
}

type SpecialHour struct {
	Date        string  `json:"date"`
	IsClosed    *bool   `json:"is_closed"`
	Start       *string `json:"start"`
	End         *string `json:"end"`
	IsOvernight *bool   `json:"is_overnight"`
}
