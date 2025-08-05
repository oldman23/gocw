package loader

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"gocrossword/pkg/crossword"
)

func LoadFromCSV(filename string) (*crossword.Crossword, error) {
	c := crossword.New()
	grid, err := loadCSV(filename)
	if err != nil {
		return nil, err
	}
	c.SetGrid(grid)
	return c, nil
}

func LoadFromCSVWithNumbers(filename, numbersFile string) (*crossword.Crossword, error) {
	c := crossword.New()
	
	grid, err := loadCSV(filename)
	if err != nil {
		return nil, err
	}
	c.SetGrid(grid)
	
	numbers, err := loadCSV(numbersFile)
	if err != nil {
		return nil, err
	}
	c.SetNumbers(numbers)
	
	return c, nil
}

func loadCSV(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open CSV file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var grid [][]string

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to read CSV: %w", err)
		}
		grid = append(grid, record)
	}

	if len(grid) == 0 {
		return nil, fmt.Errorf("empty CSV file")
	}

	return grid, nil
}