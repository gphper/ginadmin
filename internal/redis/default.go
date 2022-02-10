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

func init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     configs.App.RedisConf.Addr,
		Password: configs.App.RedisConf.Password,
		DB:       configs.App.RedisConf.Db,
	})

	err := RedisClient.Ping().Err()
	if err != nil {
		panic("redis connect error")
	}
}
