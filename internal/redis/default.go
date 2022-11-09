/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-08-20 20:46:25
 */
package redis

import (
	"fmt"

	"github.com/gphper/ginadmin/configs"

	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

func Init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     configs.App.Redis.Addr,
		Password: configs.App.Redis.Password,
		DB:       configs.App.Redis.Db,
	})

	fmt.Println(configs.App.Redis.Password)

	err := RedisClient.Ping().Err()
	if err != nil {
		panic("redis connect error")
	}
}
