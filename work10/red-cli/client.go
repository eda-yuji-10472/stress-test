package main

import (
	"context"
	goredis "github.com/redis/go-redis/v9"
	"log"
	"strconv"
	"time"
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

	log.Println("Start", time.Now())

	// コマンドの組み立て
	//cmd := redis.Set(ctx, "key", "value", goredis.KeepTTL)
	//hashes := []string{"{hash1}:", "{hash2}:"}
	//keys := []string{"key1", "key2"}
	//values := []string{"value1", "value2"}

	// 1000件のキーと値を設定
	keys := make([]string, 1000)
	values := make([]string, 1000)
	//var ifaces []interface{}
	var ifaces1 []interface{}
	var ifaces2 []interface{}
	var ifaces3 []interface{}
	var ifaces4 []interface{}
	var ifaces5 []interface{}
	var ifaces6 []interface{}
	var ifaces7 []interface{}
	var ifaces8 []interface{}
	//ifaces := make([]interface{},1000)

	for i := 0; i < 1000; i++ {
		keys[i] = "{hash1}:" + "key1" + strconv.Itoa(i)
		values[i] = "value1" + strconv.Itoa(i)
		ifaces1 = append(ifaces1, keys[i], values[i])
	}
	for i := 0; i < 1000; i++ {
		keys[i] = "{hash2}:" + "key2" + strconv.Itoa(i)
		values[i] = "value2" + strconv.Itoa(i)
		ifaces2 = append(ifaces2, keys[i], values[i])
	}
	for i := 0; i < 1000; i++ {
		keys[i] = "{hash3}:" + "key3" + strconv.Itoa(i)
		values[i] = "value3" + strconv.Itoa(i)
		ifaces3 = append(ifaces3, keys[i], values[i])
	}
	for i := 0; i < 1000; i++ {
		keys[i] = "{hash4}:" + "key4" + strconv.Itoa(i)
		values[i] = "value4" + strconv.Itoa(i)
		ifaces4 = append(ifaces4, keys[i], values[i])
	}
	for i := 0; i < 1000; i++ {
		keys[i] = "{hash5}:" + "key5" + strconv.Itoa(i)
		values[i] = "value5" + strconv.Itoa(i)
		ifaces5 = append(ifaces5, keys[i], values[i])
	}
	for i := 0; i < 1000; i++ {
		keys[i] = "{hash6}:" + "key6" + strconv.Itoa(i)
		values[i] = "value6" + strconv.Itoa(i)
		ifaces6 = append(ifaces6, keys[i], values[i])
	}
	for i := 0; i < 1000; i++ {
		keys[i] = "{hash7}:" + "key7" + strconv.Itoa(i)
		values[i] = "value7" + strconv.Itoa(i)
		ifaces7 = append(ifaces7, keys[i], values[i])
	}
	for i := 0; i < 1000; i++ {
		keys[i] = "{hash8}:" + "key8" + strconv.Itoa(i)
		values[i] = "value8" + strconv.Itoa(i)
		ifaces8 = append(ifaces8, keys[i], values[i])
	}
	//exp := time.Duration(600 * time.Second) // 10 minutes
	//var ifaces []interface{}
	//pipe := redis.TxPipeline()
	//for i := range keys {
	//	ifaces = append(ifaces, keys[i], values[i])
	//pipe.Expire(keys[i], exp)
	//}
	cmd := redis.MSet(ctx, ifaces1...)

	//cmd := redis.Set(ctx, "key1", "value1")
	for i := 0; i < 10000; i++ {
		cmd = redis.MSet(ctx, ifaces1...)
		// 実行(保存系：結果が不要な場合)
		if err := cmd.Err(); err != nil {
			panic(err)
		}
		cmd = redis.MSet(ctx, ifaces2...)
		// 実行(保存系：結果が不要な場合)
		if err := cmd.Err(); err != nil {
			panic(err)
		}
		cmd = redis.MSet(ctx, ifaces3...)
		// 実行(保存系：結果が不要な場合)
		if err := cmd.Err(); err != nil {
			panic(err)
		}
		cmd = redis.MSet(ctx, ifaces4...)
		// 実行(保存系：結果が不要な場合)
		if err := cmd.Err(); err != nil {
			panic(err)
		}
		cmd = redis.MSet(ctx, ifaces5...)
		// 実行(保存系：結果が不要な場合)
		if err := cmd.Err(); err != nil {
			panic(err)
		}
		cmd = redis.MSet(ctx, ifaces6...)
		// 実行(保存系：結果が不要な場合)
		if err := cmd.Err(); err != nil {
			panic(err)
		}
		cmd = redis.MSet(ctx, ifaces7...)
		// 実行(保存系：結果が不要な場合)
		if err := cmd.Err(); err != nil {
			panic(err)
		}
		cmd = redis.MSet(ctx, ifaces8...)
		// 実行(保存系：結果が不要な場合)
		if err := cmd.Err(); err != nil {
			panic(err)
		}

		pipe := redis.Pipeline()
		for _, key := range keys {
			pipe.Del(ctx, key)
		}
		pipe.Exec(ctx)

	}

	log.Println("End", time.Now())

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
