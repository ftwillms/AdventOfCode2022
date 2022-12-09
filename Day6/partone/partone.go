package partone

import (
	"bufio"
	"fmt"
	"os"
)

func ReadInput(filepath string) string {
	fmt.Println(fmt.Sprintf("Reading file into string: %s", filepath))
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Failed to open file: ", filepath)
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		return scanner.Text()
	}
	return ""
}

func ParseString(input string) int {
	for i := 0; i <= len(input)-4; i++ {
		currentText := map[string]bool{}
		for j := i; j < i+14; j++ {
			currentChar := string(input[j])
			if _, ok := currentText[currentChar]; ok {
				break
			}
			currentText[currentChar] = true
		}
		if len(currentText) == 14 {
			fmt.Println("Current text: ", currentText)
			return i + 14
		}
	}
	return 0
}
