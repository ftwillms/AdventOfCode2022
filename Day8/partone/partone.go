package partone

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// visibleTrees represents an x, y map with visibility of each direction as a value
var visibleTrees = map[int]map[int]bool{}

func SetVisibility(x int, y int) {
	if row, ok := visibleTrees[x]; ok {
		row[y] = true
	} else {
		visibleTrees[x] = map[int]bool{}
		visibleTrees[x][y] = true
	}
}

func SetVisibleLeftToRight(grid [][]int) {
	fmt.Println("Trees visible from left-to-right")
	for i := 0; i < len(grid); i++ {
		currentRow := grid[i]
		maxHeight := 0
		for j := 0; j < len(currentRow); j++ {
			if j == 0 {
				maxHeight = currentRow[j]
				fmt.Printf(" %d ", currentRow[j])
				SetVisibility(i, j)
			} else if maxHeight < currentRow[j] {
				fmt.Printf(" %d ", currentRow[j])
				SetVisibility(i, j)
				maxHeight = currentRow[j]
			} else {
				fmt.Printf(" * ")
			}
		}
		fmt.Println()
	}
	fmt.Println(visibleTrees)
}

func SetVisibleRightToLeft(grid [][]int) {
	fmt.Println("Trees visible from right-to-left")
	for i := 0; i < len(grid); i++ {
		currentRow := grid[i]
		maxHeight := 0
		for j := len(currentRow) - 1; j >= 0; j-- {
			if j == len(currentRow)-1 {
				maxHeight = currentRow[j]
				fmt.Printf(" %d ", currentRow[j])
				SetVisibility(i, j)
			} else if maxHeight < currentRow[j] {
				fmt.Printf(" %d ", currentRow[j])
				SetVisibility(i, j)
				maxHeight = currentRow[j]
			} else {
				fmt.Printf(" * ")
			}
		}
		fmt.Println()
	}
	fmt.Println(visibleTrees)
}

func SetVisibleTopToBottom(grid [][]int) {
	fmt.Println("Trees visible from top-to-bottom")
	// Get the very first row to get the length of columns we have
	colLength := len(grid[0])
	for i := 0; i < colLength; i++ {
		maxHeight := 0
		for j := 0; j < len(grid); j++ {
			currentTree := grid[j][i]
			if j == 0 {
				maxHeight = currentTree
				fmt.Printf(" %d ", currentTree)
				SetVisibility(j, i)
			} else if maxHeight < currentTree {
				fmt.Printf(" %d ", currentTree)
				SetVisibility(j, i)
				maxHeight = currentTree
			} else {
				fmt.Printf(" * ")
			}
		}
		fmt.Println()
	}
	fmt.Println(visibleTrees)
}

func SetVisibleBottomToTop(grid [][]int) {
	fmt.Println("Trees visible from bottom-to-top")
	// Get the very first row to get the length of columns we have
	colLength := len(grid[0])
	for i := 0; i < colLength; i++ {
		maxHeight := 0
		for j := len(grid) - 1; j >= 0; j-- {
			currentTree := grid[j][i]
			if j == len(grid)-1 {
				maxHeight = currentTree
				fmt.Printf(" %d ", currentTree)
				SetVisibility(j, i)
			} else if maxHeight < currentTree {
				fmt.Printf(" %d ", currentTree)
				SetVisibility(j, i)
				maxHeight = currentTree
			} else {
				fmt.Printf(" * ")
			}
		}
		fmt.Println()
	}
	fmt.Println(visibleTrees)
}

func CountVisibleTrees() int {
	count := 0
	for _, row := range visibleTrees {
		for _, col := range row {
			if col {
				count++
			}
		}
	}
	return count
}

func ReadInput(filepath string) [][]int {
	fmt.Println(fmt.Sprintf("Reading file into string: %s", filepath))
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Failed to open file: ", filepath)
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	theGrid := make([][]int, 0)
	i := 0
	for scanner.Scan() {
		currentLine := scanner.Text()
		gridRow := make([]int, len(currentLine))
		for j, val := range currentLine {
			intVal, _ := strconv.ParseInt(string(val), 10, 64)
			gridRow[j] = int(intVal)
		}
		theGrid = append(theGrid, gridRow)
		i++
	}
	return theGrid
}
