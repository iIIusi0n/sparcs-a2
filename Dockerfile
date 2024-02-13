FROM golang:1.19

WORKDIR /app

RUN go build api-server/cmd/api-server
