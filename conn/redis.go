package conn

import (
	"context"
	"github.com/go-redis/redis/v8"
)

func RedisConnection(addr, username, password string) (*redis.Client, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     addr,
		Username: username,
		Password: password,
	})
	err := redisClient.Ping(context.Background()).Err()
	if err != nil {
		return nil, err
	}
	return redisClient, nil
}
