package partone

import (
	"fmt"
	"strings"
	"testing"
)

func TestTheThing(t *testing.T) {
	cratesString := `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3`
	crates := NewCrates(strings.Split(cratesString, "\n"))
	instructions := []*Instruction{
		{Count: 1, SourceIdx: 2, DestIdx: 1},
		{Count: 3, SourceIdx: 1, DestIdx: 3},
		{Count: 2, SourceIdx: 2, DestIdx: 1},
		{Count: 1, SourceIdx: 1, DestIdx: 2},
	}
	for _, ins := range instructions {
		crates.RunInstruction(ins)
	}
	for _, crate := range crates.CrateStacks {
		fmt.Printf(crate.TopLetter())
	}
}
