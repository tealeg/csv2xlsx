package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"github.com/jessevdk/go-flags"
	"github.com/tealeg/xlsx"
)

type Options struct {
	InputFiles []string `short:"i" description:"Path to CSV input file(s). Multiple allowed with multiple -i flags"`
	OutputFile string   `short:"o" description:"Path to the XLSX output file"`
	Delimiter  string   `short:"d" default:"," description:"Delimiter used in the CSV file(s)"`
}

// usage prints usage of the command
func usage() {
	fmt.Println("csv2xlsx [OPTION]...")
	fmt.Println("Options:")
	fmt.Println("   -i    Path to CSV input file(s). Multiple allowed with multiple -i flags")
	fmt.Println("   -o    Path to XSLX output file.")
	fmt.Println("   -d    Delimiter in the CSV file(s). Default: ,")
}

// generateXLSXFromCSVs takes a set of input CSVs and generates a XLSX file with
// each sheet named after the input CSV files
func generateXLSXFromCSVs(inputFiles []string, XLSXPath string, delimiter string) error {
	xlsxFile := xlsx.NewFile()

	// Loop through all our input files and create our sheets
	for _, file := range inputFiles {
		csvFile, err := os.Open(file)
		if err != nil {
			return fmt.Errorf("unable to open file (%s): %v", file, err)
		}
		defer csvFile.Close()
		reader := csv.NewReader(csvFile)
		reader.Comma = rune(delimiter[0])
		sheetName := getSheetName(file)
		sheet, err := xlsxFile.AddSheet(sheetName)
		if err != nil {
			return fmt.Errorf("unable to add sheet for file (%s): %v", file, err)
		}
		fields, err := reader.Read()
		for err == nil {
			row := sheet.AddRow()
			for _, field := range fields {
				cell := row.AddCell()
				cell.Value = field
			}
			fields, err = reader.Read()
		}
		if err != nil && err.Error() != "EOF" {
			return err
		}
	}

	return xlsxFile.Save(XLSXPath)
}

// getSheetName removes the path and file extension from the input csv
func getSheetName(fileName string) string {
	// Remove .csv extension
	sheetName := strings.Replace(fileName, ".csv", "", -1)
	// Strip path
	parts := strings.Split(sheetName, "/")
	return parts[len(parts)-1]
}

func main() {
	var opts Options
	_, err := flags.Parse(&opts)
	if err != nil {
		fmt.Printf("flags error: %v", err)
		usage()
		return
	}

	err = generateXLSXFromCSVs(opts.InputFiles, opts.OutputFile, opts.Delimiter)
	if err != nil {
		fmt.Printf("Error generating XLSX file: %v", err.Error())
		return
	}
}
