package main

import "testing"

func TestChooseMyMove(t *testing.T) {
	testCases := []struct {
		pMove tool
		res   result
		want  tool
	}{
		{toolRock, resultWin, toolPaper},
		{toolPaper, resultWin, toolScissors},
		{toolScissors, resultWin, toolRock},
		{toolRock, resultDraw, toolRock},
		{toolPaper, resultDraw, toolPaper},
		{toolScissors, resultDraw, toolScissors},
		{toolRock, resultLoose, toolScissors},
		{toolPaper, resultLoose, toolRock},
		{toolScissors, resultLoose, toolPaper},
	}

	for _, tc := range testCases {
		got := chooseMyMove(tc.pMove, tc.res)
		if got != tc.want {
			t.Errorf("chooseMyMove(%v, %v) = %v, want %v", tc.pMove, tc.res,
				got, tc.want)
		}
	}
}
