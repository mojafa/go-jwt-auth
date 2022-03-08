package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mojafa/go-jwt-auth/controllers"
	"github.com/mojafa/go-jwt-auth/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.GET("/users/:userId", controllers.GetUser())
}
