FROM golang:1.20

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN go build -o main .

EXPOSE 40020

CMD ["./main"]
