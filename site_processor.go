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

// DownloadContent retrieves the content of a web page from the specified URL.
//
// Parameters:
//   - url: A string representing the URL of the web page to download.
//
// Returns:
//   - []byte: The content of the web page as a byte slice.
//   - error: An error if the download fails, or nil if successful.
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

// ProcessSite downloads and processes the content of a given URL, measuring its size and download duration.
//
// Parameters:
//   - url: A string representing the URL of the web page to process.
//   - sizeUnit: A SizeUnit enum value specifying the desired unit for content size measurement (Bytes, KB, MB, or Chars).
//
// Returns:
//   - ProcessResult: A struct containing the processed URL, content size (in the specified unit), and download duration.
//     If an error occurs during download, the Size field will be set to -1.
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

// ProcessSites concurrently processes multiple URLs, measuring their content size and download duration.
//
// Parameters:
//   - urls: A slice of strings, each representing a URL to be processed.
//   - sizeUnit: A SizeUnit enum value specifying the desired unit for content size measurement (Bytes, KB, MB, or Chars).
//
// Returns:
//   - []ProcessResult: A slice of ProcessResult structs, each containing the processed URL,
//     content size (in the specified unit), and download duration. The results are in the same
//     order as the input URLs. If an error occurs during download for any URL, its corresponding
//     ProcessResult will have a Size field set to -1.
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

// DisplayResults prints the processing results for multiple URLs to the console.
//
// Parameters:
//   - results: A slice of ProcessResult structs containing the processed data for each URL.
//   - sizeUnit: A SizeUnit enum value specifying the unit used for content size measurement
//     (Bytes, KB, MB, or Chars).
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

// SaveResultsToFile writes the processing results for multiple URLs to a file.
//
// This function takes the results of processing multiple URLs and saves them to a specified file.
// Each result is written on a separate line, including the URL, size (or error message), and processing duration.
//
// Parameters:
//   - results: A slice of ProcessResult structs containing the processed data for each URL.
//   - sizeUnit: A SizeUnit enum value specifying the unit used for content size measurement
//     (Bytes, KB, MB, or Chars).
//   - filename: A string representing the name of the file where the results will be saved.
//
// Returns:
//   - error: An error if file creation or writing fails, or nil if the operation is successful.
func SaveResultsToFile(results []ProcessResult, sizeUnit SizeUnit, filename string) error {
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer func(file *os.File) {
        err := file.Close()
        if err != nil {}
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
