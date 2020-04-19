package rediscli

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

// Get gets value from redis
func Get(key string) (string, error) {
	conn := Pool.Get()
	defer conn.Close()

	return redis.String(conn.Do("GET", key))
}

// Set sets key-value pair in redis with expiration time
func Set(key, value string, exp int32) error {
	conn := Pool.Get()
	defer conn.Close()

	// set value to redis expired after certain time
	return conn.Send("SET", key, value, "NX", "EX", exp)
}

func SetAdd(key string, value string) error {
	conn := Pool.Get()
	defer conn.Close()

	// Add value to the set at KEY
	return conn.Send("SADD", key, value)
}

func GetSetMembers(key string) (members []string, err error) {
	conn := Pool.Get()
	defer conn.Close()

	// Get values in set at given KEY
	s, err := redis.Strings(conn.Do("SMEMBERS", key))
	if err != nil {
		fmt.Printf("ERROR: fail get set %s , error %s", key, err.Error())
		return nil, err
	}

	return s, nil
}

func Invalidate(key string) error {
	conn := Pool.Get()
	defer conn.Close()

	return conn.Send("DEL", key)
}