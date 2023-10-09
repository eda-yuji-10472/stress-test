package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	// サーバーをポート8080でリッスンする
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// リクエストヘッダーを取得する
		headers := r.Header

		var env string = os.Getenv("HOSTNAME")

		filePath := "/file/" + env + "random.txt"

		// リクエストヘッダーを表示する
		fmt.Println("リクエストヘッダー：")
		for k, v := range headers {
			fmt.Println(k, ":", strings.Join(v, ","))
		}

		// 1MBのランダム文字列を生成する
		//b := make([]byte, 1024*1024)
		b := make([]byte, 64*1024)
		rand.Read(b)

		// ファイルが存在する場合は追記する
		//f, err := os.OpenFile("/file/random.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		f, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer f.Close()

		// Get the file size.
		fileSize, err := os.Stat(filePath)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Check if the file size is larger than 10GB.
		if fileSize.Size() > 10*1024*1024*1024 {
			//if fileSize.Size() > 10*1024*1024 {
			// Delete the file.
			err = os.Remove(filePath)
			if err != nil {
				fmt.Println(err)
				return
			}

			// Recreate the file.
			fi, err := os.Create(filePath)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer fi.Close()
		}

		_, err = f.Write(b)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// 200 OKレスポンスを返す
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// サーバーを起動する
	log.Fatal(http.ListenAndServe(":8080", nil))

}
