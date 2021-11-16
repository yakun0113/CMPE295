package controllers

import (
	"context"
	"net/http"
	"octopus/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		var item models.WatchlistItem
		var dbUser models.User
		user_id := c.Param("user_id")
		token_user_id := c.MustGet("uid")
		if token_user_id != user_id {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Cannot access this resource"})

		}

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		if err := userCollection.FindOne(ctx, bson.M{"user_id": user_id}).Decode(&dbUser); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Please try again!"})
			return
		}
		if err := c.BindJSON(&item); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Please try again!"})
			return
		}
		if len(dbUser.Watchlist) != 0 {
			for i := 0; i < len(dbUser.Watchlist); i++ {
				if item.Link == dbUser.Watchlist[i].Link {
					c.JSON(http.StatusBadRequest, gin.H{"error": "Already added this item!"})
					return
				}
			}

		}

		item.ID = primitive.NewObjectID()
		item.Item_id = item.ID.Hex()

		result, err := userCollection.UpdateOne(
			ctx,
			bson.M{"user_id": user_id},
			bson.M{"$push": bson.M{"watchlist": item}})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Please try again!"})
		} else if result != nil {
			c.JSON(http.StatusOK, gin.H{"message": "Item added to watchlist!"})

		}
	}
}

func DeleteItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dbUser models.User

		user_id := c.Param("user_id")
		token_user_id := c.MustGet("uid")
		if token_user_id != user_id {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Cannot access this resource"})

		}
		item_id := c.Param("item_id")[1:]

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		result, err := userCollection.UpdateOne(
			ctx,
			bson.M{"user_id": user_id},
			bson.M{"$pull": bson.M{"watchlist": bson.M{"item_id": item_id}}})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Please try again!"})
		} else if result != nil {
			err := userCollection.FindOne(ctx, bson.M{"user_id": user_id}).Decode(&dbUser)
			if err != nil {
				return
			} else {
				result := dbUser.Watchlist
				c.JSON(http.StatusOK, gin.H{"message": "Item removed from watchlist!", "result": result})
			}

		}
	}
}

func GetWatchlist() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dbUser models.User
		user_id := c.Param("user_id")
		token_user_id := c.MustGet("uid")
		if token_user_id != user_id {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Cannot access this resource"})

		}

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		err := userCollection.FindOne(ctx, bson.M{"user_id": user_id}).Decode(&dbUser)

		if err != nil {
			return
		} else {
			c.JSON(http.StatusCreated, dbUser.Watchlist)
		}
	}
}
