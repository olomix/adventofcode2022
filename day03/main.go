package main

import (
	"fmt"

	"github.com/olomix/adventofcode2022/internal"
)

func priority(in rune) int {
	if in >= 'a' && in <= 'z' {
		return int(in) - int('a') + 1
	}
	if in >= 'A' && in <= 'Z' {
		return int(in) - int('A') + 27
	}
	panic("[assertion] in >= 'a' && in <= 'z' || in >= 'A' && in <= 'Z'")
}

func commonItem(items [3]string) rune {
	for _, c1 := range items[0] {
		for _, c2 := range items[1] {
			if c1 == c2 {
				for _, c3 := range items[2] {
					if c1 == c3 {
						return c1
					}
				}
			}
		}
	}
	panic("[assertion] no common item found")
}

func main() {
	sum1 := 0
	sum2 := 0

	var groupItems [3]string

	lineNo := 0

LOOP:
	for txt := range internal.ReadLines("day03/input.txt") {
		if txt == "" {
			continue
		}

		if len(txt)%2 != 0 {
			panic("[assertion] len(txt) % 2 != 0")
		}

		groupMemberID := lineNo % 3
		lineNo++
		groupItems[groupMemberID] = txt
		if groupMemberID == 2 {
			sum2 += priority(commonItem(groupItems))
		}

		middle := len(txt) / 2
		for i := 0; i < middle; i++ {

			for j := middle; j < len(txt); j++ {
				if txt[i] == txt[j] {
					sum1 += priority(rune(txt[i]))
					continue LOOP
				}
			}
		}
	}

	fmt.Printf("sum1: %d\n", sum1)
	fmt.Printf("sum2: %d\n", sum2)
}
