package routes

import (
	authController "farmatik/app/controller/auth"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(route *gin.Engine) *gin.Engine {
	authEndpoints := authController.NewAuthHandler()

	routers := route.Group("/auth")
	{
		routers.POST("/login", authEndpoints.Login)   //get all data
		routers.POST("/logout", authEndpoints.Logout) // add new
	}

	return route
}
