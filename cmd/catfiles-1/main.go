package main

import (
	"fmt"
	"os"
)

func main() {
	argsWithoutProgName := os.Args[1:]
	for _, filepath := range argsWithoutProgName {
		_, err := os.Stat(filepath)
		if err != nil {
			fmt.Printf("could not access file: %s\n", err.Error())
			continue
		}

		fileContent, err := os.ReadFile(filepath)
		if err != nil {
			fmt.Printf("could not read from file: %s", err.Error())
		}
		fmt.Println(string(fileContent))
	}
}
