package service

import (
	"github.com/paraizofelipe/bexs/trip/model"
)

type Trip struct {
	RouteService RouteService
}

func NewTrip(fileName string) TripService {
	return &Trip{
		RouteService: NewRoute(fileName),
	}
}

// GetCheapest ---
func (t Trip) FindCheapest(origin string, destination string) (cheapestTrip model.BestRoute, err error) {
	var cheapestRoute model.Route

	for {
		if cheapestRoute, err = t.RouteService.FindCheapest(origin); err != nil || cheapestRoute.Price == 0 {
			return
		}
		cheapestTrip.TotalPrice = cheapestTrip.TotalPrice + cheapestRoute.Price

		if cheapestRoute.To == destination {
			cheapestTrip.Routes = append(cheapestTrip.Routes, cheapestRoute)
			break
		}

		cheapestTrip.Routes = append(cheapestTrip.Routes, cheapestRoute)
		origin = cheapestRoute.To
	}
	return
}
