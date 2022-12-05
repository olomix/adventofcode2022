package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/olomix/adventofcode2022/internal"
)

var field1 [9][]byte
var field2 [9][]byte

var moveRE = regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)

func parseMove(in string) (n, from, to int) {
	ss := moveRE.FindStringSubmatch(in)
	if len(ss) != 4 {
		panic("invalid move")
	}
	n64, err := strconv.ParseInt(ss[1], 10, 64)
	internal.Perr(err)
	from64, err := strconv.ParseInt(ss[2], 10, 64)
	internal.Perr(err)
	from64--
	to64, err := strconv.ParseInt(ss[3], 10, 64)
	internal.Perr(err)
	to64--

	return int(n64), int(from64), int(to64)
}

func main() {
	i := 0
	for txt := range internal.ReadLines("day05/input.txt") {
		i += 1
		if txt == "" {
			continue
		}

		if i < 9 {
			// read fields
			columnIdx := -1
			for i2 := 1; i2 < 36; i2 += 4 {
				columnIdx++
				if txt[i2] == ' ' {
					continue
				}
				field1[columnIdx] = append([]byte{txt[i2]},
					field1[columnIdx]...)
				field2[columnIdx] = append([]byte{txt[i2]},
					field2[columnIdx]...)
			}
		}

		if !strings.HasPrefix(txt, "move ") {
			continue
		}
		n, from, to := parseMove(txt)

		// move field1 1
		for i := 0; i < n; i++ {
			field1[to] = append(field1[to], field1[from][len(field1[from])-1])
			field1[from] = field1[from][:len(field1[from])-1]
		}

		// move field1 2
		field2[to] = append(field2[to],
			field2[from][len(field2[from])-n:]...)
		field2[from] = field2[from][:len(field2[from])-n]
	}

	fmt.Printf("field 1: ")
	for i := 0; i < len(field1); i++ {
		fmt.Printf("%c", field1[i][len(field1[i])-1])
	}
	fmt.Printf("\nfield 2: ")
	for i := 0; i < len(field2); i++ {
		fmt.Printf("%c", field2[i][len(field2[i])-1])
	}
}
