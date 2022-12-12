package main

import (
	"fmt"
	"sort"

	"github.com/olomix/adventofcode2022/internal"
)

type coord struct {
	row, col int
}

func (c *coord) possibleMoves() []coord {
	var res [4]coord
	i := 0
	if c.col > 0 && field[c.row][c.col-1] <= field[c.row][c.col]+1 {
		res[i] = coord{c.row, c.col - 1}
		i++
	}
	if c.col < len(field[c.row])-1 &&
		field[c.row][c.col+1] <= field[c.row][c.col]+1 {
		res[i] = coord{c.row, c.col + 1}
		i++
	}
	if c.row > 0 &&
		field[c.row-1][c.col] <= field[c.row][c.col]+1 {
		res[i] = coord{c.row - 1, c.col}
		i++
	}
	if c.row < len(field)-1 &&
		field[c.row+1][c.col] <= field[c.row][c.col]+1 {
		res[i] = coord{c.row + 1, c.col}
		i++
	}
	return res[:i]
}

type queueItem struct {
	c      coord
	weight int
}

func bfs(n coord, want coord) int {
	queue := []queueItem{{c: n, weight: 0}}
	seen := map[coord]bool{}
	for len(queue) > 0 {
		c := queue[0]
		queue = queue[1:]
		if seen[c.c] {
			continue
		}
		seen[c.c] = true
		if c.c == want {
			return c.weight
		}
		for _, cc := range c.c.possibleMoves() {
			if !seen[cc] {
				queue = append(queue, queueItem{c: cc, weight: c.weight + 1})
			}
		}
	}
	return -1
}

var field [][]byte

func main() {
	rowIdx := 0
	var startCoord, destCoord coord
	for txt := range internal.ReadLines("day12/input.txt") {
		if txt == "" {
			continue
		}

		var row []byte
		for colIdx := 0; colIdx < len(txt); colIdx++ {
			val := txt[colIdx]
			if val == 'S' {
				startCoord = coord{row: rowIdx, col: colIdx}
				val = 'a'
			}
			if val == 'E' {
				destCoord = coord{row: rowIdx, col: colIdx}
				val = 'z'
			}
			row = append(row, val)
		}
		field = append(field, row)
		rowIdx++
	}

	fmt.Println(bfs(startCoord, destCoord))

	allPathsWeights := []int{}
	for row := 0; row < len(field); row++ {
		for col := 0; col < len(field[row]); col++ {
			if field[row][col] == 'a' {
				startCoord.row = row
				startCoord.col = col
				weight := bfs(startCoord, destCoord)
				if weight == -1 {
					continue
				}
				allPathsWeights = append(allPathsWeights, weight)
			}
		}
	}
	sort.Ints(allPathsWeights)
	fmt.Println(allPathsWeights[0])
}
