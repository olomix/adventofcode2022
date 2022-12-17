package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/olomix/adventofcode2022/internal"
)

type valve struct {
	title       string
	rate        int
	connections []string
}

var allValves = make(map[string]*valve)

var valveRE = regexp.MustCompile(
	`Valve ([A-Z]{2}) has flow rate=(\d+); tunnels? leads? to valves? ([A-Z ,]+)$`)

func splitValves(s string) []string {
	return strings.Split(s, ", ")
}

func parseRate(in string) int {
	rate, err := strconv.Atoi(in)
	if err != nil {
		panic(err)
	}
	return rate
}

var weights = make(map[string]map[string]int)

func main() {
	for txt := range internal.ReadLines("day16/input.txt") {
		if txt == "" {
			continue
		}

		m := valveRE.FindStringSubmatch(txt)
		if len(m) == 0 {
			panic(1)
		}

		v := valve{
			title:       m[1],
			rate:        parseRate(m[2]),
			connections: splitValves(m[3]),
		}

		allValves[v.title] = &v
	}

	weights["AA"] = bfs2("AA")
	var valvesWithRate []string
	for _, v := range allValves {
		if v.rate > 0 {
			weights[v.title] = bfs2(v.title)
			valvesWithRate = append(valvesWithRate, v.title)
		}
	}

	valvesRates = make([]int, len(valvesWithRate)+1)

	// calculate int index for each valve
	valvesIndexes["AA"] = 0 // AA should always have index 0 to be able to skip it
	for v := range weights {
		if v == "AA" {
			continue
		}
		valvesIndexes[v] = len(valvesIndexes)
		valvesRates[valvesIndexes[v]] = allValves[v].rate
	}

	// fill depthMatrix
	ln := len(valvesIndexes)
	depthMatrix = make([][]int, ln)
	for i := range depthMatrix {
		depthMatrix[i] = make([]int, ln)
	}
	for from, tos := range weights {
		for to, depth := range tos {
			depthMatrix[valvesIndexes[from]][valvesIndexes[to]] = depth
		}
	}

	start := time.Now()
	fmt.Println(lookForMaxPath(30, 0, 1))
	fmt.Println("done in", time.Since(start))
	maxNumVariants := (1<<uint(len(valvesWithRate)) - 1) << 1
	start = time.Now()
	max := 0
	for i := 1; i < maxNumVariants; i += 2 {
		m := lookForMaxPath(26, 0, uint32(i)) +
			lookForMaxPath(26, 0, uint32(maxNumVariants^i))
		if m > max {
			max = m
		}
	}
	fmt.Println(max)
	fmt.Println("done in", time.Since(start))

}

var valvesIndexes = map[string]int{}

// matrix of path weights from one valve to another indexed by valve index
// from valvesIndexes
var depthMatrix [][]int

var valvesRates []int

var cache = map[int]int{}

func lookForMaxPath(timeLeft int, currentValueIdx int, valvesOn uint32) int {
	cacheKey := (timeLeft << 48) | (currentValueIdx << 32) | int(valvesOn)
	if v, ok := cache[cacheKey]; ok {
		return v
	}

	max := 0

	for kIdx := 1; kIdx < len(depthMatrix); kIdx++ {
		mask := uint32(1) << kIdx
		if valvesOn&mask > 0 {
			continue
		}

		newTimeLeft := timeLeft - depthMatrix[currentValueIdx][kIdx] - 1
		kWeight := valvesRates[kIdx] * newTimeLeft
		kWeight += lookForMaxPath(newTimeLeft, kIdx, valvesOn|mask)
		if kWeight > max {
			max = kWeight
		}
	}

	cache[cacheKey] = max
	return max
}

type queueEntry struct {
	valve string
	depth int
}

func bfs2(source string) map[string]int {
	depths := make(map[string]int)
	queue := []queueEntry{{valve: source, depth: 0}}
	visited := map[string]bool{source: true}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		currentValve := allValves[current.valve]
		for _, v := range currentValve.connections {
			if visited[v] {
				continue
			}
			visited[v] = true
			queue = append(
				queue,
				queueEntry{valve: v, depth: current.depth + 1})
			if allValves[v].rate > 0 {
				depths[v] = current.depth + 1
			}
		}
	}
	return depths
}
