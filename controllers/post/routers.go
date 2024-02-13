package post

import (
	"strconv"

	"api-server/data"
	"github.com/gin-gonic/gin"
)

func GetPostRouter(c *gin.Context) {
	id := c.Param("id")

	idNum, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}

	post, err := data.Manager.ReadPost(idNum)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, post)
}

func GetPostsRouter(c *gin.Context) {
	var uid int
	if c.Query("uid") != "" {
		uidNum, err := strconv.Atoi(c.Query("uid"))
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid uid"})
			return
		}
		uid = uidNum
	} else {
		uidAny, _ := c.Get("uid")
		uid = uidAny.(int)
	}

	sortMethod := c.Query("sort")

	posts, err := data.Manager.ReadPostsByUserID(uid)
	switch sortMethod {
	case SortOptionLiked:
		posts = SortPostsByLikes(posts)
	case SortOptionNearby:
		latitude, _ := strconv.ParseFloat(c.Query("latitude"), 64)
		longitude, _ := strconv.ParseFloat(c.Query("longitude"), 64)
		posts = SortPostsByDistance(posts, latitude, longitude)
	case SortOptionLatest:
	}

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	limit := c.Query("limit")
	if limit != "" {
		limitNum, err := strconv.Atoi(limit)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid limit"})
			return
		}
		if limitNum > len(posts) {
			limitNum = len(posts)
		}
		posts = posts[:limitNum]
	}

	c.JSON(200, posts)
}
