package parttwo

import (
	"fmt"
	"testing"
)

func Test_TestCountLeft(t *testing.T) {
	// 3 3 5 4 9
	inputs := []int{3, 3, 5, 4, 9}
	var trees []*Tree
	for i, val := range inputs {
		tree := &Tree{Height: val}
		if i > 0 {
			tree.left = trees[i-1]
		}
		trees = append(trees, tree)
	}
	for _, tree := range trees {
		fmt.Printf("Height: %d treeCount: %d\n", tree.Height, tree.CountDirection("left"))
	}
}

func Test_TestCountRight(t *testing.T) {
	// 3 3 5 4 9
	inputs := []int{3, 3, 5, 4, 9}
	var trees []*Tree
	for i := len(inputs) - 1; i >= 0; i-- {
		val := inputs[i]
		tree := &Tree{Height: val}
		// edge
		if i < len(inputs)-1 {
			inverseCount := len(inputs) - i - 1
			tree.right = trees[inverseCount-1]
		}
		trees = append(trees, tree)
	}
	for _, tree := range trees {
		fmt.Printf("Height: %d treeCount: %d\n", tree.Height, tree.CountDirection("right"))
	}
}

func Test_TestCountBottom(t *testing.T) {
	// 3 3 5 4 9
	inputs := []int{7, 5, 3, 3, 9, 5, 3, 7}
	var trees []*Tree
	for i := len(inputs) - 1; i >= 0; i-- {
		val := inputs[i]
		tree := &Tree{Height: val}
		// edge
		if i < len(inputs)-1 {
			inverseCount := len(inputs) - i - 1
			tree.bottom = trees[inverseCount-1]
		}
		trees = append(trees, tree)
	}
	for _, tree := range trees {
		fmt.Printf("Height: %d treeCount: %d\n", tree.Height, tree.CountDirection("bottom"))
	}
}

func Test_TestCountTop(t *testing.T) {
	// 3 3 5 4 9
	inputs := []int{3, 2, 2, 9, 0}
	var trees []*Tree
	for i := 0; i < len(inputs); i++ {
		val := inputs[i]
		tree := &Tree{Height: val}
		// edge
		if i > 0 {
			tree.top = trees[i-1]
		}
		trees = append(trees, tree)
	}
	for _, tree := range trees {
		fmt.Printf("Height: %d treeCount: %d\n", tree.Height, tree.CountDirection("top"))
	}
}

func Test_TestReadInput(t *testing.T) {
	theGrid := ReadInput("../test-input.txt")
	LinkTrees(theGrid)
	bestTree, bestScore := FindBestTree(theGrid)
	fmt.Println("Best tree and score: ", bestTree, bestScore)
}
