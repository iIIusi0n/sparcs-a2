package data

import (
	"api-server/data"
	"crypto/sha512"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"log"
)

func hashSha512(str string) string {
	hash := sha512.New()
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
}

func DropAllDataRouter(c *gin.Context) {
	key := c.Query("key")
	if hashSha512(key) != "448cfd985ef0af6917028c44ea3e116caf178ca8d3dca4b7b3d6bcd0994ad6162bc9d883042b903e745501893ace531b8cbc79ef7b03af5f5552526c2b2a494d" {
		c.JSON(400, gin.H{"error": "Invalid key"})
		return
	}

	err := data.Manager.DropAllData()
	if err != nil {
		log.Println(err)

		c.JSON(500, gin.H{"error": "Failed to drop all data"})
		return
	}

	c.JSON(200, gin.H{"message": "All data has been dropped"})
}
