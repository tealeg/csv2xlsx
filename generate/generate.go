package generate

import (
	"encoding/csv"
	"fmt"
	"github.com/tealeg/xlsx"
	"os"
	"strings"
)

// XLSXFromCSVs takes a set of input CSVs and generates a XLSX file with
// each sheet named after the input CSV files
func XLSXFromCSVs(inputFiles []string, XLSXPath string, delimiter string) error {
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