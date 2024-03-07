package redis

import (
	"github.com/mustafakemalgordesli/go-commerce/config"
	"github.com/redis/go-redis/v9"
	"github.com/streadway/amqp"
)

var (
	Client   *redis.Client
	RedisErr error
)

type Redis struct {
	*amqp.Connection
}

func Setup() error {
	configs := config.GetConfig()

	opt, err := redis.ParseURL(configs.Redis.Connection)

	if err != nil {
		RedisErr = err
		return err
	}

	Client = redis.NewClient(opt)
	return nil
}

func GetConn() *redis.Client {
	return Client
}

func GetConnErr() error {
	return RedisErr
}
