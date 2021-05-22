package service

import "github.com/paraizofelipe/bexs/trip/model"

type TripService interface {
	FindCheapest(origin string, destination string) (model.BestRoute, error)
}

type RouteService interface {
	Add(model.Route) error
	FindCheapest(code string) (model.Route, error)
}
