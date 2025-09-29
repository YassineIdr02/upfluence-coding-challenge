package main

// Generic struct for the post data (different social networks may have different fields)
type Post struct {
	Timestamp int64 `json:"timestamp"`

	// Optional fields (not all posts have them)
	Likes     int `json:"likes,omitempty"`
	Comments  int `json:"comments,omitempty"`
	Favorites int `json:"favorites,omitempty"`
	Retweets  int `json:"retweets,omitempty"`
}

// Struct for the analysis result
type AnalysisResult struct {
	TotalPosts       int   `json:"total_posts"`
	MinimumTimestamp int64 `json:"minimum_timestamp"`
	MaximumTimestamp int64 `json:"maximum_timestamp"`
	Average          int   `json:"avg_dimension"`
}