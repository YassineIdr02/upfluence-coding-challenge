package main

import (
	"fmt"
	"math"
)

func aggregatePosts(posts []Post, dimension string) AnalysisResult {
	total := len(posts)
	if total == 0 {
		return AnalysisResult{}
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
		case Likes:
			sum += post.Likes
		case Comments:
			sum += post.Comments
		case Favorites:
			sum += post.Favorites
		case Retweets:
			sum += post.Retweets
		}
	}

	avg := float64(sum) / float64(total)
	avg = math.Round(avg)


	return AnalysisResult{
		TotalPosts:       total,
		MinimumTimestamp: minTS,
		MaximumTimestamp: maxTS,
		Average:          map[string]float64{fmt.Sprintf("avg_%s", dimension): avg},
	}
}
