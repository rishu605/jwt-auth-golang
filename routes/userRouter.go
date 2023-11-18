package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/rishabh/jwt-auth/controllers"
	"github.com/rishabh/jwt-auth/middlewares"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middlewares.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
}
