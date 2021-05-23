package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/paraizofelipe/bexs/api/handler"
	"github.com/paraizofelipe/bexs/config"
)

func main() {
	var err error
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

	if _, err = os.Stat(config.Storage); err != nil {
		logger.Fatal(err)
	}
	h := handler.New(config.Storage, logger)

	http.HandleFunc("/api/routes/", h.Trip.TripHandler)

	url := fmt.Sprintf("%s:%s", config.Host, os.Getenv("PORT"))

	log.Printf("Server listening in %s", url)

	if err = http.ListenAndServe(url, nil); err != nil {
		logger.Fatal(err)
	}
}
