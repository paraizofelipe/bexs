package service

import (
	"testing"

	"github.com/paraizofelipe/bexs/storage"
)

func TestFindCheapest(t *testing.T) {
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

	file, err := storage.Create(t.TempDir(), "example", lines)
	if err != nil {
		t.Fatal(err)
	}

	for _, test := range tests {
		service := NewTrip(file.Name())
		bestTravel, err := service.FindCheapest(test.in[0], test.in[1])
		if err != nil {
			t.Fatal(err)
		}
		if test.expected != bestTravel.TotalPrice {
			t.Errorf("expected: %v --> current: %v", test.expected, bestTravel)
		}
	}
}
