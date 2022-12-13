package partone

import (
	"fmt"
	"github.com/beefsack/go-astar"
	"testing"
)

//func Test_allowedMove(t *testing.T) {
//	shouldBeTrue := allowedMove(1, 2)
//	require.True(t, shouldBeTrue)
//	shouldBeTrue = allowedMove(2, 1)
//	require.True(t, shouldBeTrue)
//	shouldBeTrue = allowedMove(0, 0)
//	require.True(t, shouldBeTrue)
//
//	shouldBeFalse := allowedMove(1, 3)
//	require.False(t, shouldBeFalse)
//	shouldBeFalse = allowedMove(3, 1)
//	require.False(t, shouldBeFalse)
//}

func Test_BestPath(t *testing.T) {
	input := `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`

	world := ParseWorld(input)
	fmt.Println(world.RenderPath([]astar.Pather{}))
	p, dist, found := astar.Path(world.From(), world.To())
	if !found {
		t.Log("Could not find a path")
	} else {
		t.Logf("Resulting path\n%s", world.RenderPath(p))
		fmt.Println("")
		fmt.Println("Distance: ", dist)
	}
	////bestSteps := 31
	//grid, startPoint, endPoint := ParseInput(strings.Split(input, "\n"))
	//fmt.Println(len(grid), startPoint, endPoint)
	//distance := endPoint.Sub(startPoint)
	//fmt.Println(distance)

}
