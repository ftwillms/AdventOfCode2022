package main

import (
	"daytwelve/partone"
	"fmt"
	"github.com/beefsack/go-astar"
)

// runPartOne find the shortest path from starting point to end point
func runPartOne() {
	fmt.Println("Part 1")
	worldAsText := partone.ReadInput("input.txt")
	world := partone.ParseWorld(worldAsText)

	fmt.Println("INPUT:")
	fmt.Println(world.RenderPath([]astar.Pather{}))
	from := world.From()
	to := world.To()
	p, dist, found := astar.Path(from, to)
	if !found {
		fmt.Println("Could not find a path")
	} else {
		fmt.Println("Total distance: ", dist)
		fmt.Println("Solution:")
		fmt.Println(world.RenderPath(p))
		fmt.Println("")
		fmt.Println("Length of path: ", len(p)-1)
	}
}

// runPartTwo
// The goal of part two is to find the shortest path amongst all starting points (S || A)
func runPartTwo() {
	fmt.Println("Part 2")
	worldAsText := partone.ReadInput("input.txt")
	world := partone.ParseWorld(worldAsText)

	fmt.Println("INPUT:")
	fmt.Println(world.RenderPath([]astar.Pather{}))
	from := world.From()
	to := world.To()
	// we know this will be found because of part one
	p, _, _ := astar.Path(from, to)
	fewestSteps := len(p) - 1
	for _, row := range world {
		for _, tile := range row {
			if tile.Kind == int('a') {
				p, _, found := astar.Path(tile, to)
				if found {
					stepCount := len(p) - 1
					if stepCount < fewestSteps {
						fewestSteps = stepCount
					}
				}
			}
		}
	}
	fmt.Println("fewest steps: ", fewestSteps)
}

func main() {
	fmt.Println("DAY 12")
	//runPartOne()
	runPartTwo()
}
