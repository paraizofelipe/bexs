package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/paraizofelipe/bexs/trip/model"
	"github.com/paraizofelipe/bexs/trip/service"
)

type Trip struct {
	Logger       *log.Logger
	TripService  service.TripService
	RouteService service.RouteService
}

func NewTripHandler(filePath string, logger *log.Logger) Trip {
	return Trip{
		Logger:       logger,
		RouteService: service.NewRoute(filePath),
		TripService:  service.NewTrip(filePath),
	}
}

// postRoute ---
func (h Trip) postRoute() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			err         error
			route       = model.Route{}
			queryString = r.URL.Query()
		)

		w.Header().Set("Content-Type", "application/json")

		route.From = queryString.Get("from")
		route.To = queryString.Get("to")
		if route.Price, err = strconv.Atoi(queryString.Get("price")); err != nil {
			http.Error(w, ErrorResponse{
				Status: http.StatusInternalServerError,
				Error:  "failed to add route",
			}.Json(), http.StatusInternalServerError)
			h.Logger.Println(err)
			return
		}

		if err = h.RouteService.Add(route); err != nil {
			http.Error(w, ErrorResponse{
				Status: http.StatusInternalServerError,
				Error:  "failed to find route",
			}.Json(), http.StatusInternalServerError)
			h.Logger.Println(err)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(SuccessResponse{
			Status: http.StatusCreated,
			Msg:    "route successfully created",
		})
	}
}

// getBestRoute ---
func (h Trip) getBestRoute() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			err         error
			bestRoute   model.BestRoute
			origin      string = r.URL.Query().Get("from")
			destination string = r.URL.Query().Get("to")
		)

		w.Header().Set("Content-Type", "application/json")

		if bestRoute, err = h.TripService.FindCheapest(origin, destination); err != nil {
			http.Error(w, ErrorResponse{
				Status: http.StatusInternalServerError,
				Error:  "failed to find route",
			}.Json(), http.StatusInternalServerError)
			h.Logger.Println(err)
			return
		}

		if len(bestRoute.Routes) == 0 {
			http.Error(w, ErrorResponse{
				Status: http.StatusNotFound,
				Error:  "route not found",
			}.Json(), http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		if _, err = w.Write([]byte(bestRoute.Json())); err != nil {
			h.Logger.Println("failed to response")
		}
	}
}

func (h Trip) TripHandler(w http.ResponseWriter, r *http.Request) {
	router := NewRouter(h.Logger)

	router.AddRoute(`routes\/?$`, http.MethodGet, h.getBestRoute())
	router.AddRoute(`routes\/?$`, http.MethodPost, h.postRoute())

	router.ServeHTTP(w, r)
}
