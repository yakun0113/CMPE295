package models

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Product struct {
	ID     int32  `json:"id"`
	Image  string `json:"image"`
	Name   string `json:"name"`
	Price  string `json:"price"`
	Rating string `json:"rating"`
	Link   string `json:"link"`
}

type SearchPost struct {
	ItemName  string `json:"itemName"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type SearchResults struct {
	Store    string    `json:"store"`
	Website  string    `json:"website"`
	Location Location  `json:"location"`
	Products []Product `json:"products"`
}
