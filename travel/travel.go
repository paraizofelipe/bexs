package travel

import (
	"fmt"
	"strconv"

	"github.com/paraizofelipe/bexs/common/csv"
)

type Travel struct {
	Routes    []Route
	BestRoute BestTravel
}

type Route struct {
	From  string
	To    string
	Price int
}

type BestTravel struct {
	Routes     []Route
	TotalPrice int
}

func NewTravel(routes []Route) Travel {
	return Travel{
		Routes: routes,
	}
}

// String ---
func (b BestTravel) String() string {
	var output string
	for index, route := range b.Routes {
		if index == (len(b.Routes) - 1) {
			output += fmt.Sprintf("%s - %s > $%d", route.From, route.To, b.TotalPrice)
			break
		}
		output += fmt.Sprintf("%s - ", route.From)
	}
	return output
}

func ImportRoutes(filepath string) (routes []Route, err error) {
	var (
		lines [][]string
		route Route
	)
	if lines, err = csv.ProcessLines(filepath, ','); err != nil {
		return
	}

	for _, line := range lines {
		if route, err = LineToRoute(line); err != nil {
			return
		}
		routes = append(routes, route)
	}
	return
}

// LineToRoute ---
func LineToRoute(line []string) (route Route, err error) {
	var price int
	if price, err = strconv.Atoi(line[2]); err != nil {
		return
	}
	route.Price = price
	route.From = line[0]
	route.To = line[1]
	return
}

// AddRoute ---
func (t *Travel) AddRoute(route Route) {
	t.Routes = append(t.Routes, route)
}

// GetCheapestRoute ---
func (t Travel) GetCheapestRoute(origin string) (cheapestRoute Route) {
	for _, route := range t.Routes {
		if route.From == origin {
			if cheapestRoute.Price == 0 {
				cheapestRoute = route
			}
			if route.Price < cheapestRoute.Price {
				cheapestRoute = route
			}
		}
	}
	return
}

// GetCheapest ---
func (t Travel) GetCheapest(origin string, destination string) (cheapestTravel BestTravel) {
	for {
		cheapestRoute := t.GetCheapestRoute(origin)
		cheapestTravel.TotalPrice = cheapestTravel.TotalPrice + cheapestRoute.Price

		if cheapestRoute.To == destination {
			cheapestTravel.Routes = append(cheapestTravel.Routes, cheapestRoute)
			break
		}

		cheapestTravel.Routes = append(cheapestTravel.Routes, cheapestRoute)
		origin = cheapestRoute.To
	}
	return
}
