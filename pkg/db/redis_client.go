package db

import (
	"context"
	"product-mall/conf"

	// Go 语言中 Redis 客户端的版本 8 的模块
	"github.com/go-redis/redis/v8"
	"github.com/go-redis/redismock/v8"
)

var (
	client *redis.Client
	Mocker redismock.ClientMock
)

// 是用于创建一个 Redis 客户端的模拟/mock 实例，它可以用于在测试过程中模拟和模仿 Redis 客户端的行为，而不实际连接到真实的 Redis 服务器。这对于编写单元测试和集成测试时，可以避免对真实的 Redis 服务器造成影响，同时能够控制测试中所需要的各种情况和返回值
func InitMockClient() {
	cli, mock := redismock.NewClientMock()
	client = cli
	Mocker = mock
}

func InitRedis(ctx context.Context) (err error) {
	client = redis.NewClient(&redis.Options{
		Addr:     conf.RedisAddr,
		Password: conf.RedisPw, // no password set
		DB:       0,            // use default DB
	})
	_, err = client.Ping(ctx).Result()
	if err != nil {
		return
	}
	return err
}

func GetRedisClient() *redis.Client {
	return client
}
