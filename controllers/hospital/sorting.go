package hospital

import (
	"api-server/data"
	"math"
	"sort"
)

func sortHospitalsByLeastWaitingTime(hospitals []*data.Hospital) []*data.Hospital {
	result := make([]*data.Hospital, len(hospitals))
	copy(result, hospitals)

	sort.Slice(result, func(i, j int) bool {
		iWaiting, err := data.Manager.ReadLatestWaitingNumber(result[i].ID)
		if err != nil {
			return false
		}

		jWaiting, err := data.Manager.ReadLatestWaitingNumber(result[j].ID)
		if err != nil {
			return true
		}

		return iWaiting.Number < jWaiting.Number
	})

	return result
}

func distance(lat1, lng1, lat2, lng2 float64) float64 {
	const earthRadius = 6371 * 1000 // meters

	lat1Rad := lat1 * (3.14159265359 / 180)
	lng1Rad := lng1 * (3.14159265359 / 180)
	lat2Rad := lat2 * (3.14159265359 / 180)
	lng2Rad := lng2 * (3.14159265359 / 180)

	latDiff := lat2Rad - lat1Rad
	lngDiff := lng2Rad - lng1Rad

	a := (latDiff/2)*(latDiff/2) + lat1Rad*(lat2Rad-lat1Rad) + (lngDiff/2)*(lngDiff/2)
	c := 2 * (math.Atan2(math.Sqrt(a), math.Sqrt(1-a)))

	return earthRadius * c
}

func sortHospitalsByDistance(hospitals []*data.Hospital, lat, lng float64) []*data.Hospital {
	result := make([]*data.Hospital, len(hospitals))
	copy(result, hospitals)

	sort.Slice(result, func(i, j int) bool {
		iDistance := distance(lat, lng, result[i].Latitude, result[i].Longitude)
		jDistance := distance(lat, lng, result[j].Latitude, result[j].Longitude)

		return iDistance < jDistance
	})

	return result
}
