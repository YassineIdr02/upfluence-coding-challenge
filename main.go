package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/analysis", analysisHandler)

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
