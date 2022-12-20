package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/olomix/adventofcode2022/internal"
)

var lineRE = regexp.MustCompile(`^Blueprint (\d+): Each ore robot costs (\d+) ore. Each clay robot costs (\d+) ore. Each obsidian robot costs (\d+) ore and (\d+) clay. Each geode robot costs (\d+) ore and (\d+) obsidian.$`)

func mustInt(in string) int {
	i, err := strconv.Atoi(in)
	if err != nil {
		panic(err)
	}
	return i
}

var oreRobotCostOre = []int{}
var clayRobotCostOre = []int{}
var obsidianRobotCostOre = []int{}
var obsidianRobotCostClay = []int{}
var geodeRobotCostOre = []int{}
var geodeRobotCostObsidian = []int{}

func main() {
	i := 0

	for txt := range internal.ReadLines("day19/input.txt") {
		if txt == "" {
			continue
		}

		m := lineRE.FindStringSubmatch(txt)
		if mustInt(m[1]) != i+1 {
			panic(1)
		}

		oreRobotCostOre = append(oreRobotCostOre, mustInt(m[2]))
		clayRobotCostOre = append(clayRobotCostOre, mustInt(m[3]))
		obsidianRobotCostOre = append(obsidianRobotCostOre, mustInt(m[4]))
		obsidianRobotCostClay = append(obsidianRobotCostClay, mustInt(m[5]))
		geodeRobotCostOre = append(geodeRobotCostOre, mustInt(m[6]))
		geodeRobotCostObsidian = append(geodeRobotCostObsidian, mustInt(m[7]))

		i++
	}

	sum := 0
	for i := 0; i < len(oreRobotCostOre); i++ {
		sum += emulate2(state{oreRobots: 1}, i, 24) * (i + 1)
	}
	fmt.Println("step1: ", sum)

	i1 := emulate2(state{oreRobots: 1}, 0, 32)
	//fmt.Println(i1)
	i2 := emulate2(state{oreRobots: 1}, 1, 32)
	//fmt.Println(i2)
	i3 := emulate2(state{oreRobots: 1}, 2, 32)
	//fmt.Println(i3)
	fmt.Println("step2: ", i1*i2*i3)
}

type state struct {
	oreRobots      int
	clayRobots     int
	obsidianRobots int
	geodeRobots    int
	haveOre        int
	haveClay       int
	haveObsidian   int
	haveGeode      int
	minute         int
}

func emulate2(s state, bI int, min int) int {
	if s.minute == min {
		return s.haveGeode
	}

	//s := state{oreRobots: 1}
	canWaitForObsidianRobot := s.clayRobots > 0
	canWaitForGeodeRobot := s.obsidianRobots > 0

	// try to produce ore robot
	way0 := 0
	if s.oreRobots < 5 {
		newState := skipTo(s, oreRobotCostOre[bI], 0, 0, min)
		way0 = newState.haveGeode
		if newState.minute < min {
			newState.haveOre += newState.oreRobots - oreRobotCostOre[bI]
			newState.haveClay += newState.clayRobots
			newState.haveObsidian += newState.obsidianRobots
			newState.haveGeode += newState.geodeRobots
			newState.oreRobots++
			newState.minute++
			way0 = emulate2(newState, bI, min)
		}
	}

	// try to produce clay robot
	newState := skipTo(s, clayRobotCostOre[bI], 0, 0, min)
	way1 := newState.haveGeode
	if newState.minute < min {
		newState.haveOre += newState.oreRobots - clayRobotCostOre[bI]
		newState.haveClay += newState.clayRobots
		newState.haveObsidian += newState.obsidianRobots
		newState.haveGeode += newState.geodeRobots
		newState.clayRobots++
		newState.minute++
		way1 = emulate2(newState, bI, min)
	}

	// try to produce obsidian robot
	way2 := 0
	if canWaitForObsidianRobot {
		newState = skipTo(s, obsidianRobotCostOre[bI],
			obsidianRobotCostClay[bI], 0, min)
		way2 = newState.haveGeode
		if newState.minute < min {
			newState.haveOre += newState.oreRobots - obsidianRobotCostOre[bI]
			newState.haveClay += newState.clayRobots - obsidianRobotCostClay[bI]
			newState.haveObsidian += newState.obsidianRobots
			newState.haveGeode += newState.geodeRobots
			newState.obsidianRobots++
			newState.minute++
			way2 = emulate2(newState, bI, min)
		}
	}

	// try to produce geode robot
	way3 := 0
	if canWaitForGeodeRobot {
		newState = skipTo(s, geodeRobotCostOre[bI], 0,
			geodeRobotCostObsidian[bI], min)
		way3 = newState.haveGeode
		if newState.minute < min {
			newState.haveOre += newState.oreRobots - geodeRobotCostOre[bI]
			newState.haveClay += newState.clayRobots
			newState.haveObsidian += newState.obsidianRobots - geodeRobotCostObsidian[bI]
			newState.haveGeode += newState.geodeRobots
			newState.geodeRobots++
			newState.minute++
			way3 = emulate2(newState, bI, min)
		}
	}

	max := way0
	if way1 > max {
		max = way1
	}
	if way2 > max {
		max = way2
	}
	if way3 > max {
		max = way3
	}
	return max
}

func skipTo(in state, wantOre, wantClay, wantObsidian int, min int) state {
	for {
		if in.minute >= min {
			return in
		}
		if in.haveOre >= wantOre &&
			in.haveClay >= wantClay &&
			in.haveObsidian >= wantObsidian {
			return in
		}
		in.minute++
		in.haveOre += in.oreRobots
		in.haveClay += in.clayRobots
		in.haveObsidian += in.obsidianRobots
		in.haveGeode += in.geodeRobots
	}
}
