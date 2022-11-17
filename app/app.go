package app

import (
	"fmt"

	"fizz-buzz-api/router"
	"fizz-buzz-api/services"
	"fizz-buzz-api/utils"

	"github.com/gin-gonic/gin"
)

type App struct {
	Config utils.AppConfig
	Router *gin.Engine
}

func New() *App {
	app := &App{}
	app.setup()
	return app
}

// Setup the server by loading config an initializing services
func (app *App) setup() {
	// Load configuration
	config := utils.LoadConfig()
	// Initialize Services
	servicesWrapper := services.InitServices(config)
	// Initialize Router
	r := router.InitializeRouter(servicesWrapper)
	app.Config = config
	app.Router = r

}

// Run the application server
func (app *App) Run() {
	// Serving application
	port := app.Config.Port
	err := app.Router.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		return
	}
}
