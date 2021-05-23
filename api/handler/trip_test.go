package handler

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/paraizofelipe/bexs/storage"
)

func TestGetBestRoute(t *testing.T) {
	var logger = log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

	t.Run("return status code 200 whit JSON body", func(t *testing.T) {
		var expected = `{"routes":[{"from":"GRU","to":"CDG","price":10}],"total_price":10}`
		file, err := storage.Create(t.TempDir(), "example", [][]string{{"GRU", "CDG", "10"}})
		if err != nil {
			t.Fatal(err)
		}

		handler := NewTripHandler(file.Name(), logger)
		ts := httptest.NewServer(http.HandlerFunc(handler.getBestRoute()))
		defer ts.Close()

		res, err := http.Get(ts.URL + "/api/routes/?from=GRU&to=CDG")
		if err != nil {
			t.Error(err)
		}

		if res.StatusCode != 200 {
			t.Errorf("current: %d --> expected: %d", res.StatusCode, 200)
		}

		body, err := io.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			t.Fatal(err)
		}
		if string(body) != expected {
			t.Errorf("current: %s --> expected: %s", string(body), expected)
		}
	})
}

func TestPostRoute(t *testing.T) {
	var logger = log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

	t.Run("return status code 201 with line GRU,CDG,100", func(t *testing.T) {
		var (
			lines [][]string
			tests = []string{"GRU", "CDG", "100"}
		)
		file, err := storage.Create(t.TempDir(), "example", lines)
		if err != nil {
			t.Fatal(err)
		}
		csv := storage.NewCSVStorage(file.Name())

		handler := NewTripHandler(file.Name(), logger)
		ts := httptest.NewServer(http.HandlerFunc(handler.TripHandler))
		url := fmt.Sprintf("%s/%s", ts.URL, "api/routes/?from=GRU&to=CDG&price=100")
		defer ts.Close()

		res, err := http.Post(url, "Content-Type: application/json", nil)
		if err != nil {
			t.Fatal(err)
		}
		if res.StatusCode != 201 {
			t.Errorf("current: %d --> expected: %d", res.StatusCode, 201)
		}
		if lines, err = csv.Lines(); err != nil {
			t.Fatal(err)
		}

		for index, test := range tests {
			if lines[0][index] != test {
				t.Errorf("current: %s --> expected: %s", lines[0][index], test)
			}
		}
	})
}
