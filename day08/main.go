package main

import (
	"fmt"

	"github.com/olomix/adventofcode2022/internal"
)

var treesField [][]int

func main() {
	for txt := range internal.ReadLines("day08/input.txt") {
		if txt == "" {
			continue
		}

		row := make([]int, 0, len(txt))
		for j := range txt {
			if txt[j] < '0' || txt[j] > '9' {
				continue
			}
			row = append(row, int(txt[j]-'0'))
		}
		treesField = append(treesField, row)
	}

	visibleCount := 0
	maxScore := 0
	for i := range treesField {
		for j := range treesField[i] {
			if isTreeVisible(i, j) {
				visibleCount += 1
			}

			score := scenicScore(i, j)
			if score > maxScore {
				maxScore = score
			}
		}
	}
	fmt.Println(visibleCount)
	fmt.Println(maxScore)
}

func scenicScore(i, j int) int {
	score1 := 0
	for i2 := i - 1; i2 >= 0; i2-- {
		score1++
		if treesField[i2][j] >= treesField[i][j] {
			break
		}
	}

	score2 := 0
	for i2 := i + 1; i2 < len(treesField); i2++ {
		score2++
		if treesField[i2][j] >= treesField[i][j] {
			break
		}
	}

	score3 := 0
	for j2 := j - 1; j2 >= 0; j2-- {
		score3++
		if treesField[i][j2] >= treesField[i][j] {
			break
		}
	}

	score4 := 0
	for j2 := j + 1; j2 < len(treesField[i]); j2++ {
		score4++
		if treesField[i][j2] >= treesField[i][j] {
			break
		}
	}

	return score1 * score2 * score3 * score4
}

func isTreeVisible(i, j int) bool {
	if i == 0 || i == len(treesField)-1 {
		return true
	}

	if j == 0 || j == len(treesField[i])-1 {
		return true
	}

	visible := true
	for i2 := 0; i2 < i; i2++ {
		if treesField[i2][j] >= treesField[i][j] {
			visible = false
			break
		}
	}

	if visible {
		return true
	}

	visible = true
	for i2 := i + 1; i2 < len(treesField); i2++ {
		if treesField[i2][j] >= treesField[i][j] {
			visible = false
			break
		}
	}
	if visible {
		return true
	}

	visible = true
	for j2 := 0; j2 < j; j2++ {
		if treesField[i][j2] >= treesField[i][j] {
			visible = false
			break
		}
	}
	if visible {
		return true
	}

	visible = true
	for j2 := j + 1; j2 < len(treesField[i]); j2++ {
		if treesField[i][j2] >= treesField[i][j] {
			visible = false
			break
		}
	}

	return visible
}
