package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sadreddinov/go-gin-restaurant-management-mongo/pkg/handler"
)

func MenuRoutes(incomingRoutes *gin.Engine)  {
	incomingRoutes.GET("/menues", handler.GetMenues())
	incomingRoutes.GET("/menues/:menu_id", handler.GetMenu())
	incomingRoutes.POST("/menues", handler.CreateMenu())
	incomingRoutes.PATCH("menues/:menu_id", handler.UpdateMenu())
}