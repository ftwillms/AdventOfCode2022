package partone

import (
	"fmt"
	"github.com/beefsack/go-astar"
	"log"
	"os"
	"strings"
)

func ReadInput(filepath string) string {
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal("Failed to parse file")
	}
	return string(bytes)
}

// https://github.com/beefsack/go-astar/blob/master/pather_test.go
// The tests used by the go-astar library almost represent
// this world to a T. The following is most of that test code with some
// minor tweaks to reflect the state of the AOC world.

// pather_test.go implements a basic world and tiles that implement Pather for
// the sake of testing.  This functionality forms the back end for
// path_test.go, and serves as an example for how to use A* for a grid.

// Kind* constants refer to tile kinds for input and output.
const (
	// KindFrom (F) is a tile which marks where the path should be calculated
	// from.
	KindFrom = int('S')
	// KindTo (T) is a tile which marks the goal of the path.
	KindTo = int('E')
	// KindPath (●) is a tile to represent where the path is in the output.
	KindPath = 7
)

// KindRunes map tile kinds to output runes.
var KindRunes = map[int]rune{
	KindFrom: 'S',
	KindTo:   'E',
	KindPath: '●',
}

// RuneKinds map input runes to tile kinds.
var RuneKinds = map[rune]int{
	'S': KindFrom,
	'E': KindTo,
}

// KindCosts map tile kinds to movement costs.
var KindCosts = map[int]float64{
	KindFrom: 96.0,
	KindTo:   122.0,
}

// A Tile is a tile in a grid which implements Pather.
type Tile struct {
	// Kind is the kind of tile, potentially affecting movement.
	Kind int
	// X and Y are the coordinates of the tile.
	X, Y int
	// W is a reference to the World that the tile is a part of.
	W World
}

// PathNeighbors returns the neighbors of the tile, excluding blockers and
// tiles off the edge of the board.
func (t *Tile) PathNeighbors() []astar.Pather {
	var neighbors []astar.Pather
	for _, offset := range [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	} {
		if n := t.W.Tile(t.X+offset[0], t.Y+offset[1]); n != nil {
			// In the Advent of Code world each tile's cost is its own value
			// except for the case of the start and end
			nCost, ok := KindCosts[n.Kind]
			if !ok {
				nCost = float64(n.Kind)
			}
			tCost, ok := KindCosts[t.Kind]
			if !ok {
				tCost = float64(t.Kind)
			}
			// The only neighbors you can jump to are ones that are AT MOST
			// n+1 higher than the current tile
			if tCost+1 >= nCost {
				neighbors = append(neighbors, n)
			}
		}
	}
	return neighbors
}

// PathNeighborCost returns the movement cost of the directly neighboring tile.
func (t *Tile) PathNeighborCost(to astar.Pather) float64 {
	toT := to.(*Tile)
	if toT.Kind == KindFrom || toT.Kind == KindTo {
		return KindCosts[toT.Kind]
	}
	return float64(toT.Kind)
}

// PathEstimatedCost uses Manhattan distance to estimate orthogonal distance
// between non-adjacent nodes.
func (t *Tile) PathEstimatedCost(to astar.Pather) float64 {
	toT := to.(*Tile)
	absX := toT.X - t.X
	if absX < 0 {
		absX = -absX
	}
	absY := toT.Y - t.Y
	if absY < 0 {
		absY = -absY
	}
	return float64(absX + absY)
}

// World is a two-dimensional map of Tiles.
type World map[int]map[int]*Tile

// Tile gets the tile at the given coordinates in the world.
func (w World) Tile(x, y int) *Tile {
	if w[x] == nil {
		return nil
	}
	return w[x][y]
}

// SetTile sets a tile at the given coordinates in the world.
func (w World) SetTile(t *Tile, x, y int) {
	if w[x] == nil {
		w[x] = map[int]*Tile{}
	}
	w[x][y] = t
	t.X = x
	t.Y = y
	t.W = w
}

// FirstOfKind gets the first tile on the board of a kind, used to get the from
// and to tiles as there should only be one of each.
func (w World) FirstOfKind(kind int) *Tile {
	for _, row := range w {
		for _, t := range row {
			if t.Kind == kind {
				return t
			}
		}
	}
	return nil
}

// From gets the from tile from the world.
func (w World) From() *Tile {
	return w.FirstOfKind(KindFrom)
}

// To gets the to tile from the world.
func (w World) To() *Tile {
	return w.FirstOfKind(KindTo)
}

// RenderPath renders a path on top of a world.
func (w World) RenderPath(path []astar.Pather) string {
	width := len(w)
	if width == 0 {
		return ""
	}
	height := len(w[0])
	pathLocs := map[string]bool{}
	for _, p := range path {
		pT := p.(*Tile)
		pathLocs[fmt.Sprintf("%d,%d", pT.X, pT.Y)] = true
	}
	rows := make([]string, height)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			t := w.Tile(x, y)
			r := ' '
			if pathLocs[fmt.Sprintf("%d,%d", x, y)] {
				r = KindRunes[KindPath]
			} else if t != nil {
				r = rune(t.Kind)
			}
			rows[y] += string(r)
		}
	}
	return strings.Join(rows, "\n")
}

// ParseWorld parses a textual representation of a world into a world map.
func ParseWorld(input string) World {
	w := World{}
	for y, row := range strings.Split(strings.TrimSpace(input), "\n") {
		for x, raw := range row {
			w.SetTile(&Tile{
				Kind: int(raw),
			}, x, y)
		}
	}
	return w
}
