package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ReadInputIntoSlices(filepath string) [][]int {
	fmt.Println(fmt.Sprintf("Reading file into slices: %s", filepath))
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Failed to open file: ", filepath)
		panic(err)
	}
	defer file.Close()

	rowCount := 0
	currentSlice := make([]int, 0)
	var allSlices [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rowText := scanner.Text()
		if rowText == "" {
			allSlices = append(allSlices, currentSlice)
			currentSlice = make([]int, 0)
			rowCount += 1
			continue
		}
		intValue, err := strconv.ParseInt(rowText, 10, 32)
		if err != nil {
			fmt.Println("Failed to parse int: ", err)
		}
		currentSlice = append(currentSlice, int(intValue))
	}
	allSlices = append(allSlices, currentSlice)

	if err := scanner.Err(); err != nil {
		fmt.Println("Failed to scan: ", err)
		panic(err)
	}

	return allSlices
}

func Sum(integers []int) int {
	sum := 0
	for _, val := range integers {
		sum += val
	}
	return sum
}

func FindMaxValue(inputSlices [][]int) int {
	fmt.Printf("Total groupings: %d\n", len(inputSlices))
	maxValue := 0
	for i, vals := range inputSlices {
		fmt.Println(i, vals)
		currentSum := Sum(vals)
		if currentSum > maxValue {
			fmt.Println("We have a new king!", currentSum)
			maxValue = currentSum
		}
	}
	return maxValue
}

func FindMaxValues(inputSlices [][]int) []int {
	maxValues := []int{0, 0, 0}
	for i, vals := range inputSlices {
		fmt.Println(i, vals)
		// max values being found: 69883, 67966, 66259
		currentSum := Sum(vals)
		for j, maxVal := range maxValues {
			if currentSum > maxVal {
				maxValues[j] = currentSum
				currentSum = maxVal
			}
		}
	}
	return maxValues
}

func main() {
	// Day1 is all about finding the max sum amongst several buckets of integers
	inputSlices := ReadInputIntoSlices("Day1/input.txt")
	maxValues := FindMaxValues(inputSlices)
	fmt.Println("The top three values: ", maxValues)
	total := Sum(maxValues)
	fmt.Println("The sum of the top three: ", total)
}
