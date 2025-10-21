FROM golang:1.25
WORKDIR /app

COPY go.mod go.sum server.go ./
RUN go mod download

EXPOSE 80

RUN go build -o server 

CMD ["./server"]
