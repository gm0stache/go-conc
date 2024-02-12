package main

import (
	"fmt"
	"slices"
	"sync"
)

func findFactors(rangeStart, rangeEnd, targetNum int) []int {
	res := []int{}
	for i := rangeStart; i < rangeEnd; i++ {
		if i == 0 {
			continue
		}
		if targetNum%i == 0 {
			res = append(res, i)
		}
	}
	return res
}

func main() {
	num := 10
	ch := make(chan []int, num)

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		ch <- findFactors(0, 5, 10)
		wg.Done()
	}()
	go func() {
		ch <- findFactors(5, 10, 10)
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(ch)
	}()

	res := []int{}
	for r := range ch {
		res = append(res, r...)
	}
	slices.Sort(res)

	fmt.Println("factors of", num)
	fmt.Printf("%+v", res)
}
