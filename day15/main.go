package main

import (
	"fmt"
	"math/big"
	"regexp"
	"sort"
	"strconv"

	"github.com/olomix/adventofcode2022/internal"
)

type coord struct{ x, y int }
type sb struct {
	s coord
	b coord
}

type interval struct{ start, end int }

var RE = regexp.MustCompile(`^Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)$`)

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func main() {
	var all []sb
	lineNo := 2000000
	maxCoord := 4000000
	for txt := range internal.ReadLines("day15/input.txt") {
		if txt == "" {
			continue
		}

		m := RE.FindStringSubmatch(txt)
		if m == nil {
			panic(txt)
		}

		signal := coord{toInt(m[1]), toInt(m[2])}
		beacon := coord{toInt(m[3]), toInt(m[4])}
		sbI := sb{s: signal, b: beacon}
		all = append(all, sbI)
	}

	allItems := make(map[int]struct{})

	beaconsOnLine := []coord{}
	for _, sbI := range all {
		if sbI.b.y == lineNo {
			beaconsOnLine = append(beaconsOnLine, sbI.b)
		}
	}

	for _, sbI := range all {
		x1, x2, err := lineBounds(sbI, lineNo)
		if err != nil {
			continue
		}
		for i := x1; i <= x2; i++ {
			beaconInCell := false
			for _, b := range beaconsOnLine {
				if b.x == i {
					beaconInCell = true
				}
			}
			if !beaconInCell {
				allItems[i] = struct{}{}
			}
		}
	}
	fmt.Println(len(allItems))

	is2 := make([]interval, len(all))
	for row := 0; row <= maxCoord; row++ {
		i := 0
		for _, sbI := range all {
			var err error
			is2[i].start, is2[i].end, err = lineBounds(sbI, row)
			if err != nil {
				continue
			}
			i += 1
		}
		c := optimizeBounds(is2[:i])
		if len(c) > 1 {
			x := new(big.Int).SetInt64(int64(c[0].end + 1))
			z := new(big.Int).Mul(x, big.NewInt(4000000))
			z = z.Add(z, big.NewInt(int64(row)))
			fmt.Printf("%v\n", z.Text(10))
			return
		}
	}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func mod(y int) int {
	if y < 0 {
		return -y
	}
	return y
}

func lineBounds(sb1 sb, lineNo int) (int, int, error) {
	b := sb1.b
	s := sb1.s
	var y1, y2 int
	if b.y < s.y {
		y1 = b.y - mod(s.x-b.x)
		y2 = s.y + (s.y - y1)
	} else {
		y2 = b.y + mod(s.x-b.x)
		y1 = s.y - (y2 - s.y)
	}
	if lineNo < y1 || lineNo > y2 {
		return 0, 0,
			fmt.Errorf("line %v not in bounds [%v, %v]", lineNo, y1, y2)
	}
	halfInterval := y2 - s.y - mod(s.y-lineNo)
	return s.x - halfInterval, s.x + halfInterval, nil
}

type boundSort []interval

func (b boundSort) Len() int {
	return len(b)
}

func (b boundSort) Less(i, j int) bool {
	if b[i].start != b[j].start {
		return b[i].start < b[j].start
	}
	return b[i].end < b[j].end
}

func (b boundSort) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

// accept bound indexes
func optimizeBounds(bounds []interval) []interval {
	if len(bounds) < 2 {
		return bounds
	}
	sort.Sort(boundSort(bounds))
	src := 0
	dst := 1
	for dst < len(bounds) {
		if bounds[dst].start <= bounds[src].end {
			bounds[src].end = max(bounds[src].end, bounds[dst].end)
		} else {
			src++
			if src != dst {
				bounds[src] = bounds[dst]
			}
		}
		dst++
	}
	return bounds[:src+1]
}
