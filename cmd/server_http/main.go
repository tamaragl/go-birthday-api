package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Routes
	r := routes()

	// Config Server
	svr := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	// Server
	fmt.Println("Server listening")
	if err := svr.ListenAndServe(); err != nil {
		log.Println("Error starting the server:", err.Error())
	}
}
