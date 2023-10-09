package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func main() {
	// ドメイン名のリスト
	hostname := "open-match-swaggerui.open-match.svc.clusterset.local"

	// 同期用オブジェクト
	var wg sync.WaitGroup

	// 並列処理
	for i := 0; i < 10000; i++ {
		time.Sleep(20 * time.Millisecond)
		wg.Add(1)
		go func(hostname string) {
			// DNS問い合わせを実行
			addrs, err := net.LookupHost(hostname)
			if err != nil {
				fmt.Println(err)
				//return
			}

			// 応答を表示
			fmt.Println(addrs)

			// 終了通知
			wg.Done()
		}(hostname)
	}

	// すべての処理が終了するまで待機
	wg.Wait()
}
