package main

import (
	"encoding/json"
	"net/http"
	"time"
)

// fetchPostsFromSSE is injectable for testing. By default, it calls ReadSSEPosts.
var fetchPostsFromSSE = ReadSSEPosts

func analysisHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method != http.MethodGet || r.URL.Path != RouteAnalysis {
		http.NotFound(w, r)
		return
	}

	duration, err := time.ParseDuration(r.URL.Query().Get("duration"))
	if err != nil {
		http.Error(w, "Invalid duration", http.StatusBadRequest)
		return
	}

	dimension := r.URL.Query().Get("dimension")
	if dimension != Likes && dimension != Comments && dimension != Favorites && dimension != Retweets {
		http.Error(w, "Invalid dimension", http.StatusBadRequest)
		return
	}

	posts, err := fetchPostsFromSSE(duration)
	if err != nil {
		http.Error(w, "Failed to read posts", http.StatusInternalServerError)
		return
	}

	result := aggregatePosts(posts, dimension)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
