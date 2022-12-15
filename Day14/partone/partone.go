package partone

import (
	"fmt"
	"image"
	"log"
	"math"
	"strconv"
	"strings"
)

// StartLocation The sand is pouring into the cave from point 500,0.
var StartLocation = image.Point{X: 500, Y: 0}
var EndLocation = image.Point{X: 500, Y: -1}

var KindMap = map[string]rune{
	"rock":  '#',
	"sand":  'o',
	"start": '+',
	"air":   '.',
	"floor": '=',
}

type Tile struct {
	loc   image.Point
	kind  string
	world World
}

type Line struct {
	start     *Tile
	end       *Tile
	direction image.Point
}

func (l *Line) GetTiles() []*Tile {
	var direction image.Point
	var delta int
	if l.start.loc.X == l.end.loc.X {
		if l.start.loc.Y > l.end.loc.Y {
			direction = image.Point{X: 0, Y: -1}
		} else {
			direction = image.Point{X: 0, Y: 1}
		}
		delta = int(math.Abs(float64(l.start.loc.Sub(l.end.loc).Y)))
	} else {
		if l.start.loc.X > l.end.loc.X {
			direction = image.Point{X: -1, Y: 0}
		} else {
			direction = image.Point{X: 1, Y: 0}
		}
		delta = int(math.Abs(float64(l.start.loc.Sub(l.end.loc).X)))
	}
	currentTile := &Tile{loc: l.start.loc, kind: "rock"}
	lineTiles := []*Tile{currentTile}
	for delta > 0 {
		delta--
		newTile := &Tile{loc: currentTile.loc.Add(direction), kind: "rock"}
		lineTiles = append(lineTiles, newTile)
		currentTile = newTile
	}
	return lineTiles
}

// World is a two-dimensional map of Tiles.
type World map[int]map[int]*Tile

// Tile gets the tile at the given coordinates in the world.
func (w World) Tile(loc image.Point) *Tile {
	if w[loc.X] == nil {
		return nil
	}
	return w[loc.X][loc.Y]
}

// SetTile sets a tile at the given coordinates in the world.
func (w World) SetTile(t *Tile) {
	if w[t.loc.X] == nil {
		w[t.loc.X] = map[int]*Tile{}
	}
	w[t.loc.X][t.loc.Y] = t
	t.world = w
}

// MaxBounds returns the max bounds of the world
func (w World) MaxBounds() (minX, minY, maxX, maxY int) {
	width := len(w)
	if width == 0 {
		// no entries
		return
	}
	maxX = StartLocation.X
	minX = StartLocation.X
	maxY = StartLocation.Y
	minY = StartLocation.Y
	for x, row := range w {
		for y := range row {
			if x < minX {
				minX = x
			}
			if x > maxX {
				maxX = x
			}
			if y < minY {
				minY = y
			}
			if y > maxY {
				maxY = y
			}
		}
	}
	return
}

// Render renders the state of the world.
func (w World) Render() string {
	width := len(w)
	if width == 0 {
		return ""
	}
	minX, minY, maxX, maxY := w.MaxBounds()
	floor := maxY + 2
	rows := make([]string, floor-minY+1)
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= floor; y++ {
			point := image.Point{x, y}
			t := w.Tile(point)
			if t == nil {
				if StartLocation.X == x && StartLocation.Y == y {
					t = &Tile{loc: point, kind: "start"}
				} else if y == floor {
					t = &Tile{loc: point, kind: "floor"}
				} else {
					t = &Tile{loc: point, kind: "air"}
				}
			}
			r := ' '
			r = KindMap[t.kind]
			rows[y] += string(r)
		}
	}
	return strings.Join(rows, "\n")
}

// ParseLine Reads a line of input and returns a series of vertexes that should be connected
func ParseLine(input string) []*Line {
	var theLines []*Line
	// 498,4 -> 498,6 -> 496,6
	var currentVertex *image.Point
	var nextVertex *image.Point
	for _, edgesText := range strings.Split(input, " -> ") {
		// 498,4
		verticesText := strings.Split(edgesText, ",")
		if currentVertex == nil {
			x, _ := strconv.Atoi(verticesText[0])
			y, err := strconv.Atoi(verticesText[1])
			if err != nil {
				log.Fatal("can't read vertex value")
			}
			currentVertex = &image.Point{X: x, Y: y}
		} else if nextVertex == nil {
			x, _ := strconv.Atoi(verticesText[0])
			y, _ := strconv.Atoi(verticesText[1])
			nextVertex = &image.Point{X: x, Y: y}
		}
		if currentVertex != nil && nextVertex != nil {
			// we have both vertices so let's create a line
			aLine := &Line{
				start: &Tile{loc: *currentVertex, kind: "rock"},
				end:   &Tile{loc: *nextVertex, kind: "rock"},
			}
			// append the line
			theLines = append(theLines, aLine)
			// and reset current and next!
			currentVertex = nextVertex
			nextVertex = nil
		}
	}
	return theLines
}

func ReadInput(input string) []*Line {
	linesText := strings.Split(input, "\n")
	var rockLines []*Line
	for _, lineText := range linesText {
		edgeCount := strings.Count(lineText, "->")
		currentRockLine := ParseLine(lineText)
		if len(currentRockLine) != edgeCount {
			log.Fatal("bad")
		}
		rockLines = append(rockLines, currentRockLine...)
	}
	return rockLines
}

func Simulate(rockTiles []*Tile) {
	w := World{}
	for _, rockTile := range rockTiles {
		w.SetTile(rockTile)
	}
	fmt.Println("INITIAL STATE")
	fmt.Println(w.Render())
	currentLoc := StartLocation
	previousLoc := StartLocation.Sub(image.Point{0, -1})
	_, _, _, maxY := w.MaxBounds()
	floor := maxY + 2
	// haven't fallen off the abyss yet...
	count := 0

	// part 1
	//for previousLoc.X >= minX && previousLoc.X <= maxX && previousLoc.Y <= maxY && previousLoc.Y >= minY {

	// part 2
	for previousLoc != StartLocation {
		count++
		motion := true
		previousLoc = currentLoc
		for motion {
			for _, offset := range [][]int{
				{0, 1},  // check down
				{-1, 1}, // check down and to the left
				{1, 1},  // check down and to the right
			} {
				tempLoc := currentLoc.Add(image.Point{offset[0], offset[1]})
				if w.Tile(tempLoc) == nil && tempLoc.Y != floor {
					currentLoc = tempLoc
					// keep falling
					break
				} else {
					currentLoc = previousLoc
				}
			}
			if previousLoc != currentLoc {
				// part 1
				//if currentLoc.X < minX || currentLoc.X > maxX || currentLoc.Y < minY || currentLoc.Y > maxY {
				//if currentLoc == EndLocation && count > 10 {
				//	motion = false
				//}
				previousLoc = currentLoc
			} else {
				motion = false
			}
		}
		//  if we get here all possible offsets were checked
		//  and the sand should settle
		w.SetTile(&Tile{loc: previousLoc, kind: "sand"})
		currentLoc = StartLocation
	}
	fmt.Println(w.Render())
	fmt.Println("COUNT: ", count-1)
}
