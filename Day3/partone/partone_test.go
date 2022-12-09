package partone

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Partone(t *testing.T) {
	cases := []struct {
		testRucksack    string
		expectedOutcome int
	}{
		{
			testRucksack:    "vJrwpWtwJgWrhcsFMMfFFhFp",
			expectedOutcome: 16,
		},
		{
			testRucksack:    "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
			expectedOutcome: 38,
		},
		{
			testRucksack:    "PmmdzqPrVvPwwTWBwg",
			expectedOutcome: 42,
		},
		{
			testRucksack:    "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
			expectedOutcome: 22,
		},
		{
			testRucksack:    "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
			expectedOutcome: 22,
		},
		{
			testRucksack:    "ttgJtRGJQctTZtZT",
			expectedOutcome: 20,
		},
		{
			testRucksack:    "CrZsJsPPZsGzwwsLwLmpwMDw",
			expectedOutcome: 19,
		},
		{
			testRucksack:    "sNbGtJbMfssNtvcnWFVmnvDd",
			expectedOutcome: 20,
		},
	}

	for _, tc := range cases {
		t.Run(tc.testRucksack, func(t *testing.T) {
			newRucksack := NewRucksack(tc.testRucksack)
			require.NotEmpty(t, newRucksack)
			require.NotEmpty(t, newRucksack.CompA)
			require.NotEmpty(t, newRucksack.CompB)
			require.Equal(t, tc.expectedOutcome, newRucksack.Score())
		})
	}
}
