package redis

import (
	"fmt"
	"talkspace/app/configs"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

func ConnectRedis() *redis.Client {
	config, err := configs.LoadConfig()
	if err != nil {
		logrus.Fatalf("failed to load redis configuration: %v", err)
		return nil
	}

	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s",
		 	  		config.REDIS.REDIS_HOST,
		      		config.REDIS.REDIS_PORT),
		DB:   1,
	})

	_, err = client.Ping().Result()
	if err != nil {
		logrus.Errorf("failed to connect to redis: %v", err)
		return nil
	}

	logrus.Info("connected to redis")
	return client
}