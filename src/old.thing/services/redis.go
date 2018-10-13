package services

import (
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis"
)

type RedisClient struct {
	*redis.Client
}

var instanceRedisCli *RedisClient = nil

func Redis_Connect() *RedisClient {
	if instanceRedisCli == nil {
		instanceRedisCli = &RedisClient{redis.NewClient(&redis.Options{
			Addr:     os.Getenv("LINK_REDIS"),
			Password: "", // no password set
			DB:       0,  // use default DB
		})}
		pong, err := instanceRedisCli.Ping().Result()
		fmt.Println(pong, err)
	}

	return instanceRedisCli
}

func (client *RedisClient) SetValue(key string, value string, expiration ...interface{}) error {
	exp := 0
	if expiration != nil {
		exp = expiration[0].(int)
	}
	err := client.Set(key, value, time.Duration(time.Second*time.Duration(exp))).Err()
	if err != nil {
		panic(err)
	}
	return err
}

func (client *RedisClient) GetValue(key string) (interface{}, error) {
	val, err := client.Get(key).Result()
	if err != nil {
		return nil, err
	}

	return val, err
}

// func (redisCli *RedisCli) GetValueString(key string) interface{} {
// 	val, err := redisCli.conn.Do("GET", key)
// 	if err == nil {
// 		intArr := val.([]int)
// 		n := bytes.Index(intArr, []byte{0})
// 		return string(intArr[:])
// 	}

// 	return err
// }
