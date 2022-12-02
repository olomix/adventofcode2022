package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func perr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	f, err := os.Open("day1/input.txt")
	perr(err)
	defer func() { perr(f.Close()) }()

	scanner := bufio.NewScanner(f)
	perElf := make([]int, 0)
	sum := 0
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			perElf = append(perElf, sum)
			sum = 0
			continue
		}

		n, err := strconv.Atoi(txt)
		perr(err)
		sum += n
	}
	perr(scanner.Err())
	if sum != 0 {
		perElf = append(perElf, sum)
	}
	sort.Ints(perElf)
	fmt.Printf("max: %d\n", perElf[len(perElf)-1])
	fmt.Printf("top 3 sum: %d\n",
		perElf[len(perElf)-1]+perElf[len(perElf)-2]+perElf[len(perElf)-3])
}
