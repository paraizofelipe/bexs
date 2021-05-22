package repository

import "github.com/paraizofelipe/bexs/trip/model"

type RouteRepository interface {
	Add(model.Route) error
	FindCheapestRoute(origin string) (model.Route, error)
}
