package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
type User struct {
	Name      string    `json:"name" bson:"name"`
	Email     string    `json:"email" bson:"email"`
	Password  string    `json:"password" bson:"password"`
	Watchlist []Product `json:"watchlist" bson:"watchlist"`
}
*/

type User struct {
	ID            primitive.ObjectID `bson:"_id"`
	Name          string             `json:"name" bson:"name" validate:"required"`
	Password      string             `json:"password" bson:"password" validate:"required"`
	Email         string             `json:"email" bson:"email" validate:"email,required"`
	Watchlist     []WatchlistItem    `json:"watchlist" bson:"watchlist"`
	Token         string             `json:"token" bson:"token"`
	User_type     string             `json:"user_type" bson:"user_type" validate:"required,eq=ADMIN|eq=USER"`
	Refresh_token string             `json:"refresh_token" bson:"refresh_token"`
	Created_at    time.Time          `json:"created_at" bson:"created_at"`
	Updated_at    time.Time          `json:"updated_at" bson:"updated_at"`
	User_id       string             `json:"user_id" bson:"user_id"`
}

type WatchlistItem struct {
	ID      primitive.ObjectID `bson:"_id"`
	Image   string             `json:"image" bson:"image"`
	Name    string             `json:"name" bson:"name"`
	Price   string             `json:"price" bson:"price"`
	Rating  string             `json:"rating" bson:"rating"`
	Link    string             `json:"link" bson:"link"`
	Item_id string             `json:"item_id" bson:"item_id"`
}

type UpdateUser struct {
	NewName         string `json:"newName"`
	NewPassword     string `json:"newPassword"`
	CurrentPassword string `json:"currentPassword"`
}
