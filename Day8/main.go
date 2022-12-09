package main

import (
	"dayeight/partone"
	"dayeight/parttwo"
	"fmt"
)

func runPartOne() {
	fmt.Println("PART 1")
	theGrid := partone.ReadInput("input.txt")
	partone.SetVisibleLeftToRight(theGrid)
	partone.SetVisibleRightToLeft(theGrid)
	partone.SetVisibleTopToBottom(theGrid)
	partone.SetVisibleBottomToTop(theGrid)
	finalCount := partone.CountVisibleTrees()
	fmt.Println("Final: ", finalCount)
}

func runPartTwo() {
	fmt.Println("PART 2")
	theGrid := parttwo.ReadInput("input.txt")
	parttwo.LinkTrees(theGrid)
	bestTree, bestScore := parttwo.FindBestTree(theGrid)
	fmt.Printf("Best score: %d with a height of: %d", bestScore, bestTree.Height)
}

func main() {
	fmt.Println("DAY 8")
	//runPartOne()
	runPartTwo()
}
