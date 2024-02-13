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

	keyword := c.Query("keyword")
	if keyword != "" {
		keywordPosts := make([]*Post, 0)
		for _, post := range posts {
			if post.Title == keyword || post.Description == keyword {
				keywordPosts = append(keywordPosts, post)
			}
		}
		posts = keywordPosts
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

func CreatePostRouter(c *gin.Context) {
	var post Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	uidAny, _ := c.Get("uid")
	uid := uidAny.(int)
	post.UserID = uid

	id, err := data.Manager.CreatePost(&post)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"id": id})
}

func UpdatePostRouter(c *gin.Context) {
	var post Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	posts, err := data.Manager.ReadPostsByUserID(post.UserID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	uid, _ := c.Get("uid")

	var found bool
	var userAccess bool
	for _, p := range posts {
		if p.ID == post.ID {
			found = true
		}

		if p.UserID == uid.(int) {
			userAccess = true
		}
	}

	if !found {
		c.JSON(400, gin.H{"error": "Post not found"})
		return
	}

	if !userAccess {
		c.JSON(403, gin.H{"error": "User has no access to update this post"})
		return
	}

	err = data.Manager.UpdatePost(&post)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"id": post.ID})
}

func DeletePostRouter(c *gin.Context) {
	id := c.Param("id")

	idNum, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}

	posts, err := data.Manager.ReadPostsByUserID(idNum)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	uid, _ := c.Get("uid")

	var found bool
	var userAccess bool
	for _, p := range posts {
		if p.ID == idNum {
			found = true
		}

		if p.UserID == uid.(int) {
			userAccess = true
		}
	}

	if !found {
		c.JSON(400, gin.H{"error": "Post not found"})
		return
	}

	if !userAccess {
		c.JSON(403, gin.H{"error": "User has no access to delete this post"})
		return
	}

	err = data.Manager.DeletePost(idNum)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"id": idNum})
}

func GetPostLikesRouter(c *gin.Context) {
	id := c.Param("id")

	idNum, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}

	count, err := data.Manager.CountLikeOnPost(idNum)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"count": count})
}

func LikePostRouter(c *gin.Context) {
	id := c.Param("id")

	idNum, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}

	uid, _ := c.Get("uid")

	liked, err := data.Manager.CheckUserLikedPost(uid.(int), idNum)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if liked {
		c.JSON(400, gin.H{"error": "User already liked this post"})
		return
	}

	like := Like{
		UserID: uid.(int),
		PostID: idNum,
	}

	_, err = data.Manager.CreateLike(&like)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"id": idNum})
}

func UnlikePostRouter(c *gin.Context) {
	id := c.Param("id")

	idNum, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}

	uid, _ := c.Get("uid")

	liked, err := data.Manager.CheckUserLikedPost(uid.(int), idNum)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if !liked {
		c.JSON(400, gin.H{"error": "User has not liked this post"})
		return
	}

	like := Like{
		UserID: uid.(int),
		PostID: idNum,
	}

	err = data.Manager.DeleteLike(like.ID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"id": idNum})
}
