package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sadreddinov/go-gin-restaurant-management-mongo/pkg/handler"
)

func UserRoutes(incomingRoutes *gin.Engine)  {
	incomingRoutes.GET("/users", handler.GetUsers())
	incomingRoutes.GET("/users/:user_id", handler.GetUser())
	incomingRoutes.POST("/users/signup", handler.SignUp())
	incomingRoutes.POST("/users/login", handler.Login())
}
