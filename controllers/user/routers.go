package user

import (
	"api-server/config"
	"api-server/controllers/auth"
	"github.com/gin-gonic/gin"
)

func TemporaryTokenRouter(c *gin.Context) {
	id := c.Param("id")
	name := c.Param("name")

	token, err := auth.BuildNewToken(id, name)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate token"})
		return
	}

	c.SetCookie("token", token, 60*60*24, "/", config.ServerName, false, true)

	c.JSON(200, gin.H{"token": token})
}
