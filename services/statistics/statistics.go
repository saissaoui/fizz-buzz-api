//go:generate mockgen -destination=../../test/mocks/services/statistics/statistics.go -package=statistics -source=./statistics.go

package statistics

import (
	"fizz-buzz-api/connectors"
	"fizz-buzz-api/services/fizzbuzz"
	"fizz-buzz-api/utils"
	"strconv"

	"github.com/pkg/errors"
)

type Service interface {
	GetStatistics() ([]RequestStats, error)
	CountRequest(request fizzbuzz.Command) error
}
type StatisticsService struct {
	RedisClient *connectors.RedisClient
}

// InitStatisticsService initialises a new StatisticsService
func InitStatisticsService(config utils.AppConfig) (*StatisticsService, error) {
	redisClient, err := connectors.NewClient(config)

	if err != nil {
		return nil, errors.Wrap(err, "InitStatisticsService")
	}
	return &StatisticsService{
		RedisClient: redisClient,
	}, nil
}

// GetStatistics Gets the requests statistics from redis
func (s StatisticsService) GetStatistics() ([]RequestStats, error) {
	// check redis availability
	if err := s.RedisClient.Ping().Err(); err != nil {
		return nil, err
	}
	fields, err := s.RedisClient.GetFields("statistics")

	if err != nil {
		return nil, err

	}
	statistics := make([]RequestStats, 0)
	for _, field := range fields {
		countString, err := s.RedisClient.Read("statistics", field)
		if err != nil {
			return nil, err
		}
		count, _ := strconv.Atoi(countString)
		statistics = append(statistics, RequestStats{
			Request: field,
			Count:   count,
		})
	}

	return statistics, nil
}

// CountRequest add a new request to count or increase counting for an already requested combination
func (s StatisticsService) CountRequest(command fizzbuzz.Command) error {

	// check redis availability
	if err := s.RedisClient.Ping().Err(); err != nil {
		return err
	}

	count := 1
	val, err := s.RedisClient.Read("statistics", command.ToString())

	if err == nil {
		count, err = strconv.Atoi(val)
		if err == nil {
			count++
		}
	}

	err = s.RedisClient.Write("statistics", command.ToString(), count)

	return err
}
