package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sadreddinov/go-gin-restaurant-management-mongo/pkg/handler"
)

func OrderItemRoutes(incomingRoutes *gin.Engine)  {
	incomingRoutes.GET("/orderItems", handler.GetOrderItems())
	incomingRoutes.GET("/orderItems/:orderItem_id", handler.GetOrderItem())
	incomingRoutes.GET("/orderItems-order/:order_id", handler.GetOrderItemsByOrder())
	incomingRoutes.POST("/orderItems", handler.CreateOrderItem())
	incomingRoutes.PATCH("/orderItems/:orderItem_id", handler.UpdateOrderItem())
}