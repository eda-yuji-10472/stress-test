package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

// クラスタと非クラスタを一律に扱うためのインタフェース
type universalClient interface {
	Ping() *redis.StatusCmd
	Get(key string) *redis.StringCmd
	Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd
}

// 接続だけはクラスタと非クラスタで異なる
func newSingleClient(addr string) universalClient {
	r := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	return r
}

func newClusterClient(addr string) universalClient {
	r := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    []string{addr},
		Password: "password",
	})
	return r
}

// ここから実際の使い方
func main() {
	// シングル構成
	/*	fmt.Println("---------- Single")
			r1 := newSingleClient("localhost:6379")
			fmt.Println(r1.Ping())
			fmt.Println(r1.Get("key1"))
			fmt.Println(r1.Set("key1", "value1", 30*time.Second))
			fmt.Println(r1.Get("key1"))
		// クラスタ構成
		fmt.Println("\n---------- Cluster")
		r2 := newClusterClient("my-redis1-redis-cluster:6379")
		fmt.Println(r2.Ping())
		fmt.Println(r2.Get("key1"))
		for i := 0; i < 10000; i++ {
			r2.Set("key0", "value0", 30*time.Second)
			r2.Set("key1", "value1", 30*time.Second)
			r2.Set("key2", "value2", 30*time.Second)
			r2.Set("key3", "value3", 30*time.Second)
			r2.Set("key4", "value4", 30*time.Second)
			r2.Set("key5", "value5", 30*time.Second)
			r2.Set("key6", "value6", 30*time.Second)
			r2.Set("key7", "value7", 30*time.Second)
			r2.Set("key8", "value8", 30*time.Second)
			r2.Set("key9", "value9", 30*time.Second)
		}
		fmt.Println(r2.Get("key1"))
	*/

	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    []string{"my-redis1-redis-cluster:6379"},
		Password: "password",
		PoolSize: 64,
	})

	// 1000件のキーと値を設定
	keys := make([]string, 1000)
	values := make([]string, 1000)
	var ifaces []interface{}
	for i := 0; i < 1000; i++ {
		keys[i] = "{hash1}:key" + strconv.Itoa(i)
		values[i] = "value" + strconv.Itoa(i)
		ifaces = append(ifaces, keys[i], values[i])
	}
	fmt.Println(keys)
	for i := 0; i < 1000; i++ {
		err := client.MSet(ifaces...).Err()
		if err != nil {
			log.Fatal(err)
		}
	}
}
