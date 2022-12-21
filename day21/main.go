package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/olomix/adventofcode2022/internal"
)

var numRE = regexp.MustCompile(`^([a-z]+): (\d+)$`)
var opsRE = regexp.MustCompile(`^([a-z]+): ([a-z]+) ([+\-*/]) ([a-z]+)$`)

func mustInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

func mustIntPtr(s string) *int {
	return &[]int{mustInt(s)}[0]
}

type opMonkey struct {
	left  string
	right string
	op    string
}

func main() {
	nodes := map[string]*node{}
	for txt := range internal.ReadLines("day21/input.txt") {
		if txt == "" {
			continue
		}
		m := numRE.FindStringSubmatch(txt)
		if m != nil {
			nodes[m[1]] = &node{
				name: m[1],
				num:  mustIntPtr(m[2]),
			}
			continue
		}
		m = opsRE.FindStringSubmatch(txt)
		if m == nil {
			panic(txt)
		}

		nodes[m[1]] = &node{
			name:  m[1],
			left:  m[2],
			right: m[4],
			op:    m[3],
		}
	}

	part1(nodes)
	part2(nodes)
}

func part1(nodes map[string]*node) {
	fmt.Println("part 1:", doMath("root", nodes))
}

type node struct {
	name  string
	op    string
	num   *int
	left  string
	right string
}

func part2(nodes map[string]*node) {
	root := nodes["root"]
	var equationSubtree string
	var wantNum int
	if containsHumn(root.left, nodes) {
		wantNum = doMath(root.right, nodes)
		equationSubtree = root.left
	}
	if containsHumn(root.right, nodes) {
		wantNum = doMath(root.left, nodes)
		equationSubtree = root.right
	}
	//fmt.Printf("Equation subtree is %v. Want: %v\n", equationSubtree, wantNum)
	fmt.Println("part 2:", resolveEquation(equationSubtree, wantNum, nodes))
}

func resolveEquation(opNode string, wantNum int, nodes map[string]*node) int {
	if opNode == "humn" {
		return wantNum
	}

	n := nodes[opNode]
	if containsHumn(n.left, nodes) {
		right := doMath(n.right, nodes)
		switch n.op {
		case "+":
			wantNum = wantNum - right
		case "-":
			wantNum = wantNum + right
		case "*":
			wantNum = wantNum / right
		case "/":
			wantNum = wantNum * right
		default:
			panic(n.op)
		}
		return resolveEquation(n.left, wantNum, nodes)
	}

	left := doMath(n.left, nodes)
	switch n.op {
	case "+":
		wantNum = wantNum - left
	case "-":
		wantNum = left - wantNum
	case "*":
		wantNum = wantNum / left
	case "/":
		wantNum = left / wantNum
	default:
		panic(n.op)
	}
	return resolveEquation(n.right, wantNum, nodes)
}

func containsHumn(n string, nodes map[string]*node) bool {
	if n == "humn" {
		return true
	}
	node := nodes[n]
	if node.num != nil {
		return false
	}
	return containsHumn(node.left, nodes) || containsHumn(node.right, nodes)
}

func doMath(n string, nodes map[string]*node) int {
	node := nodes[n]
	if node.num != nil {
		return *node.num
	}
	left := doMath(node.left, nodes)
	right := doMath(node.right, nodes)
	switch node.op {
	case "+":
		return left + right
	case "-":
		return left - right
	case "*":
		return left * right
	case "/":
		return left / right
	default:
		panic(node.op)
	}
}
