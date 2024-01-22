package main

import (
	"fmt"
	"time"
)

func main() {
	arr := [101]int{1}
	for i := 0; i < len(arr); i++ {
		go extendArr(&arr)
	}

	for arr[100] == 0 {
		fmt.Println("waiting for goroutines...")
		time.Sleep(1 * time.Second)
	}
	fmt.Println(arr)
}

func extendArr(arr *[101]int) {
	for i, v := range arr {
		if v != 0 {
			arr[i] = i + 1
			break
		}
	}
}
