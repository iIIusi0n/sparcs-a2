package image

import (
	"fmt"
	"log"
	"os"

	"api-server/config"
	"github.com/gin-gonic/gin"
)

func GetImageRouter(c *gin.Context) {
	id := c.Param("id")
	path := fmt.Sprintf("%s/%s", config.ImageDirectory, id)
	_, err := os.Stat(path)
	if err != nil {
		log.Println("File not found: ", err)

		c.JSON(404, gin.H{"error": "File not found"})
		return
	}

	c.Header("Content-Type", "image/all")
	c.File(path)
}

func UploadImageRouter(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		log.Println("No file found: ", err)

		c.JSON(400, gin.H{"error": "No file found"})
		return
	}

	path, id := newRandomImagePath()
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		log.Println("Failed to save file: ", err)

		c.JSON(500, gin.H{"error": "Failed to save file"})
		return
	}

	c.JSON(200, gin.H{"path": path, "id": id})
}

func DeleteImageRouter(c *gin.Context) {
	id := c.Param("id")
	path := fmt.Sprintf("%s/%s", config.ImageDirectory, id)
	err := os.Remove(path)
	if err != nil {
		log.Println("Failed to delete file: ", err)

		c.JSON(500, gin.H{"error": "Failed to delete file"})
		return
	}

	c.JSON(200, gin.H{"id": id})
}
