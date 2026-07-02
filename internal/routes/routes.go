package routes

import (
	"github.com/gin-gonic/gin"

	"test_project/internal/controllers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/test", controllers.TestRoute)

	return router
}
