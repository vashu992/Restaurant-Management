package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	helper "github.com/vashu992/Restaurant-Management/helpers"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) { 
		clientToken := c.Request.Header.Get("token")
		if clientToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No authentication header provided"}) 
			c.Abort()
			return
		}

		claims, err := helper.ValidateToken(clientToken)
		if err != nil { 
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) 
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Set("first_name", claims.First_name)
		c.Set("last_name", claims.Last_name)
		c.Set("uid", claims.Uid)

		c.Next()
	}
}
