package repository

import (
	"errors"
	"testing"

	"github.com/paraizofelipe/bexs/storage"
	"github.com/paraizofelipe/bexs/trip/model"
)

func TestFailureLineToRoute(t *testing.T) {
	tests := []struct {
		description string
		in          []string
		expected    error
	}{
		{
			description: "Should throw a error of number conversion error",
			in:          []string{"GRU", "CDG", "XX"},
			expected:    errors.New(`strconv.Atoi: parsing "XX": invalid syntax`),
		},
		{
			description: "should throw a error of value out of range",
			in:          []string{"GRU", "CDG", "100000000000000000000000"},
			expected:    errors.New(`strconv.Atoi: parsing "100000000000000000000000": value out of range`),
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			_, err := LineToRoute(test.in)
			if err != nil {
				if err.Error() != test.expected.Error() {
					t.Errorf("expected: %v --> current: %v", test.expected, err)
				}
			}
		})
	}
}

func TestLineToRoute(t *testing.T) {
	tests := []struct {
		description string
		in          []string
		expected    model.Route
	}{
		{
			description: "Simple case",
			in:          []string{"GRU", "CDG", "10"},
			expected:    model.Route{From: "GRU", To: "CDG", Price: 10},
		},
		{
			description: "Other simple case",
			in:          []string{"gru", "cdg", "1"},
			expected:    model.Route{From: "gru", To: "cdg", Price: 1},
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			current, err := LineToRoute(test.in)
			if err != nil {
				t.Error(err)
			}
			if current != test.expected {
				t.Errorf("expected: %v - current: %v", test.expected, current)
			}
		})
	}
}

func TestFindCheapestRoute(t *testing.T) {
	lines := [][]string{
		{"GRU,BRC,10"},
		{"BRC,SCL,5"},
		{"GRU,CDG,75"},
		{"GRU,SCL,20"},
		{"GRU,ORL,56"},
		{"ORL,CDG,5"},
		{"SCL,ORL,20"},
	}

	tests := []struct {
		description string
		in          string
		expected    model.Route
	}{
		{
			description: "Should return route GRU BRC 10",
			in:          "GRU",
			expected:    model.Route{From: "GRU", To: "BRC", Price: 10},
		},
		{
			description: "Should return route SCL ORL 20",
			in:          "SCL",
			expected:    model.Route{From: "SCL", To: "ORL", Price: 20},
		},
		{
			description: "Should return route BRC SCL 5",
			in:          "BRC",
			expected:    model.Route{From: "BRC", To: "SCL", Price: 5},
		},
		{
			description: "Must not return a route",
			in:          "XPT",
			expected:    model.Route{},
		},
	}

	for _, test := range tests {

		file, err := storage.Create(t.TempDir(), "example", lines)
		if err != nil {
			t.Fatal(err)
		}

		t.Run(test.description, func(t *testing.T) {
			repository := NewRoute(file.Name())
			cheapestRoute, err := repository.FindCheapestRoute(test.in)

			if err != nil {
				t.Fatal(err)
			}
			if cheapestRoute != test.expected {
				t.Errorf("expected: %v --> current: %v", test.expected, cheapestRoute)
			}
		})
	}
}
