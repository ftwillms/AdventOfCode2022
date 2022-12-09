package main

import (
	"daythree/partone"
	"daythree/parttwo"
	"fmt"
)

func runPartOne() {
	fmt.Println("PART 1")
	rucksacks := partone.ReadInputIntoRucksacks("input.txt")
	fmt.Printf("Inputs: %d\n", len(rucksacks))
	totalScore := 0
	for _, rucksack := range rucksacks {
		fmt.Println("Value :", rucksack)

		totalScore += rucksack.Score()
	}
	fmt.Println("Final total: ", totalScore)
}

func runPartTwo() {
	fmt.Println("PART 2")
	groups := parttwo.ReadInputIntoRucksacks("input.txt")
	fmt.Printf("Inputs: %d\n", len(groups))
	totalScore := 0
	for _, group := range groups {
		totalScore += group.Score()
	}
	fmt.Println("Final total: ", totalScore)
}

func main() {
	fmt.Println("DAY 3")
	//runPartOne()
	runPartTwo()
}
