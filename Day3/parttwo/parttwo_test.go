package parttwo

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Parttwo(t *testing.T) {
	cases := []struct {
		testRucksackA   string
		testRucksackB   string
		testRucksackC   string
		expectedOutcome int
	}{
		{
			testRucksackA:   "vJrwpWtwJgWrhcsFMMfFFhFp",
			testRucksackB:   "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
			testRucksackC:   "PmmdzqPrVvPwwTWBwg",
			expectedOutcome: 18,
		},
	}

	for _, tc := range cases {
		t.Run(tc.testRucksackA, func(t *testing.T) {
			newRucksackGroup := NewRucksackGroup(tc.testRucksackA, tc.testRucksackB, tc.testRucksackC)
			require.NotEmpty(t, newRucksackGroup)
			require.Equal(t, tc.expectedOutcome, newRucksackGroup.Score())
		})
	}
}
