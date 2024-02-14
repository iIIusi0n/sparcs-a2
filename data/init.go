package data

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"api-server/config"
	_ "github.com/go-sql-driver/mysql"
)

type ManagerModel interface{}

type manager struct {
	db *sql.DB
}

var Manager ManagerModel

func init() {
	config.MysqlUsername = os.Getenv("MYSQL_USER")
	config.MysqlPassword = os.Getenv("MYSQL_PASSWORD")
	config.MysqlDatabase = os.Getenv("MYSQL_DATABASE")

	db, err := sql.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
			config.MysqlUsername,
			config.MysqlPassword,
			config.MysqlHostname,
			config.MysqlPort,
			config.MysqlDatabase,
		))
	if err != nil {
		log.Println("Failed to connect to database: ", err)

		panic(err)
	}

	Manager = &manager{db: db}
}
