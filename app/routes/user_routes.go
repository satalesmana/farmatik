package routes

import (
	usersController "farmatik/app/controller/user"
	"farmatik/app/midleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(route *gin.Engine) *gin.Engine {

	userEndpoints := usersController.NewUserHandler()
	routeMidleware := midleware.NewAuthHandler().UseMiddleware()

	routers := route.Group("/users").Use(routeMidleware)
	{
		routers.GET("/", userEndpoints.GetAll)       //get all data
		routers.POST("/", userEndpoints.AddNew)      // add new
		routers.GET("/:id", userEndpoints.FindBy)    // find by id
		routers.PUT("/:id", userEndpoints.Edit)      // find by id
		routers.DELETE("/:id", userEndpoints.Delete) // find by id
	}

	return route
}
