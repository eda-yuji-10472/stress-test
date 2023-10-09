package main

import (
	"context"
	goredis "github.com/redis/go-redis/v9"
	"log"
)

/*
func newClusterClient(addr string) {
	r := goredis.NewClusterClient(&goredis.ClusterOptions{
		Addrs:    []string{addr},
		Password: "password",
	})
	return r
}
*/

func main() {
	ctx := context.Background()

	// クライアントの初期化
	//redis := newClusterClient("my-redis1-redis-cluster:6379")
	redis := goredis.NewClusterClient(&goredis.ClusterOptions{Addrs: []string{"my-redis1-redis-cluster:6379"}, Password: "password"})

	// 疎通確認
	if err := redis.Ping(ctx).Err(); err != nil {
		panic(err)
	}

	defer redis.Close()

	// コマンドの組み立て
	//cmd := redis.Set(ctx, "key", "value", goredis.KeepTTL)
	hashes :=  []string{"{hash1}:", "{hash2}:"}
	keys := []string{"key1", "key2"}
	values := []string{"value1", "value2"}

	//exp := time.Duration(600 * time.Second) // 10 minutes
	var ifaces []interface{}
	//pipe := redis.TxPipeline()
	for i := range keys {
		ifaces = append(ifaces, hashes[0]+keys[i], values[i])
		//pipe.Expire(keys[i], exp)
	}

	//cmd := redis.Set(ctx, "key1", "value1")
	cmd := redis.MSet(ctx, ifaces...)

	// 実行(取得系：結果が必要な場合)
	if val, err := cmd.Result(); err != nil {
		panic(err)
	} else {
		log.Printf("%s", val)
	}

	// 実行(保存系：結果が不要な場合)
	if err := cmd.Err(); err != nil {
		panic(err)
	}
}
