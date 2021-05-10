package csv

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type ProcessFunc func([]string) error

func Create(dir, filename string, lines [][]string) (f *os.File, err error) {
	if f, err = ioutil.TempFile(dir, filename); err != nil {
		return
	}
	defer f.Close()
	for _, line := range lines {
		if _, err = f.Write([]byte(fmt.Sprintf("%s\n", strings.Join(line, ",")))); err != nil {
			return
		}
	}
	return
}

// ProcessLines ---
func ProcessLines(filename string, comma rune) (lines [][]string, err error) {
	var (
		f      *os.File
		reader *csv.Reader
	)
	if f, err = os.Open(filename); err != nil {
		return
	}
	reader = csv.NewReader(f)
	reader.Comma = comma
	reader.LazyQuotes = true
	if lines, err = reader.ReadAll(); err != nil {
		log.Panicf("[ERROR] reading lines: %s", err)
	}
	return
}
