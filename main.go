package main

import (
	"fmt"
	"log"
	"mymodule/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/api/hello", handlers.HelloHandler)

	fmt.Println("Server is running on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
