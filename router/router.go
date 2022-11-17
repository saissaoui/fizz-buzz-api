package router

import (
	"fizz-buzz-api/handlers"
	"fizz-buzz-api/services"

	"github.com/gin-gonic/gin"
)

func InitializeRouter(services *services.Services) *gin.Engine {

	// Set the default gin router
	r := gin.New()

	r.Use(gin.Recovery())

	// Initialize middlewares
	initializeMiddlewares(r)

	// Initialize routes
	initializeRoutes(r, services)

	return r

}

func initializeRoutes(r *gin.Engine, services *services.Services) {

	// health
	r.GET("/", handlers.GetHealth)
	r.GET("/health", handlers.GetHealth)

	//fizz buzz
	r.POST("/fizz-buzz", handlers.GetFizzBuzz(*services))

	//stats
	r.GET("/statistics", handlers.GetStatistics(*services))

	// fallback
	r.NoRoute(handlers.NoRoute)
}
