package main

import (
	"fmt"
	"log"
	"net/http"

	"upfluence-coding-challenge/server/router"
)

func main() {
	r := router.NewRouter() 

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
