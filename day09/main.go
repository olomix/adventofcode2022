package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/olomix/adventofcode2022/internal"
)

type coord struct {
	x, y int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func moveTail(hCoord, tCoord coord) coord {
	if abs(hCoord.x-tCoord.x) < 2 && abs(hCoord.y-tCoord.y) < 2 {
		return tCoord
	}

	/* cases for head position
	13 10 11 12 14
	 9  .  ^  .  1
	 8  <  T  >  2
	 7  .  v  .  3
	16  6  5  4 15
	*/
	// 1
	if hCoord.x == tCoord.x+2 && hCoord.y == tCoord.y+1 {
		return coord{tCoord.x + 1, tCoord.y + 1}
	}
	// 2
	if hCoord.x == tCoord.x+2 && hCoord.y == tCoord.y {
		return coord{tCoord.x + 1, tCoord.y}
	}
	// 3
	if hCoord.x == tCoord.x+2 && hCoord.y == tCoord.y-1 {
		return coord{tCoord.x + 1, tCoord.y - 1}
	}
	// 4
	if hCoord.x == tCoord.x+1 && hCoord.y == tCoord.y-2 {
		return coord{tCoord.x + 1, tCoord.y - 1}
	}
	// 5
	if hCoord.x == tCoord.x && hCoord.y == tCoord.y-2 {
		return coord{tCoord.x, tCoord.y - 1}
	}
	// 6
	if hCoord.x == tCoord.x-1 && hCoord.y == tCoord.y-2 {
		return coord{tCoord.x - 1, tCoord.y - 1}
	}
	// 7
	if hCoord.x == tCoord.x-2 && hCoord.y == tCoord.y-1 {
		return coord{tCoord.x - 1, tCoord.y - 1}
	}
	// 8
	if hCoord.x == tCoord.x-2 && hCoord.y == tCoord.y {
		return coord{tCoord.x - 1, tCoord.y}
	}
	// 9
	if hCoord.x == tCoord.x-2 && hCoord.y == tCoord.y+1 {
		return coord{tCoord.x - 1, tCoord.y + 1}
	}
	// 10
	if hCoord.x == tCoord.x-1 && hCoord.y == tCoord.y+2 {
		return coord{tCoord.x - 1, tCoord.y + 1}
	}
	// 11
	if hCoord.x == tCoord.x && hCoord.y == tCoord.y+2 {
		return coord{tCoord.x, tCoord.y + 1}
	}
	// 12
	if hCoord.x == tCoord.x+1 && hCoord.y == tCoord.y+2 {
		return coord{tCoord.x + 1, tCoord.y + 1}
	}
	// 13
	if hCoord.x == tCoord.x-2 && hCoord.y == tCoord.y+2 {
		return coord{tCoord.x - 1, tCoord.y + 1}
	}
	// 14
	if hCoord.x == tCoord.x+2 && hCoord.y == tCoord.y+2 {
		return coord{tCoord.x + 1, tCoord.y + 1}
	}
	// 15
	if hCoord.x == tCoord.x+2 && hCoord.y == tCoord.y-2 {
		return coord{tCoord.x + 1, tCoord.y - 1}
	}
	// 16
	if hCoord.x == tCoord.x-2 && hCoord.y == tCoord.y-2 {
		return coord{tCoord.x - 1, tCoord.y - 1}
	}
	panic(fmt.Sprintf("no way; head: %v, tail: %v", hCoord, tCoord))
}

func moveRope(ropeCoords []coord, direction byte) {
	switch direction {
	case 'R':
		ropeCoords[0] = coord{ropeCoords[0].x + 1, ropeCoords[0].y}
	case 'L':
		ropeCoords[0] = coord{ropeCoords[0].x - 1, ropeCoords[0].y}
	case 'U':
		ropeCoords[0] = coord{ropeCoords[0].x, ropeCoords[0].y + 1}
	case 'D':
		ropeCoords[0] = coord{ropeCoords[0].x, ropeCoords[0].y - 1}
	default:
		panic("unknown direction")
	}
	for i := 1; i < len(ropeCoords); i++ {
		ropeCoords[i] = moveTail(ropeCoords[i-1], ropeCoords[i])
	}
}

func main() {
	var bigRope [10]coord
	var smallRope [2]coord

	smallSeen := make(map[coord]struct{})
	bigSeen := make(map[coord]struct{})

	for txt := range internal.ReadLines("day09/input.txt") {
		if txt == "" {
			continue
		}

		parts := strings.Split(txt, " ")
		if len(parts) != 2 {
			panic("invalid input")
		}
		if len(parts[0]) != 1 {
			panic("invalid input")
		}
		direction := parts[0][0]
		steps, err := strconv.ParseInt(parts[1], 10, 64)
		internal.Perr(err)
		for i := int64(0); i < steps; i++ {
			moveRope(smallRope[:], direction)
			smallSeen[smallRope[len(smallRope)-1]] = struct{}{}

			moveRope(bigRope[:], direction)
			bigSeen[bigRope[len(bigRope)-1]] = struct{}{}
		}
	}

	// 6522
	fmt.Println(len(smallSeen))
	// 2717
	fmt.Println(len(bigSeen))
}
