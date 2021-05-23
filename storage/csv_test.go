package storage

import (
	"strings"
	"testing"
)

func TestAppendFile(t *testing.T) {
	t.Run("should adds a route in file", func(t *testing.T) {
		var (
			line = "GRU,CDG,100"
		)

		file, err := Create(t.TempDir(), "example", [][]string{})
		if err != nil {
			t.Fatal(err)
		}

		csv := NewCSVStorage(file.Name())
		if err := csv.AppendFile(line); err != nil {
			t.Error(err)
		}
	})

	t.Run("multiples inserts in file", func(t *testing.T) {
		var (
			lines = []string{"GRU,CDG,100", "XPT,YPT,00", "XXX,YYY,999"}
		)

		file, err := Create(t.TempDir(), "example", [][]string{})
		if err != nil {
			t.Fatal(err)
		}

		csv := NewCSVStorage(file.Name())
		for _, line := range lines {
			if err := csv.AppendFile(line); err != nil {
				t.Error(err)
			}
		}

		insertLines, err := csv.Lines()
		if err != nil {
			t.Fatal(err)
		}

		current := strings.Join(insertLines[len(insertLines)-1], ",")
		expected := lines[len(lines)-1]
		if current != expected {
			t.Errorf("current: %s --> expected: %s", current, expected)
		}
	})
}
