package router

import (
	"net/http"

	"upfluence-coding-challenge/server/handlers"
)

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/analysis", handlers.AnalysisHandler)
	return mux
}