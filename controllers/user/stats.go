package user

import (
	"api-server/data"
)

func getUserTotalLikes(userID int) (int, error) {
	posts, err := data.Manager.ReadPostsByUserID(userID)
	if err != nil {
		return 0, err
	}

	totalLikes := 0
	for _, p := range posts {
		likes, err := data.Manager.CountLikeOnPost(p.ID)
		if err != nil {
			return 0, err
		}
		totalLikes += likes
	}

	return totalLikes, nil
}

func getUserTotalGatherings(userID int) (int, error) {
	return data.Manager.CountGatheringOnUser(userID)
}

func getUserTotalPosts(userID int) (int, error) {
	return data.Manager.CountPostsOnUser(userID)
}
