package travel

import (
	"errors"
	"os"
	"testing"

	"github.com/paraizofelipe/bexs/common/csv"
)

func TestLineToRoute(t *testing.T) {
	tests := []struct {
		description string
		in          []string
		expected    Route
	}{
		{
			description: "Simple case",
			in:          []string{"GRU", "CDG", "10"},
			expected:    Route{"GRU", "CDG", 10},
		},
		{
			description: "Other simple case",
			in:          []string{"gru", "cdg", "1"},
			expected:    Route{"gru", "cdg", 1},
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

func TestGetCheapestRoute(t *testing.T) {
	routes := []Route{
		{"GRU", "BRC", 10},
		{"BRC", "SCL", 5},
		{"GRU", "CDG", 75},
		{"GRU", "SCL", 20},
		{"GRU", "ORL", 56},
		{"ORL", "CDG", 5},
		{"SCL", "ORL", 20},
	}

	tests := []struct {
		description string
		in          string
		expected    Route
	}{
		{
			description: "Should return route GRU BRC 10",
			in:          "GRU",
			expected:    Route{"GRU", "BRC", 10},
		},
		{
			description: "Should return route SCL ORL 20",
			in:          "SCL",
			expected:    Route{"SCL", "ORL", 20},
		},
		{
			description: "Should return route BRC SCL 5",
			in:          "BRC",
			expected:    Route{"BRC", "SCL", 5},
		},
		{
			description: "Must not return a route",
			in:          "XPT",
			expected:    Route{},
		},
	}

	travel := NewTravel(routes)

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			cheapestRoute := travel.GetCheapestRoute(test.in)
			if cheapestRoute != test.expected {
				t.Errorf("expected: %v --> current: %v", test.expected, cheapestRoute)
			}
		})
	}
}

func TestTravelString(t *testing.T) {
	tests := []struct {
		description string
		in          []Route
		expected    string
	}{
		{
			description: "Simple case",
			in: []Route{
				{"GRU", "BRC", 10}, {"BRC", "ORL", 10},
			},
			expected: "GRU - BRC - ORL > $20",
		},
		{
			description: "Another case",
			in: []Route{
				{"SCL", "ORL", 10}, {"ORL", "CDG", 10},
			},
			expected: "SCL - ORL - CDG > $20",
		},
	}

	for _, test := range tests {
		travel := NewTravel(test.in)
		travel.BestRoute.Routes = test.in
		travel.BestRoute.TotalPrice = test.in[0].Price + test.in[1].Price

		current := travel.BestRoute.String()

		if current != test.expected {
			t.Errorf("expected: %v --> current: %v", test.expected, current)
		}
	}
}

func TestImportRoutes(t *testing.T) {
	var (
		routes []Route
	)

	tests := []struct {
		description string
		in          [][]string
		expected    []Route
	}{
		{
			description: "Simple case",
			in:          [][]string{{"XPT", "YPT", "10"}, {"YPT", "ZPT", "20"}},
			expected:    []Route{{"XPT", "YPT", 10}, {"YPT", "ZPT", 20}},
		},
		{
			description: "Another case",
			in:          [][]string{{"X", "Y", "10"}, {"Y", "Z", "20"}},
			expected:    []Route{{"X", "Y", 10}, {"Y", "Z", 20}},
		},
	}

	for _, test := range tests {
		f, err := csv.Create(t.TempDir(), "example", test.in)
		if err != nil {
			t.Error(err)
		}
		if routes, err = ImportRoutes(f.Name()); err != nil {
			t.Error(err)
		}

		trip := NewTravel(routes)
		if len(trip.Routes) != len(test.expected) {
			t.Errorf("expected: %v --> current: %v", test.expected, trip.Routes)
		}

		if trip.Routes[0] != test.expected[0] {
			t.Errorf("expected: %v --> current: %v", test.expected[0], trip.Routes[0])
		}

		err = os.Remove(f.Name())
		if err != nil {
			t.Error(err)
		}
	}
}

func TestGetCheapest(t *testing.T) {
	routes := []Route{
		{"GRU", "BRC", 10},
		{"BRC", "SCL", 5},
		{"GRU", "CDG", 75},
		{"GRU", "SCL", 20},
		{"GRU", "ORL", 56},
		{"ORL", "CDG", 5},
		{"SCL", "ORL", 20},
	}

	tests := []struct {
		description string
		in          []string
		expected    int
	}{
		{
			description: "Should return total price of $40",
			in:          []string{"GRU", "CDG"},
			expected:    40,
		},
		{
			description: "Should return totla price of $30",
			in:          []string{"BRC", "CDG"},
			expected:    30,
		},
	}

	travel := NewTravel(routes)

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			bestTravel := travel.GetCheapest(test.in[0], test.in[1])
			if test.expected != bestTravel.TotalPrice {
				t.Errorf("expected: %v --> current: %v", test.expected, bestTravel)
			}
		})
	}
}
