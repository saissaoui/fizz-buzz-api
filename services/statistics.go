package services

import (
	"fizz-buzz-api/connectors"
	"fizz-buzz-api/models"
	"fizz-buzz-api/utils"
	"strconv"

	"github.com/pkg/errors"
)

//go:generate mockery --name=StatisticsService --output ./../mocks --case=underscore

type StatisticsService interface {
	GetStatistics() (*models.StatisticsResponse, error)
	CountRequest(request *models.FizzBuzzRequest) error
}
type StatisticsServiceImpl struct {
	RedisClient *connectors.RedisClient
}

// InitStatisticsService initialises a new StatisticsService
func InitStatisticsService(config utils.AppConfig) (*StatisticsServiceImpl, error) {
	redisClient, err := connectors.NewClient(config)

	if err != nil {
		return nil, errors.Wrap(err, "InitStatisticsService")
	}
	return &StatisticsServiceImpl{
		RedisClient: redisClient,
	}, nil
}

// Gets the requests statistics from redis
func (s StatisticsServiceImpl) GetStatistics() (resp *models.StatisticsResponse, err error) {
	// check redis availability
	if err = s.RedisClient.Ping().Err(); err != nil {
		return
	}
	fields, err := s.RedisClient.GetFields("statistics")

	if err != nil {
		return
	}
	resp = new(models.StatisticsResponse)
	stats := make([]*models.RequestStats, 0)
	for _, field := range fields {
		countString, err := s.RedisClient.Read("statistics", field)
		if err != nil {
			return nil, err
		}
		count, _ := strconv.Atoi(countString)
		stats = append(stats, &models.RequestStats{
			Request: field,
			Count:   count,
		})
	}
	resp.Stats = stats

	return
}

// CountRequest add a new request to count or increase counting for an already requested combination
func (s StatisticsServiceImpl) CountRequest(request *models.FizzBuzzRequest) error {

	// check redis availability
	if err := s.RedisClient.Ping().Err(); err != nil {
		return err
	}

	count := 1
	val, err := s.RedisClient.Read("statistics", request.ToString())

	if err == nil {
		count, err = strconv.Atoi(val)
		if err == nil {
			count++
		}
	}

	err = s.RedisClient.Write("statistics", request.ToString(), count)

	return err
}
