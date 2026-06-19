package main

import (
	_ "embed"
	"strings"
)

//go:embed repos.txt
var reposFile string

func repoUrls() []string {
	var urls []string
	for line := range strings.SplitSeq(reposFile, "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		urls = append(urls, line)
	}
	return urls
}
