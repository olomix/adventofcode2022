package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/olomix/adventofcode2022/internal"
)

func allDifferent(buf []byte) bool {
	var buf2 = make([]byte, len(buf))
	copy(buf2[:], buf[:])
	sort.Slice(buf2[:], func(i, j int) bool {
		return buf2[i] < buf2[j]
	})

	for i := 0; i < len(buf2)-1; i++ {
		if buf2[i] == buf2[i+1] {
			return false
		}
	}

	return true
}

func findRnd(input []byte, n int) int {
	var buf = make([]byte, n)
	for j := range input {
		if j < n {
			buf[j] = input[j]
		} else {
			copy(buf[0:n-1], buf[1:n])
			buf[n-1] = input[j]
		}

		if j < n-1 {
			continue
		}

		if allDifferent(buf) {
			return j + 1
		}
	}
	panic("not found")
}

func main() {
	input, err := os.ReadFile("day06/input.txt")
	internal.Perr(err)
	fmt.Println(findRnd(input, 4))
	fmt.Println(findRnd(input, 14))
}
