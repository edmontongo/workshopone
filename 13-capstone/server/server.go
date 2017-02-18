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
	tm.problems["square"] = &SquareProblemGenerator{}
	tm.next[""] = "square"

	http.Handle("/task/", http.StripPrefix("/task", tm.Mux()))

	err := http.ListenAndServe(host, nil)
	if err != nil {
		log.Fatal("Could not launch HTTP server:", err.Error())
	}
}
