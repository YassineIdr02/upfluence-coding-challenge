package main

import "encoding/json"

type Post struct {
	Timestamp int64 `json:"timestamp"`
	Likes     int   `json:"likes,omitempty"`
	Comments  int   `json:"comments,omitempty"`
	Favorites int   `json:"favorites,omitempty"`
	Retweets  int   `json:"retweets,omitempty"`
}

type AnalysisResult struct {
	TotalPosts       int                `json:"total_posts"`
	MinimumTimestamp int64              `json:"minimum_timestamp"`
	MaximumTimestamp int64              `json:"maximum_timestamp"`
	Average          map[string]float64 `json:"-"`
}

func (ar AnalysisResult) MarshalJSON() ([]byte, error) {
	type Alias AnalysisResult
	alias := Alias(ar)

	data, err := json.Marshal(alias)
	if err != nil {
		return nil, err
	}

	var baseMap map[string]interface{}
	if err := json.Unmarshal(data, &baseMap); err != nil {
		return nil, err
	}

	for k, v := range ar.Average {
		baseMap[k] = v
	}

	return json.Marshal(baseMap)
}
