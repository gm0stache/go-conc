package main

import (
	"fmt"
	"time"
)

const (
	ovenTime    = 5
	defaultTime = 5
)

func PrepareTray(num int) string {
	println("preparing empty tray", num)
	time.Sleep(defaultTime * time.Second)
	return fmt.Sprintf("tray number %d", num)
}

func Mixture(tray string) string {
	println("pouring cupcake mixture in ", tray)
	time.Sleep(defaultTime * time.Second)
	return fmt.Sprintf("cupcake in tray %s", tray)
}

func Bake(mixture string) string {
	println("baking", mixture)
	time.Sleep(ovenTime * time.Second)
	return fmt.Sprintf("baked %s", mixture)
}

// main demonstrates how to do some basic data pipelining.
func main() {
	const boxesCount int = 10
	for i := 0; i < boxesCount; i++ {
		result := Bake(Mixture(PrepareTray(i)))
		println("Accepting", result)
	}
}
