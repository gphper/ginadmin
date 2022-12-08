/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-08-20 20:46:25
 */
package redis

import (
	"github.com/gphper/ginadmin/configs"

	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

func Init() error {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     configs.App.Redis.Addr,
		Password: configs.App.Redis.Password,
		DB:       configs.App.Redis.Db,
	})

	err := RedisClient.Ping().Err()
	if err != nil {
		return err
	}
	return nil
}
