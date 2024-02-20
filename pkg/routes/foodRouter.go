package routes

import (
	"github.com/sadreddinov/go-gin-restaurant-management-mongo/pkg/handler"

	"github.com/gin-gonic/gin"
)

func FoodRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/foods", handler.GetFoods())
	incomingRoutes.GET("/foods/:food_id", handler.GetFood())
	incomingRoutes.POST("/foods", handler.CreateFood())
	incomingRoutes.PATCH("foods/:food_id", handler.UpdateFood())	
}	