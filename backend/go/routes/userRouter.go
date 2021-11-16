package routes

import (
	"octopus/middleware"

	"octopus/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	// The following two routes is only for admin
	incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.GET("/users/:user_id", controllers.GetUser())

	incomingRoutes.PUT("/users", controllers.UpdateUser())
	incomingRoutes.DELETE("/users", controllers.DeleteUser())

	incomingRoutes.POST("users/logout", controllers.Logout())
	incomingRoutes.POST("/search", controllers.Search())

	incomingRoutes.GET("/users/watchlist", controllers.GetWatchlist())
	incomingRoutes.POST("/users/watchlist", controllers.AddItem())
	incomingRoutes.DELETE("/users/watchlist/:item_id", controllers.DeleteItem())

}
