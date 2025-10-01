package business


import (
	"fmt"
	"math"

	"upfluence-coding-challenge/server/constants"
	"upfluence-coding-challenge/server/models"
)

func AggregatePosts(posts []models.Post, dimension string) models.AnalysisResult {
	total := len(posts)
	if total == 0 {
		return models.AnalysisResult{}
	}

	minTS := posts[0].Timestamp
	maxTS := posts[0].Timestamp
	sum := 0

	for _, post := range posts {
		if post.Timestamp < minTS {
			minTS = post.Timestamp
		}
		if post.Timestamp > maxTS {
			maxTS = post.Timestamp
		}

		switch dimension {
		case constants.Likes:
			sum += post.Likes
		case constants.Comments:
			sum += post.Comments
		case constants.Favorites:
			sum += post.Favorites
		case constants.Retweets:
			sum += post.Retweets
		}
	}

	avg := float64(sum) / float64(total)
	avg = math.Round(avg)

	return models.AnalysisResult{
		TotalPosts:       total,
		MinimumTimestamp: minTS,
		MaximumTimestamp: maxTS,
		Average:          map[string]float64{fmt.Sprintf("avg_%s", dimension): avg},
	}
}
