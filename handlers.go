package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func analysisHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet || r.URL.Path != "/analysis" {
		http.NotFound(w, r)
		return
	}

	duration, err := time.ParseDuration(r.URL.Query().Get("duration"))
	if err != nil {
		http.Error(w, "Invalid duration", http.StatusBadRequest)
		return
	}

	dimension := r.URL.Query().Get("dimension")
	if dimension != "likes" && dimension != "comments" && dimension != "favorites" && dimension != "retweets" {
		http.Error(w, "Invalid dimension", http.StatusBadRequest)
		return
	}

	postsChan := make(chan Post)
	doneChan := make(chan struct{})

	go ReadSSEPosts(postsChan, duration)

	posts := []Post{}
loop:
	for {
		select {
		case post, ok := <-postsChan:
			if !ok {
				postsChan = nil
			} else {
				posts = append(posts, post)
			}
		case <-doneChan:
			break loop
		}
		if postsChan == nil {
			break
		}
	}

	if len(posts) == 0 {
		http.Error(w, "No posts received from stream", http.StatusInternalServerError)
		return
	}

	totalPosts := len(posts)
	minTimestamp := posts[0].Timestamp
	maxTimestamp := posts[0].Timestamp
	var sum int

	for _, post := range posts {
		if post.Timestamp < minTimestamp {
			minTimestamp = post.Timestamp
		}
		if post.Timestamp > maxTimestamp {
			maxTimestamp = post.Timestamp
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

	avg := float64(sum) / float64(totalPosts)

	result := AnalysisResult{
		TotalPosts:       totalPosts,
		MinimumTimestamp: minTimestamp,
		MaximumTimestamp: maxTimestamp,
		Average:          map[string]interface{}{},
	}
	result.Average[fmt.Sprintf("avg_%s", dimension)] = avg

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}