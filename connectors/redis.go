package connectors

import (
	"errors"
	"fizz-buzz-api/utils"
	"fmt"
	"strings"
	"sync"

	"github.com/go-redis/redis/v7"
)

type RedisClient struct {
	*redis.Client
	mux sync.Mutex
}

// NewClient create a new redis client from app config
func NewClient(config utils.AppConfig) (*RedisClient, error) {
	addr := config.RedisHost
	if len(addr) == 0 {
		return nil, errors.New("redis url required")
	}
	if !strings.Contains(addr, ":") {
		addr = addr + ":6379"
	}
	password := config.RedisPassword
	if len(password) == 0 {
		fmt.Println("no password set for redis")
	}

	return &RedisClient{
		Client: redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: password,
			DB:       0,
		}),
	}, nil
}

func (c *RedisClient) Write(key, field string, value interface{}) error {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.Client.HSet(key, field, value).Err()
}

func (c *RedisClient) Read(key, field string) (string, error) {
	result, err := c.Client.HGet(key, field).Result()
	if err != nil {
		return "", fmt.Errorf("<error>::ReadFromRedis : %w", err)
	}
	return result, nil
}

func (c *RedisClient) GetFields(key string) ([]string, error) {
	result, err := c.Client.HKeys(key).Result()
	if err != nil {
		return nil, fmt.Errorf("<error>::ReadFromRedis : %w", err)
	}
	return result, nil
}

func (c *RedisClient) Ping() *redis.StatusCmd {
	return c.Client.Ping()
}
