package main

import (
	"fmt"
	"strings"

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
		if strings.ContainsRune(items[1], c1) &&
			strings.ContainsRune(items[2], c1) {
			return c1
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
			if strings.ContainsRune(txt[middle:], rune(txt[i])) {
				sum1 += priority(rune(txt[i]))
				continue LOOP
			}
		}
	}

	fmt.Printf("sum1: %d\n", sum1) // 7597
	fmt.Printf("sum2: %d\n", sum2) // 2607
}
