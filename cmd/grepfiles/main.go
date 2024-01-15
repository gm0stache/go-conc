package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func searchStringOccurence(dirPath string, target string) {
	dir, err := os.Stat(dirPath)
	if err != nil {
		fmt.Printf("could not access file: %s", err.Error())
		return
	}

	if !dir.IsDir() {
		fmt.Printf("path '%s' is not a directory", dirPath)
		return
	}

	dirItems, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Printf("could not retrieve dir items from '%s'", dirPath)
		return
	}

	for _, item := range dirItems {
		fullPath := filepath.Join(dirPath, item.Name())
		if err != nil {
			fmt.Printf("could not convert path '%s' into info struct", fullPath)
			return
		}

		if item.IsDir() {
			go searchStringOccurence(fullPath, target)
			continue
		}

		byts, err := os.ReadFile(fullPath)
		if err != nil {
			fmt.Printf("could not read file '%s'", fullPath)
			return
		}

		fileContent := string(byts)
		containsTarget := strings.Contains(fileContent, target)
		if containsTarget {
			fmt.Println(fullPath)
		}
	}
}

func main() {
	target := os.Args[1]
	if target == "" {
		fmt.Println("target string must be provided")
		return
	}

	dirPath := os.Args[2]
	if dirPath == "" {
		fmt.Println("base directory must be provided")
		return
	}

	fmt.Printf("looking for '%s' in '%s'\n", target, dirPath)
	go searchStringOccurence(dirPath, target)
	time.Sleep(2 * time.Second)
}
