package main

import (
	"fmt"
	"net/http"

	"github.com/mr-tackie/learning-go/pkg/handlers"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on port %s", portNumber)

	err := http.ListenAndServe(portNumber, nil)

	if err != nil {
		fmt.Println("Error starting server", err)
	}
}
