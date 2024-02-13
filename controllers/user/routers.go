package user

import (
	"strconv"

	"api-server/config"
	"api-server/controllers/auth"
	"api-server/data"
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

func GetLoggedInUserRouter(c *gin.Context) {
	uid, _ := c.Get("uid")

	user, err := data.Manager.ReadUser(uid.(int))
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to read user"})
		return
	}

	c.JSON(200, user)
}

func GetStatsRouter(c *gin.Context) {
	uid, _ := c.Get("uid")

	totalLikes, err := getUserTotalLikes(uid.(int))
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get total likes"})
		return
	}

	totalGatherings, err := getUserTotalGatherings(uid.(int))
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get total gatherings"})
		return
	}

	totalPosts, err := getUserTotalPosts(uid.(int))
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get total posts"})
		return
	}

	c.JSON(200, gin.H{
		"total_likes":      totalLikes,
		"total_gatherings": totalGatherings,
		"total_posts":      totalPosts,
	})
}
