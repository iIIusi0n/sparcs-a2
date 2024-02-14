package hospital

import (
	"api-server/data"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetHospitalsRouter(c *gin.Context) {
	hospitals, err := data.Manager.ReadHospitals()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to read hospitals"})
		return
	}

	if c.Param("sorting") != "" {
		sorting := c.Query("sorting")
		if sorting == "waiting" {
			hospitals = sortHospitalsByLeastWaitingTime(hospitals)
		} else if sorting == "distance" {
			lat := c.Query("lat")
			lng := c.Query("lng")
			if lat == "" || lng == "" {
				c.JSON(400, gin.H{"error": "Invalid lat or lng"})
				return
			}

			latNum, err := strconv.ParseFloat(lat, 64)
			if err != nil {
				c.JSON(400, gin.H{"error": "Invalid lat"})
				return
			}

			lngNum, err := strconv.ParseFloat(lng, 64)
			if err != nil {
				c.JSON(400, gin.H{"error": "Invalid lng"})
				return
			}

			hospitals = sortHospitalsByDistance(hospitals, latNum, lngNum)
		}

		if sorting != "waiting" && sorting != "distance" {
			c.JSON(400, gin.H{"error": "Invalid sorting"})
			return
		}
	}

	if c.Query("limit") != "" {
		limit, err := strconv.Atoi(c.Query("limit"))
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid limit"})
			return
		}
		if limit < len(hospitals) {
			hospitals = hospitals[:limit]
		}
	}

	c.JSON(200, hospitals)
}

func GetHospitalRouter(c *gin.Context) {
	hospitalID := c.Param("id")
	hospitalIdNum, err := strconv.Atoi(hospitalID)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}

	hospital, err := data.Manager.ReadHospital(hospitalIdNum)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to read hospital"})
		return
	}

	c.JSON(200, hospital)
}

func GetHospitalWaitingNumberRouter(c *gin.Context) {
	hospitalID := c.Param("id")
	hospitalIdNum, err := strconv.Atoi(hospitalID)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}

	waitingNumber, err := data.Manager.ReadLatestWaitingNumber(hospitalIdNum)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to read waiting number"})
		return
	}

	c.JSON(200, waitingNumber)
}

func CreateHospitalRouter(c *gin.Context) {
	var hospital data.Hospital
	if err := c.ShouldBindJSON(&hospital); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	id, err := data.Manager.CreateHospital(&hospital)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create hospital"})
		return
	}

	c.JSON(200, gin.H{"id": id})
}

func CreateHospitalWaitingNumberRouter(c *gin.Context) {
	hospitalID := c.Param("id")
	hospitalIdNum, err := strconv.Atoi(hospitalID)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}

	var waitingNumber data.WaitingNumber
	if err := c.ShouldBindJSON(&waitingNumber); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	waitingNumber.HospitalID = hospitalIdNum
	err = data.Manager.CreateWaitingNumber(&waitingNumber)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create waiting number"})
		return
	}

	c.JSON(200, gin.H{"id": waitingNumber.HospitalID})

}

func UpdateHospitalRouter(c *gin.Context) {
	hospitalID := c.Param("id")
	hospitalIdNum, err := strconv.Atoi(hospitalID)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}

	var hospital data.Hospital
	if err := c.ShouldBindJSON(&hospital); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	hospital.ID = hospitalIdNum
	err = data.Manager.UpdateHospital(&hospital)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to update hospital"})
		return
	}

	c.JSON(200, gin.H{"id": hospitalIdNum})
}
