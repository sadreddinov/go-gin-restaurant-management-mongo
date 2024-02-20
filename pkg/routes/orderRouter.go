package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sadreddinov/go-gin-restaurant-management-mongo/pkg/handler"
)

func OrderRoutes(incomingRoutes *gin.Engine)  {
	incomingRoutes.GET("/orders", handler.GetOrders())
	incomingRoutes.GET("/orders/:order_id", handler.GetOrder())
	incomingRoutes.POST("/orders", handler.CreateOrder())
	incomingRoutes.PATCH("/orders/:order_id", handler.UpdateOrder())
}