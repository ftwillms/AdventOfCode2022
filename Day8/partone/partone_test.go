package partone

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestReadInput(t *testing.T) {
	theGrid := ReadInput("../test-input.txt")
	SetVisibleLeftToRight(theGrid)
	SetVisibleRightToLeft(theGrid)
	SetVisibleTopToBottom(theGrid)
	SetVisibleBottomToTop(theGrid)
	testCount := CountVisibleTrees()
	require.Equal(t, 21, testCount)
}
