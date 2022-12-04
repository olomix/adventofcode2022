package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/olomix/adventofcode2022/internal"
)

type rng struct {
	start int
	end   int
}

func decodeRange(in string) rng {
	pair := strings.Split(in, "-")
	if len(pair) != 2 {
		panic("[assertion] len(pairs) != 2")
	}
	start, err := strconv.ParseInt(pair[0], 10, 64)
	internal.Perr(err)
	end, err := strconv.ParseInt(pair[1], 10, 64)
	internal.Perr(err)

	return rng{int(start), int(end)}
}

func decodeRanges(in string) (rng, rng) {
	pairs := strings.Split(in, ",")
	if len(pairs) != 2 {
		panic("[assertion] len(pairs) != 2")
	}

	pair1 := decodeRange(pairs[0])
	pair2 := decodeRange(pairs[1])
	return pair1, pair2
}

func main() {
	sum1 := 0
	sum2 := 0

	for txt := range internal.ReadLines("day04/input.txt") {
		if txt == "" {
			continue
		}

		rng1, rng2 := decodeRanges(txt)

		if rng1.start >= rng2.start && rng1.end <= rng2.end {
			sum1++
		} else if rng2.start >= rng1.start && rng2.end <= rng1.end {
			sum1++
		}

		// check if pair is overlapping
		if rng1.start <= rng2.start && rng1.end >= rng2.start {
			sum2++
		} else if rng2.start <= rng1.start && rng2.end >= rng1.start {
			sum2++
		}
	}

	fmt.Printf("sum1: %d\n", sum1)
	fmt.Printf("sum2: %d\n", sum2)
}
