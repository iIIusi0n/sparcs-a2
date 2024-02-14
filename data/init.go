package data

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"api-server/config"
	_ "github.com/go-sql-driver/mysql"
)

type ManagerModel interface {
	CreateUser(user *User) (int, error)
	ReadUser(id int) (*User, error)
	UpdateUser(user *User) error
	DeleteUser(id int) error

	CreateHospital(hospital *Hospital) (int, error)
	ReadHospital(id int) (*Hospital, error)
	UpdateHospital(hospital *Hospital) error
	DeleteHospital(id int) error

	CreateInHospital(inHospital *InHospital) error
	ReadInHospital(hospitalID, userID int) (*InHospital, error)
	UpdateInHospital(inHospital *InHospital) error
	DeleteInHospital(hospitalID, userID int) error

	CreateWaitingNumber(waitingNumber *WaitingNumber) error
	ReadWaitingNumber(hospitalID, userID int) (*WaitingNumber, error)
	UpdateWaitingNumber(waitingNumber *WaitingNumber) error
	DeleteWaitingNumber(hospitalID, userID int) error

	CreateChatRoom(chatRoom *ChatRoom) (int, error)
	ReadChatRoom(id int) (*ChatRoom, error)
	UpdateChatRoom(chatRoom *ChatRoom) error
	DeleteChatRoom(id int) error

	CreateChat(chat *Chat) (int, error)
	ReadChat(id int) (*Chat, error)
	UpdateChat(chat *Chat) error
	DeleteChat(id int) error

	CreateChatRead(chatRead *ChatRead) error
	ReadChatRead(chatID, userID int) (*ChatRead, error)
	UpdateChatRead(chatRead *ChatRead) error
	DeleteChatRead(chatID, userID int) error
}

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
