package routes

import (
	"github.com/gin-gonic/gin"
)

func CORSMiddleware(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}

	c.Next()
}

func Routes(route *gin.Engine) *gin.Engine {

	route.Use(CORSMiddleware)

	UserRoutes(route)
	ProductRoutes(route)
	AuthRoutes(route)
	PenjualanRoutes(route)
	return route
}
