package user

import (
	"strconv"

	"api-server/config"
	"api-server/controllers/auth"
	"github.com/gin-gonic/gin"
)

func TemporaryTokenRouter(c *gin.Context) {
	id := c.Param("id")
	name := c.Param("name")

	idNumber, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}

	token, err := auth.BuildNewToken(idNumber, name)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate token"})
		return
	}

	c.SetCookie("token", token, 60*60*24, "/", config.ServerName, false, true)

	c.JSON(200, gin.H{"token": token})
}
