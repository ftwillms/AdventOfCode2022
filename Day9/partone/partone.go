package partone

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

type Direction string

const (
	RIGHT Direction = "R"
	LEFT  Direction = "L"
	DOWN  Direction = "D"
	UP    Direction = "U"
)

type Instruction struct {
	direction Direction // R, U, L, or D
	steps     int
}

func NewInstruction(input string) *Instruction {
	// Pattern: <single character direction><space><everything after is the steps>
	direction := Direction(input[0])
	stepText := input[2:]
	steps, err := strconv.ParseInt(stepText, 10, 64)
	if err != nil {
		log.Fatal("invalid step count for instruction")
	}
	return &Instruction{
		direction: direction,
		steps:     int(steps),
	}
}

type Rope struct {
	// Head is where the start of the rope is and drives the rope
	head *Coordinate
	// tail is derived from the movement of the head
	tail *Coordinate
	// After each step the head takes, the tail will need to determine if it moves
	// based on if the length is reached.
	// Each time the tail moves, record the destination coordinate
	visitedCoords map[int]map[int]bool
	// length represents the length between the head and tail
	// any time the head moves, the tail must ensure that it stays within this length.
	length int
}

func NewRope(origin *Coordinate, length int) *Rope {
	return &Rope{
		head: origin,
		tail: &Coordinate{
			X: origin.X,
			Y: origin.Y,
		},
		length: length,
		visitedCoords: map[int]map[int]bool{
			origin.X: {
				origin.Y: true,
			},
		},
	}
}

func (r *Rope) markVisited(coord *Coordinate) {
	if row, ok := r.visitedCoords[coord.X]; ok {
		row[coord.Y] = true
	} else {
		r.visitedCoords[coord.X] = map[int]bool{
			coord.Y: true,
		}
	}
}

func (r *Rope) MoveHead(instruction *Instruction) {
	// want to move the head N steps
	for i := 0; i < instruction.steps; i++ {
		switch instruction.direction {
		case UP:
			r.head.Y++
		case DOWN:
			r.head.Y--
		case LEFT:
			r.head.X--
		case RIGHT:
			r.head.X++
		default:
			log.Fatal("Invalid move direction")
		}
		// calculate distance between head and tail coordinate
		// if distance > length then the tail needs to move
		distanceFloat := Distance(r.head, r.tail)
		if int(distanceFloat) > r.length {
			r.moveTail()
		}
	}
}

func (r *Rope) moveTail() {
	// At this point we've already determined that the head is no longer within the vicinity of the tail.
	if r.tail.Y == r.head.Y {
		// we're on the same row
		if r.head.X-r.tail.X > 0 {
			r.tail.X++
		} else {
			r.tail.X--
		}
	} else if r.tail.X == r.head.X {
		// we're in the same column
		if r.head.Y-r.tail.Y > 0 {
			r.tail.Y++
		} else {
			r.tail.Y--
		}
	} else {
		// need to figure out the direction
		// up and to the right
		if (r.head.X-r.tail.X) > 0 && (r.head.Y-r.tail.Y) > 0 {
			r.tail.X++
			r.tail.Y++
		} else if (r.head.X-r.tail.X) < 0 && (r.head.Y-r.tail.Y) > 0 {
			// up and to the left
			r.tail.X--
			r.tail.Y++
		} else if (r.head.X-r.tail.X) > 0 && (r.head.Y-r.tail.Y) < 0 {
			// down and to the right
			r.tail.X++
			r.tail.Y--
		} else if (r.head.X-r.tail.X) < 0 && (r.head.Y-r.tail.Y) < 0 {
			// down and to the left
			r.tail.X--
			r.tail.Y--
		} else {
			log.Fatal("we've exceeded the length of the rope somehow...")
		}
	}
	r.markVisited(r.tail)
}

func (r *Rope) CountVisited() int {
	count := 0
	for _, row := range r.visitedCoords {
		for range row {
			count++
		}
	}
	return count
}

type Coordinate struct {
	X int
	Y int
}

func ReadInput(filepath string) []*Instruction {
	fmt.Println(fmt.Sprintf("Reading file into string: %s", filepath))
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Failed to open file: ", filepath)
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var instructions []*Instruction
	for scanner.Scan() {
		currentLine := scanner.Text()
		currentInstruction := NewInstruction(currentLine)
		instructions = append(instructions, currentInstruction)
	}
	return instructions
}

func Distance(a *Coordinate, b *Coordinate) float64 {
	xDistance := a.X - b.X
	yDistance := a.Y - b.Y
	totalDistance := math.Pow(float64(xDistance), 2) + math.Pow(float64(yDistance), 2)
	finalDistance := math.Sqrt(totalDistance)
	return finalDistance
}
