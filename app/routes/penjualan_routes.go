package routes

import (
	PenjualanController "farmatik/app/controller/penjualan"

	"github.com/gin-gonic/gin"
)

func PenjualanRoutes(route *gin.Engine) *gin.Engine {
	penjualanEndpoints := PenjualanController.NewPenjualanControllerHandler()

	routers := route.Group("/penjualan")
	{
		routers.POST("/", penjualanEndpoints.AddNew)
		routers.GET("/:id", penjualanEndpoints.FindById)
	}

	return route
}
