package service

import (
	"github.com/paraizofelipe/bexs/trip/model"
	"github.com/paraizofelipe/bexs/trip/repository"
)

type Route struct {
	RouteRepository repository.RouteRepository
}

func NewRoute(fileName string) RouteService {
	return &Route{
		RouteRepository: repository.NewRoute(fileName),
	}
}

func (r Route) Add(route model.Route) (err error) {
	return r.RouteRepository.Add(route)
}

func (r Route) FindCheapest(origin string) (route model.Route, err error) {
	return r.RouteRepository.FindCheapestRoute(origin)
}
