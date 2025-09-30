package router

import (
	"net/http"

	"upfluence-coding-challenge/server/handlers"
)

func RegisterRoutes() {
	http.HandleFunc("/analysis", handlers.AnalysisHandler)
}
