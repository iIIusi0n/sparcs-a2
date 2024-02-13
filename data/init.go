package data

import (
	"database/sql"
	"fmt"
	"log"

	"api-server/config"
	"api-server/controllers/gathering"
	"api-server/controllers/post"
	"api-server/controllers/user"

	_ "github.com/go-sql-driver/mysql"
)

type ManagerModel interface {
	CreateUser(user *user.User) (int, error)
	ReadUser(id int) (*user.User, error)
	UpdateUser(user *user.User) error
	DeleteUser(id int) error

	CreatePost(post *post.Post) (int, error)
	ReadPost(id int) (*post.Post, error)
	ReadPosts() ([]*post.Post, error)
	UpdatePost(post *post.Post) error
	DeletePost(id int) error

	ReadPostsByUserID(userID int) ([]*post.Post, error)
	CountPostsOnUser(userID int) (int, error)

	CreateLike(like *post.Like) (int, error)
	ReadLike(id int) (*post.Like, error)
	UpdateLike(like *post.Like) error
	DeleteLike(id int) error

	CountLikeOnPost(postID int) (int, error)
	CheckUserLikedPost(userID, postID int) (bool, error)

	CreateGathering(gathering *gathering.Gathering) (int, error)
	ReadGathering(id int) (*gathering.Gathering, error)
	ReadGatheringsByUserID(userID int) ([]*gathering.Gathering, error)
	ReadGatherings() ([]*gathering.Gathering, error)
	UpdateGathering(gathering *gathering.Gathering) error
	DeleteGathering(id int) error

	CountGatheringOnUser(userID int) (int, error)

	CreateParticipant(participant *gathering.Participant) (int, error)
	ReadParticipant(id int) (*gathering.Participant, error)
	ReadParticipantsByUserID(userID int) ([]*gathering.Participant, error)
	ReadParticipantsByGatheringID(gatheringID int) ([]*gathering.Participant, error)
	UpdateParticipant(participant *gathering.Participant) error
	DeleteParticipant(id int) error

	CountParticipantOnGathering(gatheringID int) (int, error)
	CheckUserParticipatedGathering(userID, gatheringID int) (bool, error)

	CreateGatheringLocation(location *gathering.Location) (int, error)
	ReadGatheringLocation(id int) (*gathering.Location, error)
	UpdateGatheringLocation(location *gathering.Location) error
	DeleteGatheringLocation(id int) error

	ReadGatheringLocationByGatheringID(gatheringID int) (*gathering.Location, error)
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
