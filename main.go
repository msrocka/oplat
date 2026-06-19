package main

import (
	"fmt"
	"os"
)

func main() {

	targetDef := readTargetDefinition()
	if targetDef == nil {
		return
	}

	urls := repoUrls()

	for _, url := range urls {
		repo, err := fetchAndParseRepo(url)
		if err != nil {
			fmt.Println("Could not fetch or parse repo:", url, err)
			continue
		}
		fmt.Println("\n-------------------------------------------")
		fmt.Printf("  <repository location=\"%s\" />\n", url)
		syncRepo(repo, targetDef)
		fmt.Println("-------------------------------------------")
	}
}

func readTargetDefinition() *TargetDef {
	targetDefFile := "platform.target"
	if len(os.Args) > 1 {
		targetDefFile = os.Args[1]
	}
	_, err := os.Stat(targetDefFile)
	if err != nil {
		fmt.Println("Error: could not find", targetDefFile)
		fmt.Println("Run this tool in a directory that contains a platform.target file, or pass the path as an argument.")
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
					fmt.Printf("  <unit id=\"%s\" version=\"%s\"/>\n", id, unit.Version)
				}
			}
		}
	}
}
