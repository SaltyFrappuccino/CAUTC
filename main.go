package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"
	"strings"
)

func main() {
	pathFlag := flag.String("path", "", "Path to the file containing URLs (relative or absolute)")
	sizeFlag := flag.String("size", "bytes", "Content size unit: bytes, kb, mb, chars")
	saveFlag := flag.Bool("save", false, "Save results to a file (true/false)")

	flag.Parse()

	if *pathFlag == "" {
		log.Fatal("The --path flag is required")
	}

	var sizeUnit SizeUnit
	switch strings.ToLower(*sizeFlag) {
	case "kb":
		sizeUnit = KB
	case "mb":
		sizeUnit = MB
	case "chars":
		sizeUnit = Chars
	case "bytes":
		sizeUnit = Bytes
	default:
		log.Fatalf("Invalid value for --size: %s. Use bytes, kb, mb, or chars", *sizeFlag)
	}

	links, err := ProcessFile(*pathFlag)
	if err != nil {
		log.Fatalf("Error processing file: %v", err)
	}
	if len(links) == 0 {
		log.Fatal("No valid URLs found in the file")
	}
	fmt.Println("URL processing completed.")

	results := ProcessSites(links, sizeUnit)
	fmt.Println("Site processing completed.")

	DisplayResults(results, sizeUnit)

	if *saveFlag {
		outputFile := filepath.Join(filepath.Dir(*pathFlag), "results.txt")
		err := SaveResultsToFile(results, sizeUnit, outputFile)
		if err != nil {
			log.Fatalf("Error saving results to file: %v", err)
		}
		fmt.Printf("Results have been saved to file: %s\n", outputFile)
	}
}
