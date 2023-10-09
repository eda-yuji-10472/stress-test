package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	// チャネルを作成する
	c := make(chan int)
	// 100回ループしてクライアントを起動する
	for i := 0; i < 10000; i++ {
		time.Sleep(10 * time.Millisecond)
		go func() {
			// クライアントを起動する
			conn, err := net.Dial("udp", "35.203.184.172:8080")
			if err != nil {
				fmt.Println(err)
				c <- 1
				return
			}
			defer conn.Close()

			// サーバーにデータを送信する
			_, err = conn.Write([]byte("Hello from client"))
			if err != nil {
				fmt.Println(err)
				c <- 1
				return
			}

			// サーバーからの応答を受信する
			buf := make([]byte, 1024)
			n, err := conn.Read(buf)
			if err != nil {
				fmt.Println(err)
				c <- 1
				return
			}

			// 受信したデータを表示する
			fmt.Printf("Received data from server: %s\n", string(buf[:n]))

			// チャネルに値を送信する
			c <- 1
		}()
	}

	// 全てのgoroutineが終了するまで待機する
	for i := 0; i < 10000; i++ {
		<-c
	}
}
