FROM golang:1.21

WORKDIR /app

#COPY go.mod ./
#COPY go.sum ./
#RUN go mod download

COPY client.go ./

RUN go build -o client ./client.go

CMD ["/app/client"]