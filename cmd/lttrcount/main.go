package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const lettrs = "abcdefghijklmnopqrstuvwxyz"

func countLttrs(url string, frequency []int) {
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

	for _, b := range body {
		c := strings.ToLower(string(b))
		cIndex := strings.Index(lettrs, c)
		if cIndex >= 0 {
			frequency[cIndex] += 1
		}
	}
	fmt.Println("completed:", url)
}

func main() {
	frequency := make([]int, len(lettrs))
	for i := 1000; i < 1030; i++ {
		url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
		go countLttrs(url, frequency)
	}

	time.Sleep(30 * time.Second)

	for i, c := range lettrs {
		fmt.Printf("%c-%d\n", c, frequency[i])
	}
}
