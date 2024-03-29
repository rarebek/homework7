FROM golang:1.20-alpine3.16 AS builder

RUN apk add --no-cache git

RUN mkdir app
COPY . /app

WORKDIR /app

# RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

RUN go build -o main cmd/main.go

CMD ["/app/main"]


