package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"

	"github.com/tealeg/xlsx"
)

var xlsxPath = flag.String("o", "", "Path to the XLSX output file")
var csvPath = flag.String("f", "", "Path to the CSV input file")
var delimiter = flag.String("d", ";", "Delimiter for felds in the CSV input.")

func usage() {
	fmt.Printf(`%s: -f=<CSV Input File> -o=<XLSX Output File> -d=<Delimiter>

`,
		os.Args[0])
}

func generateXLSXFromCSV(csvPath string, XLSXPath string, delimiter string) error {
	csvFile, err := os.Open(csvPath)
	if err != nil {
		return err
	}
	defer csvFile.Close()
	reader := csv.NewReader(csvFile)
	if len(delimiter) > 0 {
		reader.Comma = rune(delimiter[0])
	} else {
		reader.Comma = rune(';')
	}
	xlsxFile := xlsx.NewFile()
	sheet, err := xlsxFile.AddSheet(csvPath)
	if err != nil {
		return err
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
	if err != nil {
		fmt.Printf(err.Error())
	}
	return xlsxFile.Save(XLSXPath)
}

func main() {
	flag.Parse()
	if len(os.Args) < 3 {
		usage()
		return
	}
	flag.Parse()
	err := generateXLSXFromCSV(*csvPath, *xlsxPath, *delimiter)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
}
