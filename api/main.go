package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/paraizofelipe/bexs/api/handler"
)

const (
	FILE = "../storage/input-route.csv"
	HOST = "0.0.0.0"
)

func main() {
	var err error
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

	h := handler.New(FILE, logger)

	http.HandleFunc("/api/routes/", h.Trip.TripHandler)

	url := fmt.Sprintf("%s:%s", HOST, os.Getenv("PORT"))

	log.Printf("Server listening in %s", url)

	if err = http.ListenAndServe(url, nil); err != nil {
		logger.Fatal(err)
	}
}
