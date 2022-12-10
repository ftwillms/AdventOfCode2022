package partone

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_ReadInput(t *testing.T) {
	instructions := ReadInput("test-input.txt")
	require.Equal(t, 146, len(instructions))
	total := RunInstructions(instructions)
	require.Equal(t, 13140, total)
}

func Test_TestRunInstructions(t *testing.T) {
	instructions := []*instruction{{cmd: "noop"}, {cmd: "addx", value: 3}, {cmd: "addx", value: -5}}
	fmt.Println(RunInstructions(instructions))
}
