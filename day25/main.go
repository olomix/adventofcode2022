package main

import (
	"fmt"

	"github.com/olomix/adventofcode2022/internal"
)

func main() {
	var total uint
	for txt := range internal.ReadLines("day25/input.txt") {
		if txt == "" {
			continue
		}
		total += snafuToInt(txt)
	}
	fmt.Println(toSnafu(total))
}

func toSnafu(in uint) string {
	e := uint(5)
	overflow := false
	var out []byte
	for in > 0 {
		x := in % e
		in = in / e
		switch {
		case x == 0:
			if overflow {
				out = append(out, '1')
			} else {
				out = append(out, '0')
			}
			overflow = false
		case x == 1:
			if overflow {
				out = append(out, '2')
			} else {
				out = append(out, '1')
			}
			overflow = false
		case x == 2:
			if overflow {
				out = append(out, '=')
			} else {
				out = append(out, '2')
			}
		case x == 3:
			if overflow {
				out = append(out, '-')
			} else {
				out = append(out, '=')
				overflow = true
			}
		case x == 4:
			if overflow {
				out = append(out, '0')
			} else {
				out = append(out, '-')
			}
			overflow = true
		}
	}
	if overflow {
		out = append(out, '1')
	}
	for i, j := 0, len(out)-1; i < j; i, j = i+1, j-1 {
		out[i], out[j] = out[j], out[i]
	}
	return string(out)
}

func snafuToInt(in string) uint {
	data := []byte(in)
	e := 1
	for i := 0; i < len(data)-1; i++ {
		e = e * 5
	}
	var result int
	for len(data) > 0 {
		var step int
		switch data[0] {
		case '0':
			step = 0
		case '1':
			step = 1
		case '2':
			step = 2
		case '=':
			step = -2
		case '-':
			step = -1
		default:
			panic("invalid input")
		}
		result += step * e
		e = e / 5
		data = data[1:]
	}
	return uint(result)
}
