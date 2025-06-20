package main

import (
	"TasksAPI/internal"
	"log"
)

func main() {
	server := internal.NewServer(":8080")
	log.Println("Server is running on port 8080")
	log.Fatal(server.ListenAndServe())
}
