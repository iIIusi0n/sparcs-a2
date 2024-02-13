package data

import (
	"database/sql"
	"fmt"
	"log"

	"api-server/config"
	_ "github.com/go-sql-driver/mysql"
)

type ManagerModel interface {
	CreateUser(user *User) (int, error)
	ReadUser(id int) (*User, error)
	UpdateUser(user *User) error
	DeleteUser(id int) error

	CreatePost(post *Post) (int, error)
	ReadPost(id int) (*Post, error)
	ReadPosts() ([]*Post, error)
	UpdatePost(post *Post) error
	DeletePost(id int) error

	ReadPostsByUserID(userID int) ([]*Post, error)
	CountPostsOnUser(userID int) (int, error)

	CreateLike(like *Like) (int, error)
	ReadLike(id int) (*Like, error)
	UpdateLike(like *Like) error
	DeleteLike(id int) error

	CountLikeOnPost(postID int) (int, error)
	CheckUserLikedPost(userID, postID int) (bool, error)

	CreateGathering(gathering *Gathering) (int, error)
	ReadGathering(id int) (*Gathering, error)
	ReadGatheringsByUserID(userID int) ([]*Gathering, error)
	ReadGatherings() ([]*Gathering, error)
	UpdateGathering(gathering *Gathering) error
	DeleteGathering(id int) error

	CountGatheringOnUser(userID int) (int, error)

	CreateParticipant(participant *Participant) (int, error)
	ReadParticipant(id int) (*Participant, error)
	ReadParticipantsByUserID(userID int) ([]*Participant, error)
	ReadParticipantsByGatheringID(gatheringID int) ([]*Participant, error)
	UpdateParticipant(participant *Participant) error
	DeleteParticipant(id int) error

	CountParticipantOnGathering(gatheringID int) (int, error)
	CheckUserParticipatedGathering(userID, gatheringID int) (bool, error)

	CreateGatheringLocation(location *Location) (int, error)
	ReadGatheringLocation(id int) (*Location, error)
	UpdateGatheringLocation(location *Location) error
	DeleteGatheringLocation(id int) error

	ReadGatheringLocationByGatheringID(gatheringID int) (*Location, error)
}

type manager struct {
	db *sql.DB
}

var Manager ManagerModel

func init() {
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
