package partone

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var crateHeaderPattern = regexp.MustCompile(`\s*([0-9])\s*`)
var instructionTextPattern = regexp.MustCompile(`move\s+([0-9]+)\s+from\s+([0-9]+)\s+to\s+([0-9]+)`)

type CrateStack struct {
	letters []rune
}

func (cs *CrateStack) TopLetter() string {
	if len(cs.letters) == 0 {
		return ""
	}
	return string(cs.letters[len(cs.letters)-1])
}

type Crates struct {
	CrateStacks []*CrateStack
}

func NewCrates(input []string) *Crates {
	crateLines := input[0 : len(input)-1]
	crateHeaderText := input[len(input)-1]
	crates := &Crates{}
	if crateHeaderText == "" {
		log.Fatal("no header text found")
	}
	for i, columnValue := range crateHeaderText {
		if columnValue == 32 {
			continue
		}
		// loop through each crate text
		// capture the value at `i`
		crateStack := &CrateStack{}
		crateStack.letters = make([]rune, 0)
		for j := len(crateLines) - 1; j >= 0; j-- {
			crateLine := crateLines[j]
			character := crateLine[i]
			if character == 32 {
				continue
			}
			crateStack.letters = append(crateStack.letters, rune(character))
		}
		crates.CrateStacks = append(crates.CrateStacks, crateStack)

	}
	return crates
}

func (c *Crates) RunInstruction(ins *Instruction) {
	// move Count from source to dest
	sourceStack := c.CrateStacks[ins.SourceIdx-1]
	destStack := c.CrateStacks[ins.DestIdx-1]
	calcIdx := len(sourceStack.letters) - ins.Count
	lettersToMove := sourceStack.letters[calcIdx:]
	sourceStack.letters = sourceStack.letters[:calcIdx]
	// re-create source stack with letters to move no longer there
	fmt.Println("Removing letters from source stack", lettersToMove)
	// re-create dest stack with letters appended

	// Part one
	//for i := len(lettersToMove) - 1; i >= 0; i-- {
	//	destStack.letters = append(destStack.letters, lettersToMove[i])
	//}
	// Part two
	for i := 0; i <= len(lettersToMove)-1; i++ {
		destStack.letters = append(destStack.letters, lettersToMove[i])
	}
	fmt.Println("new dest stack", destStack.letters)
}

type Instruction struct {
	Count     int
	SourceIdx int
	DestIdx   int
}

func NewInstruction(input string) *Instruction {
	matches := instructionTextPattern.FindAllSubmatch([]byte(input), -1)
	if matches == nil {
		log.Fatal("failed to match instruction")
	}
	count, err := strconv.ParseInt(string(matches[0][1]), 10, 32)
	source, err := strconv.ParseInt(string(matches[0][2]), 10, 32)
	dest, err := strconv.ParseInt(string(matches[0][3]), 10, 32)
	if err != nil {
		log.Fatal("Failed to parse matches into instruction")
	}
	return &Instruction{int(count), int(source), int(dest)}
}

func ReadInputIntoObjects(filepath string) (*Crates, []*Instruction) {
	fmt.Println(fmt.Sprintf("Reading file into objects: %s", filepath))
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Failed to open file: ", filepath)
		panic(err)
	}
	defer file.Close()

	var crateTextLines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rowText := scanner.Text()
		if rowText == "" {
			break
		}
		crateTextLines = append(crateTextLines, rowText)
	}
	// process the crateTextLines array into crates!
	crates := NewCrates(crateTextLines)

	var instructions []*Instruction
	for scanner.Scan() {
		rowText := scanner.Text()
		currentInstruction := NewInstruction(rowText)
		instructions = append(instructions, currentInstruction)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Failed to scan: ", err)
		panic(err)
	}

	return crates, instructions
}
