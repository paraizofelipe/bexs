package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/paraizofelipe/bexs/trip/model"
	"github.com/paraizofelipe/bexs/trip/service"
)

// ValidateArgs ---
func ValidateArgs(args []string) (err error) {
	if len(os.Args) != 2 {
		err = errors.New("[ERROR] CSV file is required!")
		return
	}

	match, err := regexp.MatchString(``, os.Args[1])
	if err != nil || !match {
		err = errors.New("[ERROR] the first parameter must be a CSV file!")
		return
	}

	return
}

func main() {
	var (
		err       error
		input     string
		locations []string
		trip      model.Trip
	)

	if err = ValidateArgs(os.Args); err != nil {
		log.Fatal(err)
	}

	tripService := service.NewTrip(os.Args[1])

	for {
		fmt.Printf("please enter the route: ")
		fmt.Scanf("%s", &input)
		locations = strings.Split(input, "-")

		if trip.BestRoute, err = tripService.FindCheapest(locations[0], locations[1]); err != nil {
			log.Fatalf("[ERROR] %v", err)
		}
		fmt.Printf("best route: %s\n", trip.BestRoute.String())
	}
}
