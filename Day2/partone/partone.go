package partone

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ScoreComparator(choice1 int, choice2 int) int {
	// 1 ties 1 (rock ties rock)
	// 1 beats 3 (rock beats scissors)
	// 1 loses 2 (rocks loses to paper)
	// 2 beats 1 (paper beats rock)
	// 2 loses 3 (paper loses to scissors)
	// 2 ties 2
	if choice1 == choice2 {
		return 3
	}
	baseChoice := choice1 - 1
	if baseChoice <= 0 {
		baseChoice = 3
	}
	if baseChoice == choice2 {
		// 1, 2, 3
		// 3 beats 2
		// 2 beats 1,
		// 1 beats 3
		return 6
	}
	return 0
}

type RPSChoices struct {
	YourChoice     string
	OpponentChoice string
}

func (choices *RPSChoices) String() string {
	return fmt.Sprintf(
		"Your choice: %s, Opponent choice: %s, Score: %d",
		choices.YourChoice, choices.OpponentChoice, choices.Score())
}

// Score returns an integer based on:
// Your input maps to a direct value + if your input is better than opponent input.
func (choices *RPSChoices) Score() int {
	var inputScore int
	switch choices.YourChoice {
	case "X":
		inputScore = 1
	case "Y":
		inputScore = 2
	case "Z":
		inputScore = 3
	default:
		panic("YourChoice is invalid")
	}
	var opponentScore int
	switch choices.OpponentChoice {
	case "A":
		opponentScore = 1
	case "B":
		opponentScore = 2
	case "C":
		opponentScore = 3
	default:
		panic("YourChoice is invalid")
	}
	return inputScore + ScoreComparator(inputScore, opponentScore)
}

func ReadInputIntoRPSChoices(filepath string) []*RPSChoices {
	fmt.Println(fmt.Sprintf("Reading file into choices: %s", filepath))
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Failed to open file: ", filepath)
		panic(err)
	}
	defer file.Close()

	var allChoices []*RPSChoices
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rowText := scanner.Text()
		// split row text: 1 2
		choicesRaw := strings.Split(rowText, " ")
		currentChoices := &RPSChoices{
			YourChoice:     choicesRaw[1],
			OpponentChoice: choicesRaw[0],
		}
		allChoices = append(allChoices, currentChoices)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Failed to scan: ", err)
		panic(err)
	}

	return allChoices
}
