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
	depthFlag := flag.Int("depth", 1, "Depth of content downloading (default: 1)")
	exportTypeFlag := flag.String("export", "txt", "Export type: txt, json, csv")

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

	links, err := ProcessfileT(*pathFlag)
	if err != nil {
		log.Fatalf("Error processing file: %v", err)
	}
	if len(links) == 0 {
		log.Fatal("No valid URLs found in the file")
	}
	outputFile := filepath.Join(filepath.Dir(*pathFlag), "results."+*exportTypeFlag)
	log.Printf("Starting URL processing with depth: %d", *depthFlag)
	log.Printf("Export type: %s", *exportTypeFlag)
	ProcessSites(links, sizeUnit, *saveFlag, outputFile, *depthFlag, *exportTypeFlag)
	fmt.Println("Sites processing completed.")
}
