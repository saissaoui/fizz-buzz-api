package services

import (
	"fizz-buzz-api/services/fizzbuzz"
	"fizz-buzz-api/services/statistics"
	"fizz-buzz-api/utils"

	"go.uber.org/zap"
)

// Services struct wrapping all the app services
type Services struct {
	FizzBuzzService   fizzbuzz.Service
	StatisticsService statistics.Service
}

// Initializes the app services
func InitServices(config utils.AppConfig) *Services {

	fizzBuzzService := fizzbuzz.InitService()

	statisticsService, err := statistics.InitStatisticsService(config)

	if err != nil {
		utils.Logger.Error("Error initializing stats service", zap.Error(err))
	}

	return &Services{
		FizzBuzzService:   fizzBuzzService,
		StatisticsService: statisticsService,
	}

}
