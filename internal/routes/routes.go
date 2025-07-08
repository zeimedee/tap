package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/zeimedee/loverboy/internal/handlers"
	"github.com/zeimedee/loverboy/internal/services"
)

func SetUpRoutes() *gin.Engine {
	router := gin.Default()

	service := services.NewLoversService()
	tapHandler := handlers.NewTapHandler(service)

	api := router.Group("/api/v1/loverboy")
	{
		api.GET("/healthcheck", handlers.Healthcheck)
		api.POST("/register", tapHandler.Register)
		api.GET("/tap/:id", tapHandler.Tap)
	}

	return router
}
