package gathering

import (
	"strconv"
	"strings"
	"time"

	"api-server/config"
	"api-server/data"
	"github.com/gin-gonic/gin"
)

func GetGatheringRouter(c *gin.Context) {
	id := c.Param("id")

	idNum, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}

	gathering, err := data.Manager.ReadGathering(idNum)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gathering)
}

func GetGatheringsRouter(c *gin.Context) {
	sortMethod := c.Query("sort")

	gatherings, err := data.Manager.ReadGatherings()
	switch sortMethod {
	case SortOptionLiked:
		gatherings = SortGatheringsByLikes(gatherings)
	case SortOptionNearby:
		latitude, _ := strconv.ParseFloat(c.Query("latitude"), 64)
		longitude, _ := strconv.ParseFloat(c.Query("longitude"), 64)
		gatherings = SortGatheringsByDistance(gatherings, latitude, longitude)
	case SortOptionLatest:
	}

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	keyword := c.Query("keyword")
	if keyword != "" {
		keywordGatherings := make([]*data.Gathering, 0)
		for _, gathering := range gatherings {
			if strings.Contains(gathering.Title, keyword) {
				keywordGatherings = append(keywordGatherings, gathering)
			}
		}
		gatherings = keywordGatherings
	}

	limit := c.Query("limit")
	if limit != "" {
		limitNum, err := strconv.Atoi(limit)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid limit"})
			return
		}
		if limitNum < len(gatherings) {
			gatherings = gatherings[:limitNum]
		}
	}

	c.JSON(200, gatherings)
}

func CreateGatheringRouter(c *gin.Context) {
	var gathering data.Gathering
	err := c.BindJSON(&gathering)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id, err := data.Manager.CreateGathering(&gathering)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	participant := data.Participant{
		UserID:      gathering.UserID,
		GatheringID: gathering.ID,
	}
	_, err = data.Manager.CreateParticipant(&participant)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"id": id})
}

func UpdateGatheringRouter(c *gin.Context) {
	var gathering data.Gathering
	err := c.BindJSON(&gathering)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = data.Manager.UpdateGathering(&gathering)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"id": gathering.ID})
}

func DeleteGatheringRouter(c *gin.Context) {
	id := c.Param("id")

	idNum, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}

	err = data.Manager.DeleteGathering(idNum)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"id": idNum})
}

func GetGatheringParticipantsRouter(c *gin.Context) {
	id := c.Param("id")

	idNum, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}

	var result []int
	participants, err := data.Manager.ReadParticipantsByGatheringID(idNum)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	for _, participant := range participants {
		result = append(result, participant.UserID)
	}

	c.JSON(200, result)
}

func ParticipateGatheringRouter(c *gin.Context) {
	id := c.Param("id")

	idNum, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}

	uid, _ := c.Get("uid")

	participated, err := data.Manager.CheckUserParticipatedGathering(uid.(int), idNum)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if participated {
		c.JSON(400, gin.H{"error": "User already participated this gathering"})
		return
	}

	gathering, err := data.Manager.ReadGathering(idNum)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if gathering.UserID == uid.(int) {
		c.JSON(400, gin.H{"error": "User cannot participate own gathering"})
		return
	}

	participant := data.Participant{
		UserID:      uid.(int),
		GatheringID: idNum,
	}

	_, err = data.Manager.CreateParticipant(&participant)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"id": id})
}

func CancelParticipateGatheringRouter(c *gin.Context) {
	id := c.Param("id")

	idNum, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}

	uid, _ := c.Get("uid")

	participated, err := data.Manager.CheckUserParticipatedGathering(uid.(int), idNum)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if !participated {
		c.JSON(400, gin.H{"error": "User has not participated this gathering"})
		return
	}

	participant := data.Participant{
		UserID:      uid.(int),
		GatheringID: idNum,
	}

	err = data.Manager.DeleteParticipant(participant.ID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"id": id})
}

func GetUpcomingGatheringsRouter(c *gin.Context) {
	uid, _ := c.Get("uid")

	participants, err := data.Manager.ReadParticipantsByUserID(uid.(int))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	upcomingGatherings := make([]*data.Gathering, 0)
	for _, participant := range participants {
		gathering, err := data.Manager.ReadGathering(participant.GatheringID)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		gatheringDateTime, err := time.Parse(config.MysqlDateTimeLayout, gathering.DateTime)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		if gatheringDateTime.After(time.Now()) {
			upcomingGatherings = append(upcomingGatherings, gathering)
		}
	}

	c.JSON(200, upcomingGatherings)
}
