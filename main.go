package main

import (
	"fmt"
	"os"

	"gocrossword/pkg/crossword"
	"gocrossword/pkg/loader"
	"gocrossword/pkg/renderer"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run main.go <csv_file> [numbers_csv] <output_filled> <output_empty>")
		os.Exit(1)
	}

	csvFile := os.Args[1]
	var numbersFile string
	var outputFilled, outputEmpty string
	
	if len(os.Args) == 4 {
		outputFilled = os.Args[2]
		outputEmpty = os.Args[3]
	} else if len(os.Args) == 5 {
		numbersFile = os.Args[2]
		outputFilled = os.Args[3]
		outputEmpty = os.Args[4]
	} else {
		fmt.Println("Usage: go run main.go <csv_file> [numbers_csv] <output_filled> <output_empty>")
		os.Exit(1)
	}

	var crossword, err = func() (*crossword.Crossword, error) {
		if numbersFile != "" {
			return loader.LoadFromCSVWithNumbers(csvFile, numbersFile)
		} else {
			return loader.LoadFromCSV(csvFile)
		}
	}()
	
	if err != nil {
		fmt.Printf("Error loading CSV: %v\n", err)
		os.Exit(1)
	}

	pngRenderer := renderer.NewPNGRenderer()

	err = pngRenderer.Render(crossword, outputFilled)
	if err != nil {
		fmt.Printf("Error rendering filled crossword: %v\n", err)
		os.Exit(1)
	}

	err = pngRenderer.RenderEmpty(crossword, outputEmpty)
	if err != nil {
		fmt.Printf("Error rendering empty crossword: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Crossword rendered successfully!\n")
	fmt.Printf("Filled version: %s\n", outputFilled) 
	fmt.Printf("Empty version: %s\n", outputEmpty)
}