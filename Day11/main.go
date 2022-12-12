package main

import (
	"dayeleven/partone"
	"fmt"
)

func runPartOne() {
	monkies := partone.SetupMonkies()
	// 306162 is too high
	// 400 is too low

	// partone answer was 54054
	// 14398800016 too high
	// 11521878962 too low
	// 11524130415 also too low
	fmt.Println(partone.Simulate(monkies, 10000))
}

func main() {
	runPartOne()
}
