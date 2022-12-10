package partone

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type instruction struct {
	cmd   string
	value int
}

func NewInstruction(input string) *instruction {
	cmd := input[:4]
	var value int64
	if cmd == "addx" {
		value, _ = strconv.ParseInt(input[5:], 10, 64)
	}
	return &instruction{
		cmd:   cmd,
		value: int(value),
	}

}

func RunInstructions(instructions []*instruction) int {
	// noop cmd += 1 tick
	// addx cmd += 2 ticks
	instructionsToAdd := make([]*instruction, 0)

	totalNoise := 0
	x := 1
	tickCount := 1
	cycleCount := 1
	// for part 2 there is an off by one issue that distorted the generated text
	fmt.Printf(".")
	for tickCount > 0 {
		tickCount--

		// for every cycle check to see if we have instructions to add
		if len(instructionsToAdd) > 0 {
			addIns := instructionsToAdd[0]
			x += addIns.value
			instructionsToAdd = instructionsToAdd[1:]
		}

		if cycleCount-1 < len(instructions) {
			ins := instructions[cycleCount-1]
			if ins.cmd == "noop" {
				tickCount++
				instructionsToAdd = append(instructionsToAdd, NewInstruction("noop"))
			} else if ins.cmd == "addx" {
				tickCount += 2
				instructionsToAdd = append(instructionsToAdd, NewInstruction("noop"))
				instructionsToAdd = append(instructionsToAdd, ins)
			}
		}

		// PART 1
		//if cycleCount == 20 {
		//	fmt.Println("20th cycle: ", x)
		//	totalNoise += cycleCount * x
		//} else if (cycleCount-20)%40 == 0 {
		//	fmt.Printf("%dth cycle: %d\n", cycleCount, x)
		//	totalNoise += cycleCount * x
		//}

		// PART 2
		pixelCol := cycleCount % 40
		if pixelCol == 0 {
			fmt.Println()
		}
		if x-1 <= (pixelCol-1) && (pixelCol-1) <= x+1 {
			fmt.Printf("#")
		} else {
			fmt.Printf(".")
		}

		cycleCount++
	}
	return totalNoise
}

func ReadInput(filepath string) []*instruction {
	fmt.Println(fmt.Sprintf("Reading file into string: %s", filepath))
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Failed to open file: ", filepath)
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var instructions []*instruction
	for scanner.Scan() {
		currentLine := scanner.Text()
		currentInstruction := NewInstruction(currentLine)
		instructions = append(instructions, currentInstruction)
	}
	return instructions
}
