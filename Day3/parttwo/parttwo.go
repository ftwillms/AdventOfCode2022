package parttwo

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// ItemValueMapper
// Lowercase item types a through z have priorities 1 through 26.
// Uppercase item types A through Z have priorities 27 through 52.
func ItemValueMapper(val rune) int {
	// ascii values for upper:
	// A -> 65
	// Z -> 90
	// ascii values for lower:
	// a -> 97
	// z -> 122
	asciiVal := int(val)
	newVal := 0
	// lowercase
	if asciiVal >= 97 {
		newVal = asciiVal - (97 - 1)
	} else if asciiVal >= 65 {
		// 27 through 52
		newVal = asciiVal - (65 - 27)
	}

	return newVal
}

type Component struct {
	raw   string
	runes map[rune]int
}

func NewComponent(val string) *Component {

	runes := make(map[rune]int, len(val))
	for _, char := range val {
		runes[char] = ItemValueMapper(char)
	}
	return &Component{
		raw:   val,
		runes: runes,
	}
}

func (c *Component) String() string {
	return c.raw
}

type RucksackGroup struct {
	compA *Component
	compB *Component
	compC *Component
}

// Score should return the priority of the highest rune
// that exists in both components
func (r *RucksackGroup) Score() int {
	highestRune := 0
	for k, v := range r.compA.runes {
		if _, ok := r.compB.runes[k]; ok {
			if _, ok := r.compC.runes[k]; ok {
				if v > highestRune {
					highestRune = v
				}
			}
		}
	}
	if highestRune == 0 {
		log.Fatal("Failed to find a matching rune")
	}
	return highestRune
}

func (r *RucksackGroup) String() string {
	return fmt.Sprintf("CompA: %s, CompB: %s, CompC: %s", r.compA, r.compB, r.compC)
}

func NewRucksackGroup(a string, b string, c string) *RucksackGroup {
	// Ensure the value has equal characters for the two compartments
	return &RucksackGroup{
		compA: NewComponent(a),
		compB: NewComponent(b),
		compC: NewComponent(c),
	}
}

func ReadInputIntoRucksacks(filepath string) []*RucksackGroup {
	fmt.Println(fmt.Sprintf("Reading file into rucksacks: %s", filepath))
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Failed to open file: ", filepath)
		panic(err)
	}
	defer file.Close()

	rucksacks := make([]string, 0)
	var rucksackGroups []*RucksackGroup
	scanner := bufio.NewScanner(file)
	counter := 0
	for scanner.Scan() {
		rowText := scanner.Text()
		rucksacks = append(rucksacks, rowText)
		if (counter+1)%3 == 0 {
			currentRucksackGroup := NewRucksackGroup(rucksacks[0], rucksacks[1], rucksacks[2])
			rucksackGroups = append(rucksackGroups, currentRucksackGroup)
			rucksacks = make([]string, 0)
		}
		counter++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Failed to scan: ", err)
		panic(err)
	}

	return rucksackGroups
}
