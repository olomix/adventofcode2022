package main

import (
	"fmt"
	"math"
	"sort"

	"github.com/olomix/adventofcode2022/internal"
)

type node struct {
	x, y         int
	pathWeight   int
	destinations []*node
}

func (n *node) coord() [2]int {
	return [2]int{n.y, n.x}
}

func (n *node) possibleWays() (res [][2]int) {
	if n.x > 0 && field[n.y][n.x-1] <= field[n.y][n.x]+1 {
		res = append(res, [2]int{n.y, n.x - 1})
	}
	if n.x < len(field[n.y])-1 && field[n.y][n.x+1] <= field[n.y][n.x]+1 {
		res = append(res, [2]int{n.y, n.x + 1})
	}
	if n.y > 0 && field[n.y-1][n.x] <= field[n.y][n.x]+1 {
		res = append(res, [2]int{n.y - 1, n.x})
	}
	if n.y < len(field)-1 && field[n.y+1][n.x] <= field[n.y][n.x]+1 {
		res = append(res, [2]int{n.y + 1, n.x})
	}
	return
}

func allNodes(n *node) []*node {
	seen := map[[2]int]bool{}
	var all []*node
	queue := []*node{n}
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		if seen[n.coord()] {
			continue
		}
		seen[n.coord()] = true
		all = append(all, n)
		for _, nn := range n.destinations {
			if !seen[nn.coord()] {
				queue = append(queue, nn)
			}
		}
	}
	return all
}

// Dijkstra
func try3(n *node, want *node) int {
	all := allNodes(n)
	weights := map[[2]int]int{}
	for _, n := range all {
		weights[n.coord()] = math.MaxInt
	}
	weights[n.coord()] = 0
	seen := map[[2]int]bool{}
	for len(all) > 0 {
		sort.Slice(all, func(i, j int) bool {
			return weights[all[i].coord()] < weights[all[j].coord()]
		})

		n := all[0]
		all = all[1:]
		seen[n.coord()] = true

		newWeight := weights[n.coord()] + 1
		for _, nn := range n.destinations {
			if !seen[nn.coord()] && newWeight < weights[nn.coord()] {
				weights[nn.coord()] = newWeight
			}
		}
	}
	return weights[want.coord()]
}

var field [][]byte

func buildGraph() {
	for rowIdx := range field {
		for colIdx := range field[rowIdx] {
			n := getNode(rowIdx, colIdx)
			for _, coord := range n.possibleWays() {
				n.destinations = append(n.destinations,
					getNode(coord[0], coord[1]))
			}
		}
	}
}

var nodes = map[[2]int]*node{}

func getNode(row, col int) *node {
	key := [2]int{row, col}
	n, seen := nodes[key]
	if !seen {
		n = &node{x: col, y: row}
		nodes[key] = n
	}
	return n
}

func main() {
	rowIdx := 0
	var curPos [2]int
	var destPos [2]int
	for txt := range internal.ReadLines("day12/input.txt") {
		if txt == "" {
			continue
		}

		var row []byte
		for colIdx := 0; colIdx < len(txt); colIdx++ {
			val := txt[colIdx]
			if val == 'S' {
				curPos = [2]int{rowIdx, colIdx}
				val = 'a'
			}
			if val == 'E' {
				destPos = [2]int{rowIdx, colIdx}
				val = 'z'
			}
			row = append(row, val)
		}
		field = append(field, row)
		rowIdx++
	}
	buildGraph()
	startNode := getNode(curPos[0], curPos[1])
	wantNode := getNode(destPos[0], destPos[1])
	fmt.Println(try3(startNode, wantNode))
	allPathsWeights := []int{}
	for row := 0; row < len(field); row++ {
		col := 0
		// only from first column we can reach destination
		//for col := 0; col < len(field[row]); col++ {
		if field[row][col] == 'a' {
			startNode := getNode(row, col)
			weight := try3(startNode, wantNode)
			//fmt.Printf("%v: %v\n", startNode.coord(), weight)
			allPathsWeights = append(allPathsWeights, weight)
		}
		//}
	}
	sort.Ints(allPathsWeights)
	fmt.Println(allPathsWeights[0])
}
