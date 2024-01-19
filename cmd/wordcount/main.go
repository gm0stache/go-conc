package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const specialSymbols = ".,;-:\"'"

func countWords(url string, wordCounts map[string]int) {
	resp, err := http.Get(url)
	if err != nil {
		panic("http request caused error")
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		panic("resp signals error")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic("could not read resp bytes")
	}

	bodyContent := string(body)
	bodyContentWords := strings.Split(bodyContent, " ")

	for _, w := range bodyContentWords {
		cleaned := strings.Trim(w, specialSymbols)
		normalized := strings.ToLower(cleaned)

		count, found := wordCounts[normalized]
		if !found {
			wordCounts[normalized] = 1
			continue
		}
		wordCounts[normalized] = count + 1
	}

	fmt.Println("completed:", url)
}

func main() {
	wordCounts := map[string]int{}
	for i := 1000; i < 1030; i++ {
		url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
		go countWords(url, wordCounts)
	}

	time.Sleep(30 * time.Second)

	for word, occurrencesCount := range wordCounts {
		fmt.Printf("%s-%d\n", word, occurrencesCount)
	}
}
