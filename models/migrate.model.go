package models

type APIResponse struct {
	Status bool   `json:"status"`
	Code   int    `json:"code"`
	Data   []Beer `json:"data"`
}

type Beer struct {
	ID       uint
	Name     string // Name of the beer
	Category string // Category of the beer (e.g., Lager, Ale, etc.)
	Detail   string // Detailed description of the beer
	Picture  string // URL or path to the picture of the beer
}
