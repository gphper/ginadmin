/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-08-20 20:46:25
 */
package redisx

import (
	"github.com/gphper/ginadmin/configs"

	"github.com/go-redis/redis"
)

var redisClient *redis.Client

func Init() error {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     configs.App.Redis.Addr,
		Password: configs.App.Redis.Password,
		DB:       configs.App.Redis.Db,
	})

	err := redisClient.Ping().Err()
	if err != nil {
		return err
	}
	return nil
}

// 获取redis客户端
func GetRedisClient() *redis.Client {
	return redisClient
}
