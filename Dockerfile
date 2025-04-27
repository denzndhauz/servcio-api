# Dockerfile for Ecommerce API
FROM golang:1.23.0-alpine3.18

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o main ./cmd/server

CMD ["./main"]
