package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sadreddinov/go-gin-restaurant-management-mongo/pkg/handler"
)

func TableRoutes(incomingRoutes *gin.Engine)  {
	incomingRoutes.GET("/tables", handler.GetTables())
	incomingRoutes.GET("/tables/:table_id", handler.GetTable())
	incomingRoutes.POST("/tables", handler.CreateTable())
	incomingRoutes.PATCH("/tables/:table_id", handler.UpdateTable())
}