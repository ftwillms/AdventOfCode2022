package main

import (
	"dayfive/partone"
	"fmt"
)

func runPartOne() {
	fmt.Println("PART 1")
	// read in state of the crates
	crates, instructions := partone.ReadInputIntoObjects("input.txt")
	fmt.Printf("Inputs: %d crates, %d instructions \n", len(crates.CrateStacks), len(instructions))
	for _, instruction := range instructions {
		crates.RunInstruction(instruction)
	}
	fmt.Println()
	for _, stack := range crates.CrateStacks {
		fmt.Printf(stack.TopLetter())
	}
	fmt.Println()
}

//func runPartTwo() {
//	fmt.Println("PART 2")
//	sectionRanges := parttwo.ReadInputIntoSectionRangePairs("input.txt")
//	fmt.Printf("Inputs: %d\n", len(sectionRanges))
//	count := 0
//	for _, sectionRange := range sectionRanges {
//		if sectionRange.Overlaps() {
//			count++
//		}
//	}
//	fmt.Println("Overlap count: ", count)
//}

func main() {
	fmt.Println("DAY 5")
	runPartOne()
}
