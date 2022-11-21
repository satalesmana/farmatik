package routes

import (
	appconfigController "farmatik/app/controller/appconfig"
	"farmatik/app/midleware"

	"github.com/gin-gonic/gin"
)

func AppConfigRoutes(route *gin.Engine) *gin.Engine {
	appConfigEndpoints := appconfigController.NewAppConfigControllerHandler()
	routeMidleware := midleware.NewAuthHandler().UseMiddleware()

	routers := route.Group("/appconfig").Use(routeMidleware)
	{
		routers.POST("/", appConfigEndpoints.FindById)
	}

	return route
}
