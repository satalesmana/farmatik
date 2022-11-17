package routes

import (
	ProductController "farmatik/app/controller/product"
	"farmatik/app/midleware"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(route *gin.Engine) *gin.Engine {
	productEndpoints := ProductController.NewProductControllerHandler()
	routeMidleware := midleware.NewAuthHandler().UseMiddleware()

	routers := route.Group("/product").Use(routeMidleware)
	{
		routers.GET("", productEndpoints.GetAll)
		routers.POST("", productEndpoints.AddNew)
		routers.GET("/:id", productEndpoints.FindById)
		routers.PUT("/:id", productEndpoints.Edit)
		routers.DELETE("/:id", productEndpoints.Delete)
	}

	return route
}
