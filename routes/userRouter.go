package routes

import (
	controller "github.com/vashu992/Restaurant-Management/controllers"

	"github.com/gin-gonic/gin"
)
 
func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
	incomingRoutes.POST("/users/signup", controller.CreateSignUp())
	incomingRoutes.PATCH("/users/:login", controller.UpdateLogin())
}