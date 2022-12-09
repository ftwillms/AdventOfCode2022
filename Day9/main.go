package main

import (
	"daynine/other"
	"daynine/partone"
	"daynine/parttwo"
	"fmt"
)

func runPartOne() {
	fmt.Println("PART 1")
	instructions := partone.ReadInput("input.txt")

	rope := partone.NewRope(&partone.Coordinate{0, 0}, 1)
	for _, instruction := range instructions {
		rope.MoveHead(instruction)
	}
	fmt.Println(rope.CountVisited())
}

func runPartTwo() {
	fmt.Println("PART 2")
	instructions := parttwo.ReadInput("input.txt")

	rope := parttwo.NewRope(10)
	for _, instruction := range instructions {
		rope.MoveHead(instruction)
	}
	fmt.Println(rope.CountVisited())
}

func runOther() {
	fmt.Println("OTHER SOLUTIONS")
	other.WhatIsImage()
}

func main() {
	fmt.Println("DAY 9")
	// How many positions does the tail of the rope visit at least once?
	//runPartOne()
	//runPartTwo()
	runOther()
}
