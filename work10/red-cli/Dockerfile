FROM golang:1.21

WORKDIR /app

#COPY go.mod ./
#COPY go.sum ./
RUN go mod init redis-test
RUN go get github.com/go-redis/redis
RUN go get "github.com/gomodule/redigo/redis"
RUN go get "github.com/redis/go-redis/v9"

COPY client.go ./

RUN go build -o client ./client.go

CMD ["/app/client"]