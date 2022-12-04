package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/olomix/adventofcode2022/internal"
)

func main() {
	sum1 := 0
	sum2 := 0

	for txt := range internal.ReadLines("day04/input.txt") {
		if txt == "" {
			continue
		}

		pairs := strings.Split(txt, ",")
		if len(pairs) != 2 {
			panic("[assertion] len(pairs) != 2")
		}

		pair1 := strings.Split(pairs[0], "-")
		if len(pair1) != 2 {
			panic("[assertion] len(pairs1) != 2")
		}
		pair1Start, err := strconv.ParseInt(pair1[0], 10, 64)
		internal.Perr(err)
		pair1End, err := strconv.ParseInt(pair1[1], 10, 64)
		internal.Perr(err)

		pair2 := strings.Split(pairs[1], "-")
		if len(pair2) != 2 {
			panic("[assertion] len(pairs2) != 2")
		}
		pair2Start, err := strconv.ParseInt(pair2[0], 10, 64)
		internal.Perr(err)
		pair2End, err := strconv.ParseInt(pair2[1], 10, 64)
		internal.Perr(err)

		if pair1Start >= pair2Start && pair1End <= pair2End {
			sum1++
		} else if pair2Start >= pair1Start && pair2End <= pair1End {
			sum1++
		}

		// check if pair is overlapping
		if pair1Start <= pair2Start && pair1End >= pair2Start {
			sum2++
		} else if pair2Start <= pair1Start && pair2End >= pair1Start {
			sum2++
		}
	}

	fmt.Printf("sum1: %d\n", sum1)
	fmt.Printf("sum2: %d\n", sum2)
}
