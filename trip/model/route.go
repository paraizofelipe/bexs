package model

import (
	"encoding/json"
	"fmt"
)

type Route struct {
	From  string `json:"from"`
	To    string `json:"to"`
	Price int    `json:"price"`
}

type BestRoute struct {
	Routes     []Route `json:"routes"`
	TotalPrice int     `json:"total_price"`
}

// ToLine ---
func (r Route) ToLine() string {
	return fmt.Sprintf("%s,%s,%d", r.From, r.To, r.Price)
}

func (r Route) Json() string {
	j, err := json.Marshal(r)
	if err != nil {
		return ""
	}
	return string(j)
}

func (b BestRoute) Json() string {
	j, err := json.Marshal(b)
	if err != nil {
		return ""
	}
	return string(j)
}

// String ---
func (b BestRoute) String() string {
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
