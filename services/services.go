package services

import (
	"fizz-buzz-api/utils"

	"go.uber.org/zap"
)

// Services struct wrapping all the app services
type Services struct {
	FizzBuzzService   FizzBuzzService
	StatisticsService StatisticsService
}

// Initializes the app services
func InitServices(config utils.AppConfig) *Services {

	fizzBuzzService := &FizzBuzzServiceImpl{}

	statisticsService, err := InitStatisticsService(config)

	if err != nil {
		utils.Logger.Error("Error initializing stats service", zap.Error(err))
	}

	return &Services{
		FizzBuzzService:   fizzBuzzService,
		StatisticsService: statisticsService,
	}

}
