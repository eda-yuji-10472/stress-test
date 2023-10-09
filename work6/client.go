package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

func main() {
	// 並行処理用のセマフォを作成します。
	loop := 100000
	sem := make(chan bool, loop)
	wg := sync.WaitGroup{}

	// 10000個のゴルーチンを作成します。
	for i := 0; i < loop; i++ {
		wg.Add(1)
		time.Sleep(10 * time.Millisecond)
		go func() {
			defer wg.Done()

			// セマフォを取得します。
			sem <- true

			// HTTP リクエストを行います。
			resp, err := http.Get("http://104.196.247.155:8080")
			//resp, err := http.Get("http://localhost:8080")
			if err != nil {
				fmt.Println(err)
				return
			}
			defer resp.Body.Close()

			// 応答を読み取ります。
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println(err)
				return
			}

			// 応答を出力します。
			fmt.Println(string(b))

			// セマフォを解放します。
			<-sem
		}()
	}

	// すべてのゴルーチンが終了するのを待ちます。
	wg.Wait()
}
