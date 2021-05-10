package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/paraizofelipe/bexs/travel"
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
		trip      travel.Travel
		input     string
		routes    []travel.Route
		locations []string
	)
	if err = ValidateArgs(os.Args); err != nil {
		log.Fatal(err)
	}

	if routes, err = travel.ImportRoutes(os.Args[1]); err != nil {
		log.Fatal(err)
	}

	trip = travel.NewTravel(routes)

	for {
		fmt.Printf("please enter the route: ")
		fmt.Scanf("%s", &input)
		locations = strings.Split(input, "-")

		trip.BestRoute = trip.GetCheapest(locations[0], locations[1])
		fmt.Printf("best route: %s\n", trip.BestRoute.String())
	}
}
