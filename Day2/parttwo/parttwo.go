package parttwo

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ScoreComparator in round2 returns the value of YourChoice that we should make
// given an opponent choice and a roundEnd value.
func ScoreComparator(oppChoice int, roundEnd int) int {
	if roundEnd == 3 { // tie so return what the opponent choice was
		return oppChoice
	}
	var baseChoice int
	if roundEnd == 6 {
		baseChoice = (oppChoice + 1) % 3
	} else {
		baseChoice = (oppChoice - 1) % 3
	}
	if baseChoice <= 0 {
		baseChoice += 3
	}
	return baseChoice
}

type RPSChoices struct {
	YourChoice     string
	OpponentChoice string
	RoundEnd       string
}

func (choices *RPSChoices) String() string {
	return fmt.Sprintf(
		"Your choice: %s, Opponent choice: %s, Score: %d",
		choices.YourChoice, choices.OpponentChoice, choices.Score())
}

// Score returns an integer based on:
// Your input maps to a direct value + if your input is better than opponent input.
func (choices *RPSChoices) Score() int {
	var roundEndScore int
	switch choices.RoundEnd {
	case "X":
		roundEndScore = 0
	case "Y":
		roundEndScore = 3
	case "Z":
		roundEndScore = 6
	default:
		panic("RoundEnd is invalid")
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

	inputScore := ScoreComparator(opponentScore, roundEndScore)

	return inputScore + roundEndScore
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
			RoundEnd:       choicesRaw[1],
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
