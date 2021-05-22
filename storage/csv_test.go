package storage

import (
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
}
