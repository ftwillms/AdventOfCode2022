package partone

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

type Rucksack struct {
	Raw   string
	CompA *Component
	CompB *Component
}

// Score should return the priority of the highest rune
// that exists in both components
func (r *Rucksack) Score() int {
	highestRune := 0
	for k, v := range r.CompA.runes {
		if _, ok := r.CompB.runes[k]; ok {
			if v > highestRune {
				highestRune = v
			}
		}
	}
	if highestRune == 0 {
		log.Fatal("Failed to find a matching rune")
	}
	return highestRune
}

func (r *Rucksack) String() string {
	return fmt.Sprintf("CompA: %s, CompB: %s", r.CompA, r.CompB)
}

func NewRucksack(value string) *Rucksack {
	// Ensure the value has equal characters for the two compartments
	valueLength := len(value)
	if valueLength%2 != 0 {
		log.Fatal("this rucksack value does not have equal parts")
	}
	midpoint := len(value) / 2
	compA := value[0:midpoint]
	compB := value[midpoint:valueLength]
	return &Rucksack{
		Raw:   value,
		CompA: NewComponent(compA),
		CompB: NewComponent(compB),
	}
}

func ReadInputIntoRucksacks(filepath string) []*Rucksack {
	fmt.Println(fmt.Sprintf("Reading file into rucksacks: %s", filepath))
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Failed to open file: ", filepath)
		panic(err)
	}
	defer file.Close()

	var rucksacks []*Rucksack
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rowText := scanner.Text()
		currentRacksack := NewRucksack(rowText)
		rucksacks = append(rucksacks, currentRacksack)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Failed to scan: ", err)
		panic(err)
	}

	return rucksacks
}
