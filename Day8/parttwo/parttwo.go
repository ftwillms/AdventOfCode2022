package parttwo

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Tree struct {
	Height int
	left   *Tree
	right  *Tree
	top    *Tree
	bottom *Tree
}

func (t *Tree) getNext(direction string) *Tree {
	var next *Tree
	switch direction {
	case "top":
		next = t.top
	case "bottom":
		next = t.bottom
	case "left":
		next = t.left
	case "right":
		next = t.right
	default:
		log.Fatal("unknown direction")
	}
	return next
}

func (t *Tree) CountDirection(direction string) int {
	next := t.getNext(direction)
	if next != nil {
		count := 0
		for next != nil {
			count++
			if next.Height >= t.Height {
				break
			}
			next = next.getNext(direction)
		}
		return count
	}
	return 0
}

func (t *Tree) Score() int {
	leftCount := t.CountDirection("left")
	rightCount := t.CountDirection("right")
	topCount := t.CountDirection("top")
	bottomCount := t.CountDirection("bottom")
	return leftCount * rightCount * topCount * bottomCount
}

func ReadInput(filepath string) [][]*Tree {
	// Convert the text into a 2d int array
	// Once we have the 2d array
	fmt.Println(fmt.Sprintf("Reading file into string: %s", filepath))
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Failed to open file: ", filepath)
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	theGrid := make([][]*Tree, 0)
	i := 0
	for scanner.Scan() {
		currentLine := scanner.Text()
		gridRow := make([]*Tree, len(currentLine))

		for j, val := range currentLine {
			intVal, _ := strconv.ParseInt(string(val), 10, 64)
			gridRow[j] = &Tree{Height: int(intVal)}
		}
		theGrid = append(theGrid, gridRow)
		i++
	}
	return theGrid
}

func LinkTrees(grid [][]*Tree) {
	// Iterate through the trees linking their friends
	for i, currentRow := range grid {
		for j, currentTree := range currentRow {
			fmt.Printf(" %d ", currentTree.Height)
			if j > 0 {
				currentTree.left = grid[i][j-1]
			}
			if i > 0 {
				currentTree.top = grid[i-1][j]
			}
			if j < len(currentRow)-1 {
				currentTree.right = grid[i][j+1]
			}
			if i < len(grid)-1 {
				currentTree.bottom = grid[i+1][j]
			}
		}
		fmt.Println()
	}
}

func FindBestTree(theGrid [][]*Tree) (*Tree, int) {
	var bestTree *Tree
	bestScore := 0
	for _, row := range theGrid {
		for _, currentTree := range row {
			score := currentTree.Score()
			if bestScore == 0 {
				bestTree = currentTree
				bestScore = score
			}
			if score > bestScore {
				bestTree = currentTree
				bestScore = score
			}
		}
	}
	return bestTree, bestScore
}
