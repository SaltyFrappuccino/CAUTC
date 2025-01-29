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

	links, err := ProcessFile_t(*pathFlag)
	if err != nil {
		log.Fatalf("Error processing file: %v", err)
	}
	if len(links) == 0 {
		log.Fatal("No valid URLs found in the file")
	}

	outputFile := filepath.Join(filepath.Dir(*pathFlag), "results.txt")
	ProcessSites(links, sizeUnit, *saveFlag, outputFile)
	fmt.Println("Sites processing completed.")
}
