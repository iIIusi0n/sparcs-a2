package chat

import (
	"api-server/data"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func GetChatRoomsRouter(c *gin.Context) {
	rooms, err := data.Manager.ReadChatRooms()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to read chat rooms"})
		return
	}

	c.JSON(200, rooms)
}

func GetChatRoomRouter(c *gin.Context) {
	id := c.Param("id")
	idNum, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}

	room, err := data.Manager.ReadChatRoom(idNum)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to read chat room"})
		return
	}

	c.JSON(200, room)
}

func CreateChatRoomRouter(c *gin.Context) {
	var room data.ChatRoom
	err := c.BindJSON(&room)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	id, err := data.Manager.CreateChatRoom(&room)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create chat room"})
		return
	}

	c.JSON(200, gin.H{"id": id})
}

func GetChatMessagesRouter(c *gin.Context) {
	id := c.Param("room_id")
	idNum, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}

	messages, err := data.Manager.ReadChatsByRoomID(idNum)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to read chat messages"})
		return
	}

	if c.Query("limit") != "" {
		limit, err := strconv.Atoi(c.Query("limit"))
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid limit"})
			return
		}
		if limit < len(messages) {
			messages = messages[:limit]
		}
	}

	c.JSON(200, messages)
}

func CreateChatMessageRouter(c *gin.Context) {
	id := c.Param("room_id")
	idNum, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}

	var chat data.Chat
	err = c.BindJSON(&chat)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	chat.RoomID = idNum
	chat.SenderID = c.GetInt("uid")

	chatID, err := data.Manager.CreateChat(&chat)
	if err != nil {
		log.Println(err)

		c.JSON(500, gin.H{"error": "Failed to create chat message"})
		return
	}

	c.JSON(200, gin.H{"id": chatID})
}
