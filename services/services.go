package services

import (
	"fizz-buzz-api/utils"

	"go.uber.org/zap"
)

type Services struct {
	FizzBuzzService   *FizzBuzzService
	StatisticsService *StatisticsService
}

func InitServices(config utils.AppConfig) *Services {

	fizzBuzzService := &FizzBuzzService{}

	statisticsService, err := InitStatisticsService(config)

	if err != nil {
		utils.Logger.Error("Error initializing stats service", zap.Error(err))
	}

	return &Services{
		FizzBuzzService:   fizzBuzzService,
		StatisticsService: statisticsService,
	}

}
