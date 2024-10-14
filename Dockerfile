FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

ENV GOPROXY https://goproxy.cn,direct
RUN go mod download

COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN  go build -o start ./logToFile/start.go
