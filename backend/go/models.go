package main

type User struct {
	Name      string    `json:"name" bson:"name"`
	Email     string    `json:"email" bson:"email"`
	Password  string    `json:"password" bson:"password"`
	Watchlist []Product `json:"watchlist" bson:"watchlist"`
}

type SearchPost struct {
	ItemName  string `json:"itemName"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type SearchResults struct {
	Store    string    `json:"store"`
	Location Location  `json:"location"`
	Products []Product `json:"products"`
}

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
