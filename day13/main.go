package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sort"

	"github.com/olomix/adventofcode2022/internal"
)

func main() {
	var pairs [][2][]any
	var first []any
	for txt := range internal.ReadLines("day13/input.txt") {
		if txt == "" {
			continue
		}

		if first == nil {
			err := json.Unmarshal([]byte(txt), &first)
			internal.Perr(err)
		} else {
			var second []any
			err := json.Unmarshal([]byte(txt), &second)
			internal.Perr(err)
			pairs = append(pairs, [2][]any{first, second})
			first = nil
		}
	}

	rightIdxs := []int{}
	for i := range pairs {
		r := cmp(pairs[i][0], pairs[i][1])
		if r == -1 {
			rightIdxs = append(rightIdxs, i+1)
		}
	}

	sum := 0
	for _, idx := range rightIdxs {
		sum += idx
	}
	fmt.Println(sum)

	msg2 := []any{[]any{float64(2)}}
	msg6 := []any{[]any{float64(6)}}
	flatList := [][]any{msg2, msg6}
	for _, pair := range pairs {
		flatList = append(flatList, pair[0], pair[1])
	}
	sort.Slice(flatList, func(i, j int) bool {
		return cmp(flatList[i], flatList[j]) == -1
	})

	idx2 := -1
	idx6 := -1
	for i, list := range flatList {
		if reflect.DeepEqual(msg2, list) {
			idx2 = i + 1
		}
		if reflect.DeepEqual(msg6, list) {
			idx6 = i + 1
		}
	}

	fmt.Println(idx2 * idx6)
}

func cmp(left, right []any) int {
	if len(left) == 0 && len(right) == 0 {
		return 0
	}

	i := 0
	for {
		// if both list are empty, then they are equal
		if i > len(left)-1 && i > len(right)-1 {
			return 0
		}
		// if left is empty, then it is less
		if i > len(left)-1 && i <= len(right)-1 {
			return -1
		}
		// if right is empty, then left is not less
		if i <= len(left)-1 && i > len(right)-1 {
			return 1
		}

		fL, okL := left[i].(float64)
		fR, okR := right[i].(float64)
		if !okL || !okR {
			var newLeft, newRight []any
			if !okL {
				newLeft = left[i].([]any)
			} else {
				newLeft = []any{left[i]}
			}
			if !okR {
				newRight = right[i].([]any)
			} else {
				newRight = []any{right[i]}
			}
			switch cmp(newLeft, newRight) {
			case -1:
				return -1
			case 1:
				return 1
			}
		} else {
			if fL < fR {
				return -1
			}
			if fL > fR {
				return 1
			}
		}
		i++
	}
}
