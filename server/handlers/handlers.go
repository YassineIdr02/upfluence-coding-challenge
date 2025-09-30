package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"upfluence-coding-challenge/server/constants"
	"upfluence-coding-challenge/server/business"
	"upfluence-coding-challenge/server/helpers"
)

// fetchPostsFromSSE is injectable for testing. By default, it calls ReadSSEPosts.
var FetchPostsFromSSE = business.ReadSSEPosts

func AnalysisHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method != http.MethodGet || r.URL.Path != constants.RouteAnalysis {
		http.NotFound(w, r)
		return
	}

	duration, err := time.ParseDuration(r.URL.Query().Get("duration"))
	if err != nil {
		http.Error(w, "Invalid duration", http.StatusBadRequest)
		return
	}

	dimension := r.URL.Query().Get("dimension")
	if dimension != constants.Likes && dimension != constants.Comments && dimension != constants.Favorites && dimension != constants.Retweets {
		http.Error(w, "Invalid dimension", http.StatusBadRequest)
		return
	}

	posts, err := FetchPostsFromSSE(duration)
	if err != nil {
		http.Error(w, "Failed to read posts", http.StatusInternalServerError)
		return
	}

	result := helpers.AggregatePosts(posts, dimension)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
