package parttwo

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRPSChoices_Score(t *testing.T) {
	cases := []struct {
		testChoice      *RPSChoices
		expectedOutcome int
	}{
		{
			testChoice: &RPSChoices{
				RoundEnd:       "Y",
				OpponentChoice: "A",
			},
			expectedOutcome: 4,
		},
		{
			testChoice: &RPSChoices{
				RoundEnd:       "X",
				OpponentChoice: "B",
			},
			expectedOutcome: 1,
		},
		{
			testChoice: &RPSChoices{
				RoundEnd:       "Z",
				OpponentChoice: "C",
			},
			expectedOutcome: 7,
		},
		{
			testChoice: &RPSChoices{
				RoundEnd:       "X",
				OpponentChoice: "A",
			},
			expectedOutcome: 3,
		},
		{
			testChoice: &RPSChoices{
				RoundEnd:       "Z",
				OpponentChoice: "A",
			},
			expectedOutcome: 8,
		},
		{
			testChoice: &RPSChoices{
				RoundEnd:       "X",
				OpponentChoice: "C",
			},
			expectedOutcome: 2,
		},
		{
			testChoice: &RPSChoices{
				RoundEnd:       "Z",
				OpponentChoice: "B",
			},
			expectedOutcome: 9,
		},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			fmt.Println(fmt.Sprintf("Test %d: %d? %d", i, tc.expectedOutcome, tc.testChoice.Score()))
			require.Equal(t, tc.expectedOutcome, tc.testChoice.Score())
		})
	}
}
