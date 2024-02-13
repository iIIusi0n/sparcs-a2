package gathering

import (
	"math"
	"sort"

	"api-server/data"
)

func getLikesFromGathering(gathering *data.Gathering) (int, error) {
	location, err := data.Manager.ReadGatheringLocationByGatheringID(gathering.ID)
	if err != nil {
		return 0, err
	}

	likes, err := data.Manager.CountLikeOnPost(location.PostID)
	if err != nil {
		return 0, err
	}

	return likes, nil
}

func SortGatheringsByLikes(gatherings []*data.Gathering) []*data.Gathering {
	result := make([]*data.Gathering, len(gatherings))
	copy(result, gatherings)

	sort.Slice(result, func(i, j int) bool {
		iLikes, err := getLikesFromGathering(result[i])
		if err != nil {
			return false
		}

		jLikes, err := getLikesFromGathering(result[j])
		if err != nil {
			return false
		}

		return iLikes > jLikes
	})

	return result
}

func distance(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371e3 // metres
	const PI = 3.14159265358979323846

	phi1 := lat1 * PI / 180
	phi2 := lat2 * PI / 180
	deltaPhi := (lat2 - lat1) * PI / 180
	deltaLambda := (lon2 - lon1) * PI / 180

	a := (deltaPhi / 2) * (deltaPhi / 2)
	b := (phi1 / 2) * (phi2 / 2)
	c := (deltaLambda / 2) * (deltaLambda / 2)
	d := a + b + c
	e := 2 * math.Atan2(math.Sqrt(d), math.Sqrt(1-d))
	return R * e
}

func SortGatheringsByDistance(gatherings []*data.Gathering, latitude, longitude float64) []*data.Gathering {
	result := make([]*data.Gathering, len(gatherings))
	copy(result, gatherings)

	sort.Slice(result, func(i, j int) bool {
		iDistance := distance(latitude, longitude, result[i].Latitude, result[i].Longitude)
		jDistance := distance(latitude, longitude, result[j].Latitude, result[j].Longitude)

		return iDistance < jDistance
	})

	return result
}
