package main

import (
	"fmt"
	"github.com/dkoston/csv2xlsx/generate"
	"github.com/jessevdk/go-flags"
)

type Options struct {
	InputFiles []string `short:"i" description:"Path to CSV input file(s). Multiple allowed with multiple -i flags" required:"true"`
	OutputFile string   `short:"o" description:"Path to the XLSX output file" required:"true"`
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

func main() {
	var opts Options
	_, err := flags.Parse(&opts)
	if err != nil {
		fmt.Printf("flags error: %v", err)
		usage()
		return
	}

	err = generate.XLSXFromCSVs(opts.InputFiles, opts.OutputFile, opts.Delimiter)
	if err != nil {
		fmt.Printf("Error generating XLSX file: %v", err.Error())
		return
	}
}
