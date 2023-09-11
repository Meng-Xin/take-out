package initialize

import (
	"fmt"
	"github.com/go-redis/redis"
	"take-out/global"
)

func initRedis() *redis.Client {
	redisOpt := global.Config.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisOpt.Host, redisOpt.Port),
		Password: redisOpt.Password, // no password set
		DB:       redisOpt.DataBase, // use default DB
	})
	ping := client.Ping()
	err := ping.Err()
	if err != nil {
		panic(err)
	}
	return client
}
