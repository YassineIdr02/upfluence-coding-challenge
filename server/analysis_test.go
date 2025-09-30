package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func fakePosts() []Post {
	return []Post{
		{Timestamp: 100, Likes: 10, Comments: 5, Favorites: 2, Retweets: 1},
		{Timestamp: 200, Likes: 30, Comments: 15, Favorites: 4, Retweets: 2},
	}
}

func TestAnalysisHandler_Comments(t *testing.T) {
	originalFetch := fetchPostsFromSSE
	fetchPostsFromSSE = func(duration time.Duration) ([]Post, error) {
		return fakePosts(), nil
	}
	defer func() { fetchPostsFromSSE = originalFetch }()

	req := httptest.NewRequest(http.MethodGet, RouteAnalysis+"?duration=1s&dimension=comments", nil)
	w := httptest.NewRecorder()

	analysisHandler(w, req)

	res := w.Result()
	defer res.Body.Close()

	var result map[string]interface{}
	_ = json.NewDecoder(res.Body).Decode(&result)

	if result["avg"+Comments].(float64) != 10 {
		t.Errorf("expected avg_comments 10, got %v", result["avg_"+Comments])
	}
}
