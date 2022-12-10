package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/olomix/adventofcode2022/internal"
)

var sum = 0
var field [6][40]byte
var tickNo = 1
var regX = 1

func afterTickCallback() {
	accumulateSignalStrength()
	updateCRT()
}

func accumulateSignalStrength() {
	if tickNo >= 20 && tickNo <= 220 && (tickNo-20)%40 == 0 {
		sum += tickNo * regX
	}

	if tickNo == 220 {
		// 11780
		fmt.Printf("step 1: %v\n", sum)
	}
}

func printField() {
	// PZULBAUA
	fmt.Println("step 2")
	for i := range field {
		for j := range field[i] {
			fmt.Printf("%c", field[i][j])
		}
		fmt.Println()
	}
}

func updateCRT() {
	tickNo0 := tickNo - 1
	row, column := tickNo0/40%6, tickNo0%40
	sym := byte('.')
	if column >= regX-1 && column <= regX+1 {
		sym = '#'
	}
	field[row][column] = sym
}

func main() {
	for txt := range internal.ReadLines("day10/input.txt") {
		if txt == "" {
			continue
		}
		switch {
		case txt == "noop":
			tickNo++
			afterTickCallback()
		case strings.HasPrefix(txt, "addx "):
			arg, err := strconv.Atoi(txt[5:])
			internal.Perr(err)
			tickNo++
			afterTickCallback()
			tickNo++
			regX += arg
			afterTickCallback()
		default:
			panic("unknown command")
		}
	}
	printField()
}
