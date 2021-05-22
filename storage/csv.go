package storage

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type CSVStorage struct {
	FileName string
}

func NewCSVStorage(fileName string) *CSVStorage {
	return &CSVStorage{
		FileName: fileName,
	}
}

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

func (c *CSVStorage) LoadFile() (file *os.File, err error) {
	if file, err = os.OpenFile(c.FileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, os.ModeAppend); err != nil {
		return
	}
	return
}

func (c *CSVStorage) AppendFile(line string) (err error) {
	var file *os.File

	if file, err = c.LoadFile(); err != nil {
		return
	}

	defer file.Close()

	_, err = file.WriteString(line)
	return
}

func (c *CSVStorage) Lines() (lines [][]string, err error) {
	var (
		file   *os.File
		reader *csv.Reader
	)

	if file, err = c.LoadFile(); err != nil {
		return
	}

	defer file.Close()

	reader = csv.NewReader(file)
	reader.Comma = ','
	if lines, err = reader.ReadAll(); err != nil {
		return
	}
	return
}
