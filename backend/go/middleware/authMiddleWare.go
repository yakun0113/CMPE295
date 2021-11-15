package middleware

import (
	"net/http"
	"octopus/helpers"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		//clientToken := c.Request.Header.Get("token")
		clientToken, noTokenErr := c.Cookie("token")
		if noTokenErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Login to visit this page!"})
			c.Abort()
			return
		}
		if clientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Login to visit this page!"})
			c.Abort()
			return
		}

		claims, err := helpers.ValidateToken(clientToken)
		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}
		c.Set("email", claims.Email)
		c.Set("name", claims.Name)
		c.Set("uid", claims.Uid)
		c.Set("user_type", claims.User_type)
		c.Next()
	}
}
