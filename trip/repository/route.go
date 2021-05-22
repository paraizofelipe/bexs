package repository

import (
	"strconv"

	"github.com/paraizofelipe/bexs/storage"
	"github.com/paraizofelipe/bexs/trip/model"
)

type Route struct {
	storage *storage.CSVStorage
}

func NewRoute(fileName string) RouteRepository {
	return &Route{
		storage: storage.NewCSVStorage(fileName),
	}
}

// LineToRoute ---
func LineToRoute(line []string) (route model.Route, err error) {
	var price int
	if price, err = strconv.Atoi(line[2]); err != nil {
		return
	}
	route.Price = price
	route.From = line[0]
	route.To = line[1]
	return
}

// Add ---
func (r Route) Add(route model.Route) (err error) {
	return r.storage.AppendFile(route.ToLine())
}

func (r Route) FindCheapestRoute(origin string) (cheapestRoute model.Route, err error) {
	var (
		lines [][]string
		route model.Route
	)

	if lines, err = r.storage.Lines(); err != nil {
		return
	}

	for _, line := range lines {
		if route, err = LineToRoute(line); err != nil {
			return
		}
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
