package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	// サーバーをポート8080でリッスンする
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// リクエストヘッダーを取得する
		headers := r.Header

		// リクエストヘッダーを表示する
		fmt.Println("リクエストヘッダー：")
		for k, v := range headers {
			fmt.Println(k, ":", strings.Join(v, ","))
		}

		// 200 OKレスポンスを返す
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("gke-usr-OK"))
	})

	// サーバーを起動する
	log.Fatal(http.ListenAndServe(":8080", nil))
}
