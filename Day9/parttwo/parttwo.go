package parttwo

import (
	"bufio"
	"fmt"
	"image"
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
	knots       []*Coordinate
	currentHead *Coordinate
	currentTail *Coordinate
	currentIdx  int
	// After each step the head takes, the tail will need to determine if it moves
	// based on if the length is reached.
	// Each time the tail moves, record the destination coordinate
	visitedCoords map[int]map[int]bool
}

func NewRope(length int) *Rope {
	knots := make([]*Coordinate, length)
	for i := 0; i < length; i++ {
		knots[i] = &Coordinate{X: 0, Y: 0}
	}
	return &Rope{
		knots: knots,
		visitedCoords: map[int]map[int]bool{
			0: {
				0: true,
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
		r.currentIdx = 0
		r.currentHead = r.knots[r.currentIdx]
		r.currentTail = r.knots[r.currentIdx+1]
		switch instruction.direction {
		case UP:
			r.currentHead.Y++
		case DOWN:
			r.currentHead.Y--
		case LEFT:
			r.currentHead.X--
		case RIGHT:
			r.currentHead.X++
		default:
			log.Fatal("Invalid move direction")
		}
		r.moveTail()
	}
}

func (r *Rope) moveTail() {

	// calculate distance between head and tail coordinate
	// if distance > length then the tail needs to move
	distanceFloat := Distance(r.currentHead, r.currentTail)
	if int(distanceFloat) <= 1 {
		return
	}
	// At this point we've already determined that the head is no longer within the vicinity of the tail.
	if r.currentTail.Y == r.currentHead.Y {
		// we're on the same row
		if r.currentHead.X-r.currentTail.X > 0 {
			r.currentTail.X++
		} else {
			r.currentTail.X--
		}
	} else if r.currentTail.X == r.currentHead.X {
		// we're in the same column
		if r.currentHead.Y-r.currentTail.Y > 0 {
			r.currentTail.Y++
		} else {
			r.currentTail.Y--
		}
	} else {
		// need to figure out the direction
		// up and to the right
		if (r.currentHead.X-r.currentTail.X) > 0 && (r.currentHead.Y-r.currentTail.Y) > 0 {
			r.currentTail.X++
			r.currentTail.Y++
		} else if (r.currentHead.X-r.currentTail.X) < 0 && (r.currentHead.Y-r.currentTail.Y) > 0 {
			// up and to the left
			r.currentTail.X--
			r.currentTail.Y++
		} else if (r.currentHead.X-r.currentTail.X) > 0 && (r.currentHead.Y-r.currentTail.Y) < 0 {
			// down and to the right
			r.currentTail.X++
			r.currentTail.Y--
		} else if (r.currentHead.X-r.currentTail.X) < 0 && (r.currentHead.Y-r.currentTail.Y) < 0 {
			// down and to the left
			r.currentTail.X--
			r.currentTail.Y--
		} else {
			log.Fatal("we've exceeded the length of the rope somehow...")
		}
	}
	if r.currentTail == r.knots[len(r.knots)-1] {
		// only mark if it's the last tail to be moving
		r.markVisited(r.currentTail)
		return
	} else {
		r.currentIdx++
		r.currentHead = r.currentTail
		r.currentTail = r.knots[r.currentIdx+1]
		r.moveTail()
	}
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
	// for exploring the image library...
	pointA := image.Point{
		X: a.X,
		Y: a.Y,
	}
	pointB := image.Point{
		X: b.X,
		Y: b.Y,
	}
	what := pointA.Sub(pointB)
	xDistance := a.X - b.X
	yDistance := a.Y - b.Y
	if what.X == xDistance || what.Y == yDistance {
		fmt.Println("its equal")
	}
	totalDistance := math.Pow(float64(xDistance), 2) + math.Pow(float64(yDistance), 2)
	finalDistance := math.Sqrt(totalDistance)
	return finalDistance
}
