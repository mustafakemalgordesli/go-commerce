package service

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/mustafakemalgordesli/go-commerce/config"
	redispkg "github.com/mustafakemalgordesli/go-commerce/pkg/redis"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

var client *redis.Client

func TestMain(m *testing.M) {
	if err := config.SetupPath("./../.."); err != nil {
		log.Fatalf("config.Setup() error: %s", err)
		os.Exit(1)
	}
	if err := redispkg.Setup(); err != nil {
		log.Fatalf("redis.Setup() error: %s", err)
		fmt.Println("redis" + err.Error())
		os.Exit(1)
	}
	client = redispkg.Client
	fmt.Println("Before all tests")
	exitCode := m.Run()
	fmt.Println("After all tests")
	os.Exit(exitCode)
}

func setup() {
	ctx := context.Background()

	err := client.Set(ctx, "foo", "bar", 0).Err()
	if err != nil {
		panic(err)
	}
}
func clear() {
	ctx := context.Background()

	err := client.Set(ctx, "foo", "", 0).Err()
	if err != nil {
		panic(err)
	}
}

func TestRedis(t *testing.T) {
	setup()
	t.Run("Redis", func(t *testing.T) {
		ctx := context.Background()
		val, err := client.Get(ctx, "foo").Result()
		if err != nil {
			panic(err)
		}
		assert.Equal(t, "bar", val)
	})
	clear()
}
