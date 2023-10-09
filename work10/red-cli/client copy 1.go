package main

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

func newPool(server string, password string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp",
				server,
				redis.DialPassword(password))
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

// Connection
func Connection() redis.Conn {
	const Addr = "my-redis1-redis-cluster:6379"

	c, err := redis.Dial("tcp", Addr, redis.DialPassword("password"))
	if err != nil {
		panic(err)
	}
	return c
}

type Data struct {
	Key   string
	Value string
}

// 複数のデータの登録(Redis: MSET key [key...])
func Mset(datas []Data, c redis.Conn) {
	var query []interface{}
	for _, v := range datas {
		query = append(query, v.Key, v.Value)
	}
	fmt.Println(query) // [key1 value1 key2 value2]

	c.Do("MSET", query...)
}

// 複数の値を取得 (Redis: MGET key [key...])
func Mget(keys []string, c redis.Conn) []string {
	var query []interface{}
	for _, v := range keys {
		query = append(query, v)
	}
	fmt.Println("MGET query:", query) // [key1 key2]

	res, err := redis.Strings(c.Do("MGET", query...))
	if err != nil {
		panic(err)
	}
	return res
}

// TTLの設定(Redis: EXPIRE key ttl)
func Expire(key string, ttl int, c redis.Conn) {
	c.Do("EXPIRE", key, ttl)
}

func main() {
	// 接続
	c := Connection()
	defer c.Close()

	// 複数データの登録
	datas := []Data{
		Data{Key: "key1", Value: "value1"},
		Data{Key: "key2", Value: "value2"},
	}
	Mset(datas, c)

	// 複数データの取得
	keys := []string{"key1", "key2"}
	res_mget := Mget(keys, c)
	fmt.Println(res_mget)

	// TTLの設定
	Expire("key1", 10, c)
}
