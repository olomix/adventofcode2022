package main

import (
	"fmt"
	"strconv"

	"github.com/olomix/adventofcode2022/internal"
)

type node struct {
	val  int
	prev *node
	next *node
}

func mustInt(in string) int {
	out, err := strconv.Atoi(in)
	if err != nil {
		panic(err)
	}
	return out
}

func main() {
	var input []int
	for txt := range internal.ReadLines("day20/input.txt") {
		if txt == "" {
			continue
		}
		input = append(input, mustInt(txt))
	}

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func findSum(list *node) int {
	zero := findZero(list)
	sum := 0
	for i := 1; i <= 3000; i++ {
		zero = zero.next
		if i == 1000 {
			sum += zero.val
		}
		if i == 2000 {
			sum += zero.val
		}
		if i == 3000 {
			sum += zero.val
		}
	}
	return sum
}

func part1(in []int) int {
	start := makeLinkedList(in)
	original := listToSlice(start)
	mix(original)
	return findSum(start)
}

func part2(in []int) int {
	for i := range in {
		in[i] = in[i] * 811589153
	}
	start := makeLinkedList(in)
	original := listToSlice(start)
	for z := 0; z < 10; z++ {
		mix(original)
	}

	return findSum(start)
}

func mix(nodes []*node) {
	for i, n := range nodes {
		_ = i
		if n.val == 0 {
			continue
		}

		n.prev.next = n.next
		n.next.prev = n.prev

		var after *node
		if n.val > 0 {
			v := n.val % (len(nodes) - 1)
			after = n
			for i := 0; i < v; i++ {
				after = after.next
			}
		} else if n.val < 0 {
			v := n.val % (len(nodes) - 1)
			after = n.prev
			for i := 0; i < -v; i++ {
				after = after.prev
			}
		}
		n.prev.next = n.next
		n.next.prev = n.prev
		n.next = after.next
		n.prev = after
		n.next.prev = n
		after.next = n
	}
}

func makeLinkedList(in []int) *node {
	var prev *node
	var start *node
	for _, v := range in {
		n := &node{
			val:  v,
			prev: prev,
		}
		if prev != nil {
			prev.next = n
		}
		prev = n
		if start == nil {
			start = n
		}
	}
	prev.next = start
	start.prev = prev
	return start
}

func listToSlice(in *node) []*node {
	var out []*node
	for {
		out = append(out, in)
		if in.next == out[0] {
			break
		}
		in = in.next
	}
	return out
}

func findZero(in *node) *node {
	for {
		if in.val == 0 {
			return in
		}
		in = in.next
	}
}
