package main

import (
	"daysix/partone"
	"fmt"
)

func runPartOne() {
	fmt.Println("PART 1")
	input := partone.ReadInput("input.txt")
	indexFound := partone.ParseString(input)
	fmt.Println(indexFound)
}

func main() {
	fmt.Println("DAY 6")
	runPartOne()
}
