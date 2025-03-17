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
		fmt.Println(repo.Name)
		break
	}

	for _, loc := range targetDef.Locations {
		fmt.Println(loc.Type)
		fmt.Println(loc.Repo.URL)
		for _, unit := range loc.Units {
			fmt.Println(unit.ID)
			fmt.Println(unit.Version)
		}
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
