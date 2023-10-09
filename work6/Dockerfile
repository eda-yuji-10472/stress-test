# Go 1.16をベースとするDockerイメージを使用します
FROM golang:1.16

# 作業ディレクトリを作成し、コピー元のファイルを追加します
WORKDIR /app
COPY server.go .
RUN mkdir /file

# Goプログラムをビルドします
RUN go build server.go

# ポート8080を開放します
EXPOSE 8080

# イメージを実行した時のデフォルトコマンドを設定します
CMD ["./server"]

VOLUME ["/file"]