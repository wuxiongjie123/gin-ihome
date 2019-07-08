package grides

import "github.com/gomodule/redigo/redis"

var RedisConn *redis.Pool

func SetUp() error {
	RedisConn = &redis.Pool{
		
	}
	return nil
}