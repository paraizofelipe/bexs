package travel

type Travel struct {
    Routes []Route
    BestRoute BestRoute
}

type Route struct {
    from string
    to string
    price int
}

type BestRoute struct {
    travel []Route
    totalPrice int
}

func NewTravel() (Travel) {
    return Travel{}
}

func (t *Travel) AddRoute(route Route) {
    t.Routes = append(t.Routes, route)
}

func (t Travel) GetCheapestRoute(origin string) (cheapestRoute Route) {
    for _, route := range t.Routes {
        if route.from == origin {
            if cheapestRoute.price == 0 {
                cheapestRoute = route
            }
            if route.price < cheapestRoute.price {
                cheapestRoute = route
            }
        }
    }
    return
}

func (t Travel) GetCheapest(origin string, destination string) (cheapestTravel BestRoute) {
    for {
        cheapestRoute := t.GetCheapestRoute(origin)
        cheapestTravel.totalPrice = cheapestTravel.totalPrice + cheapestRoute.price
        
        if cheapestRoute.to == destination {
            cheapestTravel.travel = append(cheapestTravel.travel, cheapestRoute)
            break
        }

        cheapestTravel.travel = append(cheapestTravel.travel, cheapestRoute)
        origin = cheapestRoute.to
    }
    return
}
