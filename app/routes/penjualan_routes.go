package routes

import (
	PenjualanController "farmatik/app/controller/penjualan"
	"farmatik/app/midleware"

	"github.com/gin-gonic/gin"
)

func PenjualanRoutes(route *gin.Engine) *gin.Engine {
	penjualanEndpoints := PenjualanController.NewPenjualanControllerHandler()
	routeMidleware := midleware.NewAuthHandler().UseMiddleware()

	routers := route.Group("/penjualan").Use(routeMidleware)
	{
		routers.POST("/", penjualanEndpoints.AddNew)
		routers.GET("/:id", penjualanEndpoints.FindById)
		routers.GET("/", penjualanEndpoints.GetAll)
	}

	return route
}
