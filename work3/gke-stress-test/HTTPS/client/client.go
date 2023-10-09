package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

func main() {
	// 並行処理用のセマフォを作成します。
	loop := 10000
	sem := make(chan bool, loop)
	wg := sync.WaitGroup{}

	// 1000 個のゴルーチンを作成します。
	for i := 0; i < loop; i++ {
		wg.Add(1)
		time.Sleep(20 * time.Millisecond)
		go func() {
			defer wg.Done()

			// セマフォを取得します。
			sem <- true

			// HTTP リクエストを行います。
			//resp, err := http.Get("http://104.196.247.155:8080")
			//resp, err := http.Get("https://haproxy2-8-juulzjpopq-uc.a.run.app")
			//resp, err := http.Get("https://haproxy2-8-1-juulzjpopq-uc.a.run.app")
			//resp, err := http.Get("http://34.36.1.70")
			//resp, err := http.Get("http://130.211.11.61") // lb
			//resp, err := http.Get("http://34.132.121.47:8080") // us-eda-svc
			//resp, err := http.Get("http://split-test1.verification-gcp.colopl.jp.")
			resp, err := http.Get("http://active-standby1.verification-gcp.colopl.jp.")
			if err != nil {
				fmt.Println(err)
				return
			}
			defer resp.Body.Close()

			// 応答を読み取ります。
			b, err := io.ReadAll(resp.Body)
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
