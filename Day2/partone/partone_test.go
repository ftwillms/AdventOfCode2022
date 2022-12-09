package partone

import (
	"fmt"
	"testing"
)

func TestRPSChoices_Score(t *testing.T) {
	cases := []struct {
		testChoice      *RPSChoices
		expectedOutcome int
	}{
		{
			testChoice: &RPSChoices{
				YourChoice:     "Y",
				OpponentChoice: "A",
			},
			expectedOutcome: 8,
		},
		{
			testChoice: &RPSChoices{
				YourChoice:     "X",
				OpponentChoice: "B",
			},
			expectedOutcome: 1,
		},
		{
			testChoice: &RPSChoices{
				YourChoice:     "Y",
				OpponentChoice: "B",
			},
			expectedOutcome: 6,
		},
	}
	for i, tc := range cases {
		fmt.Println(fmt.Sprintf("Test %d: %d? %d", i, tc.expectedOutcome, tc.testChoice.Score()))
	}
}
