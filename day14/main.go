package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/olomix/adventofcode2022/internal"
)

type cord struct {
	x, y int
}

var paths [][]cord

var occupiedPoints = map[cord]struct{}{}

var MaxY = 0

func main() {
	for txt := range internal.ReadLines("day14/input.txt") {
		if txt == "" {
			continue
		}

		coords := strings.Split(txt, " -> ")
		var line []cord
		for _, coord := range coords {
			points := strings.Split(coord, ",")
			if len(points) != 2 {
				panic("1")
			}

			y := strToInt(points[1])
			x := strToInt(points[0])
			if y > MaxY {
				MaxY = y
			}
			line = append(line, cord{x: x, y: y})
		}
		paths = append(paths, line)
	}

	fillOccupiedPoints()
	c := 0
	for fallNext() {
		c++
	}
	fmt.Println(c)

	c2 := 0
	startPoint := cord{x: 500, y: 0}
	for {
		fallTillY(MaxY + 2)
		c2++
		if _, ok := occupiedPoints[startPoint]; ok {
			break
		}
	}
	fmt.Println(c2 + c)
}

// return true if sand tuns to rest of false if it goes away
func fallNext() bool {
	currentCoord := cord{x: 500, y: 0}
	for {
		if currentCoord.y+1 > MaxY {
			return false
		}
		nextCoord := cord{y: currentCoord.y + 1, x: currentCoord.x}
		if _, ok := occupiedPoints[nextCoord]; !ok {
			currentCoord = nextCoord
			continue
		}

		nextCoord.x -= 1
		if _, ok := occupiedPoints[nextCoord]; !ok {
			currentCoord = nextCoord
			continue
		}

		nextCoord.x += 2
		if _, ok := occupiedPoints[nextCoord]; !ok {
			currentCoord = nextCoord
			continue
		}

		occupiedPoints[currentCoord] = struct{}{}
		return true
	}
}

func fallTillY(maxY int) {
	currentCoord := cord{x: 500, y: 0}
	if _, ok := occupiedPoints[currentCoord]; ok {
		panic(4)
	}

	for {
		if currentCoord.y == maxY-1 {
			occupiedPoints[currentCoord] = struct{}{}
			return
		}

		nextCoord := cord{y: currentCoord.y + 1, x: currentCoord.x}
		if _, ok := occupiedPoints[nextCoord]; !ok {
			currentCoord = nextCoord
			continue
		}

		nextCoord.x -= 1
		if _, ok := occupiedPoints[nextCoord]; !ok {
			currentCoord = nextCoord
			continue
		}

		nextCoord.x += 2
		if _, ok := occupiedPoints[nextCoord]; !ok {
			currentCoord = nextCoord
			continue
		}

		occupiedPoints[currentCoord] = struct{}{}
		return
	}
}

func fillOccupiedPoints() {
	for i := range paths {
		path := paths[i]

		for j := 1; j < len(path); j++ {
			point1 := path[j-1]
			point2 := path[j]
			x1 := point1.x
			y1 := point1.y
			x2 := point2.x
			y2 := point2.y
			if x1 == x2 {
				if y1 < y2 {
					for yi := y1; yi <= y2; yi++ {
						occupiedPoints[cord{x: x1, y: yi}] = struct{}{}
					}
				} else {
					for yi := y2; yi <= y1; yi++ {
						occupiedPoints[cord{x: x1, y: yi}] = struct{}{}
					}
				}
			} else if y1 == y2 {
				if x1 < x2 {
					for xi := x1; xi <= x2; xi++ {
						occupiedPoints[cord{x: xi, y: y1}] = struct{}{}
					}
				} else {
					for xi := x2; xi <= x1; xi++ {
						occupiedPoints[cord{x: xi, y: y1}] = struct{}{}
					}
				}
			} else {
				panic(3)
			}
		}
	}
}

func strToInt(in string) int {
	i, err := strconv.ParseInt(in, 10, 64)
	internal.Perr(err)
	return int(i)
}
