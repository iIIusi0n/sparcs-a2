package main

import (
	"math/rand"
	"os"
	"time"

	"api-server/config"
	"api-server/server"
)

func ReadEnvorinmentVariables() {
	config.ServerName = os.Getenv("SERVER_DOMAIN")
	config.ServerSecret = os.Getenv("SERVER_SECRET_KEY")

	config.MysqlUsername = os.Getenv("MYSQL_USER")
	config.MysqlPassword = os.Getenv("MYSQL_PASSWORD")
	config.MysqlDatabase = os.Getenv("MYSQL_DATABASE")
}

func main() {
	ReadEnvorinmentVariables()

	rand.Seed(time.Now().Unix())

	err := server.Start()
	if err != nil {
		panic(err)
	}
}
