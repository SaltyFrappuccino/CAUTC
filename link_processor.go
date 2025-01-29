package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func FindAndReadFile(filePath string) (string, error) {
	if !filepath.IsAbs(filePath) {
		cwd, err := os.Getwd()
		if err != nil {
			return "", fmt.Errorf("error getting current working directory: %v", err)
		}
		filePath = filepath.Join(cwd, filePath)
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("error reading file: %v", err)
	}

	return string(content), nil
}

func ExtractAndNormalizeLinks(content string) []string {
	var links []string
	scanner := bufio.NewScanner(strings.NewReader(content))

	urlRegex := regexp.MustCompile(`https?://[^\s"']+`)

	for scanner.Scan() {
		line := scanner.Text()
		matches := urlRegex.FindAllString(line, -1)
		for _, match := range matches {
			if link := normalizeLink(match); link != "" {
				links = append(links, link)
			}
		}
	}

	return links
}

func normalizeLink(word string) string {
	if !strings.HasPrefix(word, "http://") && !strings.HasPrefix(word, "https://") {
		word = "https://" + word
	}

	u, err := url.Parse(word)
	if err != nil || u.Hostname() == "" {
		return ""
	}

	u.Scheme = "https"

	return u.String()
}

func ProcessfileT(filePath string) ([]string, error) {
	content, err := FindAndReadFile(filePath)
	if err != nil {
		return nil, err
	}

	links := ExtractAndNormalizeLinks(content)
	return links, nil
}
