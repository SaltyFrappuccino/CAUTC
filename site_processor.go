package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

type SizeUnit int

const (
	Bytes SizeUnit = iota
	KB
	MB
	Chars
)

type ProcessResult struct {
	URL      string
	Size     int64
	Duration time.Duration
}

func DownloadContent(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	return io.ReadAll(resp.Body)
}

func ProcessSite(url string, sizeUnit SizeUnit) ProcessResult {
	start := time.Now()
	content, err := DownloadContent(url)
	duration := time.Since(start)

	if err != nil {
		return ProcessResult{URL: url, Size: -1, Duration: duration}
	}

	size := int64(len(content))
	switch sizeUnit {
	case KB:
		size /= 1024
	case MB:
		size /= 1024 * 1024
	case Chars:
		// size is already in number of characters
	}

	return ProcessResult{URL: url, Size: size, Duration: duration}
}

func ProcessSites(urls []string, sizeUnit SizeUnit) []ProcessResult {
	results := make([]ProcessResult, len(urls))
	var wg sync.WaitGroup

	for i, url := range urls {
		wg.Add(1)
		go func(i int, url string) {
			defer wg.Done()
			results[i] = ProcessSite(url, sizeUnit)
		}(i, url)
	}

	wg.Wait()
	return results
}

func DisplayResults(results []ProcessResult, sizeUnit SizeUnit) {
	unitStr := "bytes"
	switch sizeUnit {
	case KB:
		unitStr = "KB"
	case MB:
		unitStr = "MB"
	case Chars:
		unitStr = "characters"
	}

	for _, result := range results {
		if result.Size == -1 {
			fmt.Printf("URL: %s - Error downloading content\n", result.URL)
		} else {
			fmt.Printf("URL: %s - Size: %d %s - Time: %v\n", result.URL, result.Size, unitStr, result.Duration)
		}
	}
}

func SaveResultsToFile(results []ProcessResult, sizeUnit SizeUnit, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	unitStr := "bytes"
	switch sizeUnit {
	case KB:
		unitStr = "KB"
	case MB:
		unitStr = "MB"
	case Chars:
		unitStr = "characters"
	}

	for _, result := range results {
		if result.Size == -1 {
			_, err = fmt.Fprintf(file, "%s - Error downloading content - %v\n", result.URL, result.Duration)
		} else {
			_, err = fmt.Fprintf(file, "%s - %d %s - %v\n", result.URL, result.Size, unitStr, result.Duration)
		}
		if err != nil {
			return err
		}
	}

	return nil
}
