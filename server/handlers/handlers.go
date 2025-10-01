package handlers

import (
	"log"
	"net/http"
	"time"

	"upfluence-coding-challenge/server/business"
	"upfluence-coding-challenge/server/constants"
	"upfluence-coding-challenge/server/helpers"
)

// fetchPostsFromSSE is injectable for testing. By default, it calls ReadSSEPosts.
var FetchPostsFromSSE = business.ReadSSEPosts

// AnalysisHandler handles GET /analysis requests to analyze social media posts from SSE stream.
// It validates query parameters (duration and dimension), fetches posts from SSE, aggregates them,
// and returns the analysis results as JSON.
func AnalysisHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method != http.MethodGet || r.URL.Path != "/analysis" {
		helpers.WriteJSONError(w, http.StatusNotFound, "not_found", "Not Found")
		return
	}

	duration, err := time.ParseDuration(r.URL.Query().Get("duration"))
	if err != nil {
		helpers.WriteJSONError(w, http.StatusBadRequest, "invalid_duration", "Invalid duration")
		return
	}

	dimension := r.URL.Query().Get("dimension")
	if dimension != constants.Likes && dimension != constants.Comments && dimension != constants.Favorites && dimension != constants.Retweets {
		helpers.WriteJSONError(w, http.StatusBadRequest, "invalid_dimension", "Invalid dimension")
		return
	}

	log.Printf("[INFO] /analysis called with duration=%s dimension=%s", duration, dimension)

	posts, err := FetchPostsFromSSE(duration)
	if err != nil {
		helpers.WriteJSONError(w, http.StatusInternalServerError, "failed_to_read_posts", "Failed to read posts")
		return
	}

	result := business.AggregatePosts(posts, dimension)
	log.Printf("[INFO] Aggregated %d posts for dimension=%s", len(posts), dimension)

	helpers.WriteJSON(w, http.StatusOK, result)
}
