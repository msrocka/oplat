package main

import (
	"bufio"
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/ulikunitz/xz"
)

// Repo represents the repository metadata
type Repo struct {
	XMLName    xml.Name   `xml:"repository"`
	Name       string     `xml:"name,attr"`
	Properties []Property `xml:"properties>property"`
	Units      []Unit     `xml:"units>unit"`
}

type Property struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

type Unit struct {
	ID      string `xml:"id,attr"`
	Version string `xml:"version,attr"`
}

func readRepoUrls(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var urls []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		urls = append(urls, line)
	}
	return urls, scanner.Err()
}

func fetchAndParseRepo(baseURL string) (*Repo, error) {
	contentURL := strings.TrimSuffix(baseURL, "/") + "/content.xml.xz"

	// Download the content.xml.xz file
	resp, err := http.Get(contentURL)
	if err != nil {
		return nil, fmt.Errorf("failed to download: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	compressed, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err)
	}

	// Create XZ reader
	xzReader, err := xz.NewReader(bytes.NewReader(compressed))
	if err != nil {
		return nil, fmt.Errorf("failed to create XZ reader: %v", err)
	}

	// Read and decompress
	decompressed, err := io.ReadAll(xzReader)
	if err != nil {
		return nil, fmt.Errorf("failed to decompress: %v", err)
	}

	// Parse XML
	var repo Repo
	if err := xml.Unmarshal(decompressed, &repo); err != nil {
		return nil, fmt.Errorf("failed to parse XML: %v", err)
	}

	return &repo, nil
}
