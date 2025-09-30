package main

import (
	"fmt"
	"log"
	"net/http"

	"upfluence-coding-challenge/server/constants"
	"upfluence-coding-challenge/server/handlers"
)

func main() {
	http.HandleFunc(constants.RouteAnalysis, handlers.AnalysisHandler)

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
