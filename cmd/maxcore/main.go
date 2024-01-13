package main

import(
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Number of CPU's:", runtime.NumCPU())	
	fmt.Println("Number of max. procedures:", runtime.GOMAXPROCS(0))
}
