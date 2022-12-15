package main

import (
	"dayfourteen/partone"
	"os"
)

func runPartOne() {
	contents, _ := os.ReadFile("input.txt")
	rockLines := partone.ReadInput(string(contents))
	allRockTiles := make([]*partone.Tile, 0)
	// part two is a minor tweak to part one
	for _, rockLine := range rockLines {
		// Enumerate the tiles
		rockTiles := rockLine.GetTiles()
		allRockTiles = append(allRockTiles, rockTiles...)
	}
	partone.Simulate(allRockTiles)
}

func main() {
	runPartOne()
}
