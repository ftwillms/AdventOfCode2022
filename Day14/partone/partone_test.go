package partone

import (
	"fmt"
	"image"
	"log"
	"testing"
)

func Test_ReadInput(t *testing.T) {
	input := `498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9`
	rockLines := ReadInput(input)
	fmt.Println(rockLines)
	allRockTiles := make([]*Tile, 0)
	for _, rockLine := range rockLines {
		// Enumerate the tiles
		allRockTiles = append(allRockTiles, rockLine.GetTiles()...)
	}

	Simulate(allRockTiles)
}

func Test_GetTiles(t *testing.T) {
	cases := []struct {
		name          string
		tileStart     *Tile
		tileEnd       *Tile
		expectedStart image.Point
	}{
		{
			name:          "x-increasing",
			tileStart:     &Tile{loc: image.Point{1, 0}},
			tileEnd:       &Tile{loc: image.Point{5, 0}},
			expectedStart: image.Point{5, 0},
		},
		{
			name:          "y-increasing",
			tileStart:     &Tile{loc: image.Point{0, 1}},
			tileEnd:       &Tile{loc: image.Point{0, 5}},
			expectedStart: image.Point{0, 5},
		},
		{
			name:          "x-decreasing",
			tileStart:     &Tile{loc: image.Point{5, 0}},
			tileEnd:       &Tile{loc: image.Point{1, 0}},
			expectedStart: image.Point{5, 0},
		},
		{
			name:          "y-decreasing",
			tileStart:     &Tile{loc: image.Point{0, 5}},
			tileEnd:       &Tile{loc: image.Point{0, 1}},
			expectedStart: image.Point{0, 5},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			rockLine := Line{
				start: tc.tileStart,
				end:   tc.tileEnd,
			}
			tiles := rockLine.GetTiles()
			if tiles[0].loc != tc.expectedStart {
				log.Fatal("nope")
			}
		})
	}
}
