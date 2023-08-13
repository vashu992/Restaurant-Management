package routes

import (
	controller "github.com/vashu992/Restaurant-Management/controllers"

	"github.com/gin-gonic/gin"
)
 
func TableRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/tables", controller.GetTables())
	incomingRoutes.GET("/tables/:order_id", controller.GetTable())
	incomingRoutes.POST("/ashutosh/", controller.CreateTable())
	incomingRoutes.PATCH("/tables/:order_id", controller.UpdateTable())
}