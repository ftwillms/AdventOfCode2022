package main

import (
	"dayseven/partone"
	"fmt"
)

func runPartOne() {
	fmt.Println("PART 1")
	fileTree := partone.ReadInput("input.txt")
	dirSizes := partone.ParseTree(fileTree, "", 0)
	currentTotalUsed := 70000000 - dirSizes["/"]
	delta := 30000000 - currentTotalUsed
	currentMin := 0
	for _, size := range dirSizes {
		if size >= delta {
			if currentMin == 0 {
				currentMin = size
			} else {
				if size < currentMin {
					currentMin = size
				}
			}
		}
	}
	fmt.Println()
	fmt.Println("Total: ", currentMin)
}

func main() {
	fmt.Println("DAY 7")
	runPartOne()
}
