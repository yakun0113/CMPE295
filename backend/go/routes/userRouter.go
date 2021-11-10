package routes

import (
	"octopus/middleware"

	"octopus/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.GET("/users/:user_id", controllers.GetUser())
	incomingRoutes.PUT("/users/:user_id", controllers.UpdateUser())
	incomingRoutes.DELETE("/users/:user_id", controllers.DeleteUser())

	incomingRoutes.GET("/users/watchlist/:user_id", controllers.GetWatchlist())
	incomingRoutes.POST("/users/watchlist/:user_id", controllers.AddItem())
	incomingRoutes.DELETE("/users/watchlist/:user_id/*item_id", controllers.DeleteItem())

}
