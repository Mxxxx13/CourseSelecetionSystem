// @Title : redis
// @Description :连接redis
// @Author : MX
// @Update : 2022/4/21 22:34

package dao

import (
	"fmt"

	"github.com/go-redis/redis"
)

var Redis *redis.Client

func RedisInit() {
	client := NewRedisClient()
	Redis = client
}

func NewRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(pong)
	return client
}
