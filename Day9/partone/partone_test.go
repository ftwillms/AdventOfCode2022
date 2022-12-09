package partone

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewInstruction(t *testing.T) {
	firstText := "R 4"
	firstInstruction := NewInstruction(firstText)
	require.Equal(t, firstInstruction.direction, RIGHT)
	doubleDigitText := "U 234"
	doubleDigitInstruction := NewInstruction(doubleDigitText)
	require.Equal(t, doubleDigitInstruction.direction, UP)
	require.Equal(t, doubleDigitInstruction.steps, 234)
}

func TestReadInput(t *testing.T) {
	instructions := ReadInput("../test-input.txt")
	require.Equal(t, len(instructions), 8)

	rope := NewRope(&Coordinate{0, 0}, 1)
	for _, instruction := range instructions {
		rope.MoveHead(instruction)
	}
	fmt.Println(rope.visitedCoords)
	fmt.Println(rope.CountVisited())
}

func TestReadInputPart2(t *testing.T) {
	instructions := ReadInput("../test-input-part2.txt")
	require.Equal(t, len(instructions), 8)

	rope := NewRope(&Coordinate{0, 0}, 10)
	for _, instruction := range instructions {
		rope.MoveHead(instruction)
	}
	fmt.Println(rope.visitedCoords)
	fmt.Println(rope.CountVisited())
}

func TestDistance(t *testing.T) {
	coordA := &Coordinate{X: 0, Y: 0}
	coordB := &Coordinate{X: 1, Y: 0}
	val := Distance(coordA, coordB)
	require.Equal(t, 1, int(val))

	coordA = &Coordinate{X: 0, Y: 0}
	coordB = &Coordinate{X: 0, Y: 1}
	val = Distance(coordA, coordB)
	require.Equal(t, 1, int(val))

	coordA = &Coordinate{X: 0, Y: 0}
	coordB = &Coordinate{X: 1, Y: 1}
	val = Distance(coordA, coordB)
	require.Equal(t, 1, int(val))

	coordA = &Coordinate{X: 0, Y: 0}
	coordB = &Coordinate{X: 1, Y: 2}
	val = Distance(coordA, coordB)
	require.Equal(t, 2, int(val))

	coordA = &Coordinate{X: 0, Y: 0}
	coordB = &Coordinate{X: 0, Y: 2}
	val = Distance(coordA, coordB)
	require.Equal(t, 2, int(val))

	coordA = &Coordinate{X: 0, Y: 2}
	coordB = &Coordinate{X: 0, Y: 0}
	val = Distance(coordA, coordB)
	require.Equal(t, 2, int(val))

	coordA = &Coordinate{X: 2, Y: 0}
	coordB = &Coordinate{X: 0, Y: 0}
	val = Distance(coordA, coordB)
	require.Equal(t, 2, int(val))

}
