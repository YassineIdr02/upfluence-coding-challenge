package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"upfluence-coding-challenge/server/handlers"
	"upfluence-coding-challenge/server/models"
	"upfluence-coding-challenge/server/constants"
)

func fakePosts() []models.Post {
	return []models.Post{
		{Timestamp: 100, Likes: 10, Comments: 5, Favorites: 2, Retweets: 1},
		{Timestamp: 200, Likes: 30, Comments: 15, Favorites: 4, Retweets: 2},
	}
}

func TestAnalysisHandler_Comments(t *testing.T) {
	originalFetch := handlers.FetchPostsFromSSE
	handlers.FetchPostsFromSSE = func(duration time.Duration) ([]models.Post, error) {
		return fakePosts(), nil
	}
	defer func() { handlers.FetchPostsFromSSE = originalFetch }()

	req := httptest.NewRequest(http.MethodGet, "/analysis?duration=1s&dimension=comments", nil)
	w := httptest.NewRecorder()

	handlers.AnalysisHandler(w, req)

	res := w.Result()
	defer res.Body.Close()

	var result map[string]interface{}
	_ = json.NewDecoder(res.Body).Decode(&result)

	if result["avg_"+constants.Comments].(float64) != 10 {
		t.Errorf("expected avg_comments 10, got %v", result["avg_"+constants.Comments])
	}
}
