package main

import (
	"daytwo/partone"
	"daytwo/parttwo"
	"fmt"
)

func runPartOne() {
	fmt.Println("PART 1")
	allChoices := partone.ReadInputIntoRPSChoices("day2/input.txt")
	fmt.Printf("Inputs: %d\n", len(allChoices))
	totalScore := 0
	for _, choices := range allChoices {
		totalScore += choices.Score()
	}
	fmt.Println("Final total: ", totalScore)
}

func runPartTwo() {
	fmt.Println("PART 2")
	allChoices := parttwo.ReadInputIntoRPSChoices("input.txt")
	fmt.Printf("Inputs: %d\n", len(allChoices))
	totalScore := 0
	for _, choices := range allChoices {
		totalScore += choices.Score()
	}
	fmt.Println("Final total: ", totalScore)
}

func main() {
	fmt.Println("DAY 2")
	runPartTwo()
}
