package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func main() {
	// チャネルを作成する
	c := make(chan int)

	// 100回ループしてgoroutineを起動する
	for i := 0; i < 100; i++ {
		time.Sleep(10 * time.Millisecond)
		go func() {
			// TCPサーバーに接続する
			conn, err := net.Dial("tcp", "tcp-test1.verification-gcp.colopl.jp.:30000")
			if err != nil {
				fmt.Println(err)
				c <- 1
				return
			}

			// サーバーにデータを送信する
			message := "Hello, World " + strconv.Itoa(i)
			conn.Write([]byte(message))

			// サーバーからのレスポンスを受信する
			buf := make([]byte, 1024)
			n, err := conn.Read(buf)
			if err != nil {
				fmt.Println(err)
				c <- 1
				return
			}

			fmt.Println(string(buf[:n]))

			// チャネルに値を送信する
			c <- 1
		}()
	}

	// 全てのgoroutineが終了するまで待機する
	for i := 0; i < 10000; i++ {
		<-c
	}
}
