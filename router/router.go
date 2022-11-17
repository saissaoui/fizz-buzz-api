package router

import (
	"fizz-buzz-api/handlers"
	"fizz-buzz-api/handlers/fizzbuzz"
	"fizz-buzz-api/handlers/health"
	"fizz-buzz-api/handlers/stats"
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
	r.GET("/", health.GetHealth)
	r.GET("/health", health.GetHealth)

	//fizz buzz
	r.POST("/fizz-buzz", fizzbuzz.GetFizzBuzz(*services))

	//stats
	r.GET("/statistics", stats.GetStatistics(*services))

	// fallback
	r.NoRoute(handlers.NoRoute)
}
