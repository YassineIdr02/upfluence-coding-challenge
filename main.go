package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func analysisHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet || r.URL.Path != "/analysis" {
		http.NotFound(w, r)
		return
	}

	result := AnalysisResult{TotalPosts: 10, MinimumTimestamp: 1717171717171, MaximumTimestamp: 1717171717171, Average: 10}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func main() {
	http.HandleFunc("/analysis", analysisHandler)

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
