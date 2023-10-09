package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	// ポート番号を指定してリッスン開始
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer listener.Close()

	fmt.Println("Listening on port 8080")

	for {
		// 接続を待ちます
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting:", err.Error())
			continue
		}
		// ゴルーチンで接続を処理します
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	// コネクションからメッセージを受信します
	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
		return
	}

	// メッセージを出力します
	fmt.Printf("Received message: %s\n", string(buffer))

	// 10秒間待機する
	time.Sleep(100 * time.Second)

	// メッセージを返します
	_, err = conn.Write([]byte("Message received.\n"))
	if err != nil {
		fmt.Println("Error writing:", err.Error())
	}

	// コネクションを閉じます
	conn.Close()
}
