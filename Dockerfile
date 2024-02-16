FROM golang:1.19

WORKDIR /app

COPY . .

ENV LANG en_US.UTF-8
ENV LANGUAGE en_US:en
ENV LC_ALL en_US.UTF-8

RUN go get -d -v ./...
RUN go build api-server/cmd/api-server

CMD ["./api-server"]
