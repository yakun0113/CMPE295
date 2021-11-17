package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"octopus/helpers"

	"octopus/database"

	"octopus/models"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

//var validate = validator.New()

func Signup() gin.HandlerFunc {

	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var user models.User
		if err := c.BindJSON(&user); err != nil {
			log.Panic(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		/*
			validationErr := validate.Struct(user)
			if validationErr != nil {
				println("errs")

				c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
				return
			}
		*/
		count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occured while checking for the email"})
		}

		password := helpers.HashPassword(user.Password)
		user.Password = password

		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "This email already exists!"})
			return
		}
		user.User_type = "USER"
		user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.User_id = user.ID.Hex()
		token, refreshToken, _ := helpers.GenerateAllTokens(user.Email, user.Name, user.User_type, user.User_id)
		user.Token = token
		user.Refresh_token = refreshToken

		resultInsertionNumber, insertErr := userCollection.InsertOne(ctx, user)
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to sign up, please try again!"})
			return
		}
		defer cancel()
		fmt.Println(resultInsertionNumber)
		c.JSON(http.StatusOK, gin.H{"message": "Sign up successfully!"})
	}

}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var user models.User
		var foundUser models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "This email is not registered!"})
			return
		}

		passwordIsValid := helpers.VerifyPassword(user.Password, foundUser.Password)
		defer cancel()
		if !passwordIsValid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong password, please try again!"})
			return
		}
		/*
			if foundUser.Email == "" {
				c.JSON(http.StatusOK, gin.H{"message": "This email is not registered!"})
			}
		*/
		token, refreshToken, _ := helpers.GenerateAllTokens(foundUser.Email, foundUser.Name, foundUser.User_type, foundUser.User_id)
		helpers.UpdateAllTokens(token, refreshToken, foundUser.User_id)
		err = userCollection.FindOne(ctx, bson.M{"user_id": foundUser.User_id}).Decode(&foundUser)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.SetCookie("token", token, 900, "/", "localhost", false, true)
		c.SetCookie("user_id", foundUser.User_id, 900, "/", "localhost", false, false)
		c.JSON(http.StatusOK, foundUser)
	}
}

func Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.SetCookie("token", "0", 1, "/", "localhost", false, true)
		c.SetCookie("user_id", "0", 1, "/", "localhost", false, false)

		c.JSON(http.StatusOK, gin.H{"message": "Signed out successfully!"})
	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		user_id := c.Param("user_id")
		token_user_id := c.MustGet("uid")
		if token_user_id != user_id {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Cannot access this resource"})

		}
		/*
			if err := helpers.MatchUserTypeToUid(c, user_id); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}*/
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var user models.User
		err := userCollection.FindOne(ctx, bson.M{"user_id": user_id}).Decode(&user)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, user)
	}
}

func DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		user_id := c.Param("user_id")
		token_user_id := c.MustGet("uid")
		if token_user_id != user_id {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Cannot access this resource"})

		}

		result, err := userCollection.DeleteOne(ctx, bson.M{"user_id": user_id})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		} else if result != nil {
			c.JSON(http.StatusOK, gin.H{"message": "Account deleted, hope to see you again!"})
		}

	}
}

func UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		user_id := c.Param("user_id")
		token_user_id := c.MustGet("uid")
		if token_user_id != user_id {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Cannot access this resource"})

		}

		var updateUser models.UpdateUser
		var foundUser models.User
		if err := c.BindJSON(&updateUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := userCollection.FindOne(ctx, bson.M{"user_id": user_id}).Decode(&foundUser)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Please try again!"})
			return
		}

		passwordIsValid := helpers.VerifyPassword(updateUser.CurrentPassword, foundUser.Password)
		defer cancel()
		if !passwordIsValid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong password, please try again!"})
			return
		}

		if updateUser.NewPassword != "" {
			newPassword := helpers.HashPassword(updateUser.NewPassword)
			result, err := userCollection.UpdateOne(
				ctx,
				bson.M{"user_id": user_id},
				bson.M{"$set": bson.M{"password": newPassword, "name": updateUser.NewName}})
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Please try again!"})
			} else if result != nil {
				c.JSON(http.StatusOK, gin.H{"message": "Updated successfully!"})

			}
		} else {
			result, err := userCollection.UpdateOne(
				ctx,
				bson.M{"user_id": user_id},
				bson.M{"$set": bson.M{"name": updateUser.NewName}})
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Please try again!"})
			} else if result != nil {
				c.JSON(http.StatusOK, gin.H{"message": "Updated successfully!"})

			}

		}
	}
}

/*
func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helpers.CheckUserType(c, "ADMIN"); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		recordPerPage, err := strconv.Atoi(c.Query("recordPerPage"))
		if err != nil || recordPerPage < 1 {
			recordPerPage = 10
		}
		page, err1 := strconv.Atoi(c.Query("page"))
		if err1 != nil || page < 1 {
			page = 1
		}

		startIndex := (page - 1) * recordPerPage
		startIndex, err = strconv.Atoi(c.Query("startIndex"))

		matchStage := bson.D{{"$match", bson.D{{}}}}
		groupStage := bson.D{{"$group", bson.D{
			{"_id", bson.D{{"_id", "null"}}},
			{"total_count", bson.D{{"$sum", 1}}},
			{"data", bson.D{{"$push", "$$ROOT"}}}}}}
		projectStage := bson.D{
			{"$project", bson.D{
				{"_id", 0},
				{"total_count", 1},
				{"user_items", bson.D{{"$slice", []interface{}{"$data", startIndex, recordPerPage}}}}}}}
		result, err := userCollection.Aggregate(ctx, mongo.Pipeline{
			matchStage, groupStage, projectStage})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing user items"})
		}
		var allusers []bson.M
		if err = result.All(ctx, &allusers); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, allusers[0])
	}
}
*/
