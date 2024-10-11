package common

/**
 * @Author elastic·H
 * @Date 2024-10-09
 * @File: redisService.go
 * @Description:
 */

import (
	"fmt"
	"gin-init/config"
	"github.com/go-redis/redis/v8"
)

var Client *redis.Client

type RedisService struct {
	Client *redis.Client
}

func NewRedisService() *RedisService {
	initRedis()

	return &RedisService{
		Client: Client,
	}
}

func initRedis() {
	// 连接 Redis, 获得 DB 实例
	Client = redis.NewClient(&redis.Options{
		Addr:     config.Conf.Redis.GetAddr(),
		Password: config.Conf.Redis.Password,
		DB:       config.Conf.Redis.DB,
	})
	fmt.Println("[Success] Redis数据库连接成功！！！")

}
