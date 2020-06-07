package rediscli

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fbessez/octo-org/config"
	"github.com/gomodule/redigo/redis"
)

// Pool is redis pool instance
var (
	Pool *redis.Pool
)

func newPool(server string) *redis.Pool {
	return &redis.Pool{
		IdleTimeout: 240 * time.Second,
		MaxIdle:     10,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", server, redis.DialConnectTimeout(10*time.Second))
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func cleanupHook() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	signal.Notify(c, syscall.SIGKILL)
	go func() {
		<-c
		Pool.Close()
		os.Exit(0)
	}()
}

func init() {
	address := "127.0.0.1"
	port := "6379"
	if len(config.CONSTANTS.Redis.Address) > 0 {
		address = config.CONSTANTS.Redis.Address
	}
	if len(config.CONSTANTS.Redis.Port) > 0 {
		port = config.CONSTANTS.Redis.Port
	}
	Pool = newPool(address + ":" + port)
	cleanupHook()
}
