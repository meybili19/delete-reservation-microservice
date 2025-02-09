FROM golang:1.23.4

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN go build -o delete-reservation cmd/main.go

EXPOSE 4002

CMD ["./delete-reservation"]
