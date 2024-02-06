package main

import (
	"fmt"
	"myapp/pkg/handlers"
	"net/http"
)

const portNum = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	fmt.Printf("Starting the app on port %s\n", portNum)

	// If it doesn't work, nothing serious
	_ = http.ListenAndServe(portNum, nil)
}
