package routes

import (
	"github.com/andremmoreno/go-api/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("/")
	{
		users := main.Group("users")
		{
			users.GET("/", controllers.GetUsers)
		}
	}

	return router
}
