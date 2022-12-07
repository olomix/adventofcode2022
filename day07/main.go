package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"

	"github.com/olomix/adventofcode2022/internal"
)

type Dir struct {
	files    map[string]int
	parent   *Dir
	children map[string]*Dir
}

func (d *Dir) Size() int {
	size := 0
	for _, s := range d.files {
		size += s
	}
	for _, c := range d.children {
		size += c.Size()
	}
	return size
}

func (d *Dir) AllDirs() []*Dir {
	all := []*Dir{d}
	for _, c := range d.children {
		all = append(all, c.AllDirs()...)
	}
	return all
}

var cmdCD = regexp.MustCompile(`\$ cd (.+)`)
var cmdLS = regexp.MustCompile(`\$ ls`)
var outDIR = regexp.MustCompile(`dir (.*)`)
var outFILE = regexp.MustCompile(`(\d+) (.+)`)

func main() {
	i := 0

	root := &Dir{
		parent:   nil,
		files:    make(map[string]int),
		children: make(map[string]*Dir),
	}
	current := root
	for txt := range internal.ReadLines("day07/input.txt") {
		i += 1
		if txt == "" {
			continue
		}

		m := cmdCD.FindStringSubmatch(txt)
		if m != nil {
			dirName := m[1]
			if dirName == ".." {
				current = current.parent
				if current == nil {
					current = root
				}
				continue
			}

			if dirName == "/" {
				current = root
				continue
			}

			if _, ok := current.children[dirName]; !ok {
				current.children[dirName] = &Dir{
					files:    make(map[string]int),
					parent:   current,
					children: make(map[string]*Dir),
				}
			}
			current = current.children[dirName]
			continue
		}

		m = cmdLS.FindStringSubmatch(txt)
		if m != nil {
			continue
		}

		m = outDIR.FindStringSubmatch(txt)
		if m != nil {
			dirName := m[1]
			if _, ok := current.children[dirName]; !ok {
				current.children[dirName] = &Dir{
					files:    make(map[string]int),
					parent:   current,
					children: make(map[string]*Dir),
				}
			}
			continue
		}

		m = outFILE.FindStringSubmatch(txt)
		if m != nil {
			fileName := m[2]
			fileSize, err := strconv.ParseInt(m[1], 10, 64)
			internal.Perr(err)
			current.files[fileName] = int(fileSize)
			continue
		}

		panic(fmt.Sprintf("invalid line %d: %s", i, txt))
	}

	allDirs := root.AllDirs()
	allSums := []int{}
	sum := 0
	for _, d := range allDirs {
		s := d.Size()
		allSums = append(allSums, s)
		if s <= 100000 {
			sum += s
		}
	}
	// 1423358
	fmt.Println(sum)

	totalSize := 70000000
	requiredFreeSize := 30000000
	rootSize := root.Size()
	currentFreeSize := totalSize - rootSize
	requiredSize := requiredFreeSize - currentFreeSize
	sort.Ints(allSums)
	for _, s := range allSums {
		if s >= requiredSize {
			fmt.Println(s) // 545729
			break
		}
	}
}
