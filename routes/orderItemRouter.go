package routes

import (
	controller "github.com/vashu992/Restaurant-Management/controllers"

	"github.com/gin-gonic/gin"
)
 
func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orderItems", controller.GetOrderItems())
	incomingRoutes.GET("/orderItems/:orderItem_id", controller.GetOrderItemsByOrder())
	incomingRoutes.GET("/orderItems-order/:order_id", controller.GetOrderItemsByOrder())
	incomingRoutes.POST("/tables/", controller.CreateOrderItem())
	incomingRoutes.PATCH("/tables/:orderItem_id", controller.UpdateOrderItem())
}