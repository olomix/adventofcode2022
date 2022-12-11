package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/olomix/adventofcode2022/internal"
)

type monkey struct {
	items     []int
	op        func(int) int
	test      int
	direction [2]int
	cnt       int
}

var monkeyRE = regexp.MustCompile(`(?m)Monkey \d+:\s*
  Starting items: (.+)$
  Operation: new = (.+)$
  Test: divisible by (\d+)$
    If true: throw to monkey (\d+)$
    If false: throw to monkey (\d+)$
`)

func parseStartingItems(s string) []int {
	var res []int
	for _, i := range strings.Split(s, ", ") {
		i2, err := strconv.ParseInt(i, 10, 32)
		internal.Perr(err)
		res = append(res, int(i2))
	}
	return res
}

func parseOperation(s string) func(int) int {
	if !strings.HasPrefix(s, "old + ") && !strings.HasPrefix(s, "old * ") {
		panic("unknown operation: " + s)
	}
	switch {
	case s == "old * old":
		return func(i int) int { return i * i }
	case s == "old + old":
		return func(i int) int { return i + i }
	default:
		arg, err := strconv.ParseInt(s[6:], 10, 32)
		internal.Perr(err)
		if s[4] == '+' {
			return func(i int) int { return i + int(arg) }
		} else if s[4] == '*' {
			return func(i int) int { return i * int(arg) }
		}
	}
	panic("unknown operation: " + s)
}

func parseMonkeys(input string) []monkey {
	matches := monkeyRE.FindAllStringSubmatch(input, -1)
	monkeys := make([]monkey, len(matches))
	for i := range matches {
		monkeys[i].items = parseStartingItems(matches[i][1])
		monkeys[i].op = parseOperation(matches[i][2])
		monkeys[i].test = parseInt(matches[i][3])
		monkeys[i].direction[0] = parseInt(matches[i][4])
		monkeys[i].direction[1] = parseInt(matches[i][5])
	}
	return monkeys
}

func parseInt(s string) int {
	i, err := strconv.ParseInt(s, 10, 32)
	internal.Perr(err)
	return int(i)
}

func step1(monkeys []monkey, iterCount int, worryLevelDevidor int) int {
	divisors := []int{}
	for i := range monkeys {
		divisors = append(divisors, monkeys[i].test)
	}
	g := lcm(divisors[0], divisors[1], divisors[2:]...)
	for i := 0; i < iterCount; i++ {
		for j := range monkeys {
			for _, item := range monkeys[j].items {
				item = monkeys[j].op(item)
				item %= g
				item = item / worryLevelDevidor
				var newMonkeyIdx int
				if item%monkeys[j].test == 0 {
					newMonkeyIdx = monkeys[j].direction[0]
				} else {
					newMonkeyIdx = monkeys[j].direction[1]
				}
				monkeys[newMonkeyIdx].items = append(
					monkeys[newMonkeyIdx].items, item)
				monkeys[j].cnt++
			}
			monkeys[j].items = nil
		}
	}

	return mulGreatest(monkeys)
}

func mulGreatest(ms []monkey) int {
	var lst []int
	for i := range ms {
		lst = append(lst, ms[i].cnt)
	}

	sort.Ints(lst)
	return lst[len(lst)-1] * lst[len(lst)-2]
}

func main() {
	fileBytes, err := os.ReadFile("day11/input.txt")
	internal.Perr(err)
	monkeys1 := parseMonkeys(string(fileBytes))
	monkeys2 := parseMonkeys(string(fileBytes))
	res1 := step1(monkeys1, 20, 3)
	fmt.Println(res1)
	res2 := step1(monkeys2, 10000, 1)
	fmt.Println(res2)
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(a, b int, integers ...int) int {
	result := a * b / gcd(a, b)
	for i := 0; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}
	return result
}
