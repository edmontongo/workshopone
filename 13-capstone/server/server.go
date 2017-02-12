package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	host := os.Getenv("EDMONTONGO_WORKSHOPONE_HOST")
	if host == "" {
		host = ":8080"
	}

	tm := NewTaskMaster()
	http.Handle("/task/", http.StripPrefix("/task/", tm))

	err := http.ListenAndServe(host, nil)
	if err != nil {
		log.Fatal("Could not launch HTTP server:", err.Error())
	}
}

// GET  /task/     - Provides documentation
// GET  /task/new  - Gets a new task ID
// GET  /task/id   - Gets a task description
// POST /task/id   - Posts a task solution

// GET  /status         - HTML showing status
// GET  /status/updates - Websocket providing updates to status
