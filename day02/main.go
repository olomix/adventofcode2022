package main

import (
	"fmt"
	"strings"

	"github.com/olomix/adventofcode2022/internal"
)

type tool uint8

const (
	toolRock     tool = iota
	toolPaper    tool = iota
	toolScissors tool = iota
)

type result string

const (
	resultLoose result = "X"
	resultDraw  result = "Y"
	resultWin   result = "Z"
)

func (t tool) String() string {
	switch t {
	case toolRock:
		return "Rock"
	case toolPaper:
		return "Paper"
	case toolScissors:
		return "Scissors"
	default:
		return "unknown"
	}
}

var defeats = map[tool]tool{
	toolRock:     toolScissors,
	toolScissors: toolPaper,
	toolPaper:    toolRock,
}

var (
	peerMoveRock     = "A"
	peerMovePaper    = "B"
	peerMoveScissors = "C"
	myMoveRock       = "X"
	myMovePaper      = "Y"
	myMoveScissors   = "Z"
)

var toolScore = map[tool]int{
	toolRock:     1,
	toolPaper:    2,
	toolScissors: 3,
}

func decodePeerMove(in string) tool {
	switch in {
	case peerMoveRock:
		return toolRock
	case peerMovePaper:
		return toolPaper
	case peerMoveScissors:
		return toolScissors
	default:
		panic("invalid peer move")
	}
}

func decodeMyMove(in string) tool {
	switch in {
	case myMoveRock:
		return toolRock
	case myMovePaper:
		return toolPaper
	case myMoveScissors:
		return toolScissors
	default:
		panic("invalid my move")
	}
}

func chooseMyMove(pMove tool, res result) tool {
	switch res {
	case resultLoose:
		r, ok := defeats[pMove]
		if !ok {
			panic("[2] unexpected move")
		}
		return r
	case resultDraw:
		return pMove
	case resultWin:
		for k, v := range defeats {
			if v == pMove {
				return k
			}
		}
		panic("unexpected move")
	default:
		panic("unexpected result")
	}
}

func moveScore(pMove, mMove tool) int {
	if pMove == mMove {
		return 3
	}
	isPeerWon := defeats[pMove] == mMove
	if isPeerWon {
		return 0
	}
	return 6
}

func main() {
	totalScore1 := 0
	totalScore2 := 0

	for txt := range internal.ReadLines("day02/input.txt") {
		if txt == "" {
			continue
		}

		txts := strings.Split(txt, " ")
		if len(txts) != 2 {
			panic("invalid input")
		}

		peerMove := decodePeerMove(txts[0])
		myMove1 := decodeMyMove(txts[1])

		roundScore := moveScore(peerMove, myMove1) + toolScore[myMove1]
		totalScore1 += roundScore

		myMove2 := chooseMyMove(peerMove, result(txts[1]))
		roundScore = moveScore(peerMove, myMove2) + toolScore[myMove2]
		totalScore2 += roundScore
	}

	fmt.Printf("totalScore1: %d\n", totalScore1)
	fmt.Printf("totalScore2: %d\n", totalScore2)
}
