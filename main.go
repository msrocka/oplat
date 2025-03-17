package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	targetDef := readTargetDefinition()
	if targetDef == nil {
		return
	}

	urls, err := readRepoUrls("repos.txt")
	if err != nil {
		fmt.Println("Could not read repo URLs:", err)
		return
	}

	for _, url := range urls {
		repo, err := fetchAndParseRepo(url)
		if err != nil {
			fmt.Println("Could not fetch or parse repo:", url, err)
			continue
		}
		fmt.Printf("<repository location=\"%s\" />\n", url)
		syncRepo(repo, targetDef)
	}
}

func readTargetDefinition() *TargetDef {
	var targetDefFile string
	if len(os.Args) > 1 {
		targetDefFile = os.Args[1]
	} else {
		targetDefFile = filepath.Join(os.Getenv("HOME"),
			"Projects/openLCA/repos/olca-app/olca-app/platform.target")
	}
	_, err := os.Stat(targetDefFile)
	if err != nil {
		fmt.Println("Could not read:", targetDefFile, err)
		os.Exit(1)
	}

	targetDef, err := parseTargetDefinition(targetDefFile)
	if err != nil {
		fmt.Println("Could not read target definition from:", targetDefFile, err)
		os.Exit(1)
	}
	return targetDef
}

func syncRepo(repo *Repo, targetDef *TargetDef) {
	for _, unit := range repo.Units {
		id := unit.ID
		for _, loc := range targetDef.Locations {
			for _, targetUnit := range loc.Units {
				if targetUnit.ID == id && targetUnit.Version != unit.Version {
					fmt.Printf("  <unit id=\"%s\" version=\"%s\"/>\n\n", id, unit.Version)
				}
			}
		}
	}
}
