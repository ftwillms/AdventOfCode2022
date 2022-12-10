package main

import (
	"dayten/partone"
	"fmt"
)

func runPartOne() {
	fmt.Println("Part 1")
	instructions := partone.ReadInput("input.txt")
	total := partone.RunInstructions(instructions)
	fmt.Println("total: ", total)
}

func main() {
	fmt.Println("DAY 10")
	runPartOne()
}
