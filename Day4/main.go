package main

import (
	"dayfour/partone"
	"dayfour/parttwo"
	"fmt"
)

func runPartOne() {
	fmt.Println("PART 1")
	sectionRanges := partone.ReadInputIntoSectionRangePairs("input.txt")
	fmt.Printf("Inputs: %d\n", len(sectionRanges))
	count := 0
	for _, sectionRange := range sectionRanges {
		if sectionRange.Overlaps() {
			count++
		}
	}
	fmt.Println("Overlap count: ", count)
}

func runPartTwo() {
	fmt.Println("PART 2")
	sectionRanges := parttwo.ReadInputIntoSectionRangePairs("input.txt")
	fmt.Printf("Inputs: %d\n", len(sectionRanges))
	count := 0
	for _, sectionRange := range sectionRanges {
		if sectionRange.Overlaps() {
			count++
		}
	}
	fmt.Println("Overlap count: ", count)
}

func main() {
	fmt.Println("DAY 4")
	//runPartOne()
	runPartTwo()
}
