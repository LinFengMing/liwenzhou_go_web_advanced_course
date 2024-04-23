package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		PoolSize: 100,
	})
	_, err = rdb.Ping().Result()
	return err
}

func watchDemo() {
	key := "watch_count"
	err := rdb.Watch(func(tx *redis.Tx) error {
		n, err := tx.Get(key).Int()
		if err != nil && err != redis.Nil {
			return err
		}
		_, err = tx.TxPipelined(func(pipe redis.Pipeliner) error {
			time.Sleep(10 * time.Second)
			pipe.Set(key, n+1, 0)
			return nil
		})
		return err
	}, key)
	if err != nil {
		fmt.Printf("tx exec failed, err:%v\n", err)
		return
	}
	fmt.Println("tx exec success")
}

func main() {
	if err := initClient(); err != nil {
		fmt.Printf("init redis client failed, err:%v\n", err)
	}
	fmt.Println("init redis success")
	defer rdb.Close()
	watchDemo()
}
