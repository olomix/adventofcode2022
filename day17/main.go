package main

import (
	"fmt"
	"time"

	"github.com/olomix/adventofcode2022/internal"
)

type Figure interface {
	Init(field *Field)
	CanGoDown(field *Field) bool
	GoDown()
	GoLeft(field *Field)
	GoRight(field *Field)
	DoRest(field *Field)
}

type FigureHorizontalDash struct {
	// coordinate of the leftmost point
	x, y int
}

func (f *FigureHorizontalDash) Init(field *Field) {
	f.x = 4
	f.y = field.highestY + 4
	field.Grow(f.y + 1)
}

func (f *FigureHorizontalDash) CanGoDown(field *Field) bool {
	if f.y == 0 {
		return false
	}
	mask := byte(0b1111 << (f.x - 3))
	return (field.rows[f.y-1] & mask) == 0
}

func (f *FigureHorizontalDash) GoDown() {
	f.y--
}

func (f *FigureHorizontalDash) GoLeft(field *Field) {
	if f.x == 6 {
		return
	}
	if field.rows[f.y]&(1<<(f.x+1)) != 0 {
		return
	}
	f.x++
}

func (f *FigureHorizontalDash) GoRight(field *Field) {
	if f.x <= 3 {
		return
	}
	if field.rows[f.y]&(1<<(f.x-4)) != 0 {
		return
	}
	f.x--
}

func (f *FigureHorizontalDash) DoRest(field *Field) {
	field.rows[f.y] |= 0b1111 << (f.x - 3)
	if f.y > field.highestY {
		field.highestY = f.y
	}
}

type FigureVerticalDash struct {
	// coordinate of the bottom point
	x, y int
}

func (f *FigureVerticalDash) Init(field *Field) {
	f.x = 4
	f.y = field.highestY + 4
	field.Grow(f.y + 4)
}

func (f *FigureVerticalDash) CanGoDown(field *Field) bool {
	if f.y == 0 {
		return false
	}
	return field.rows[f.y-1]&(0b1<<f.x) == 0
}

func (f *FigureVerticalDash) GoDown() {
	f.y--
}

func (f *FigureVerticalDash) GoLeft(field *Field) {
	if f.x == 6 {
		return
	}
	if field.rows[f.y]&(1<<(f.x+1)) != 0 ||
		field.rows[f.y+1]&(1<<(f.x+1)) != 0 ||
		field.rows[f.y+2]&(1<<(f.x+1)) != 0 ||
		field.rows[f.y+3]&(1<<(f.x+1)) != 0 {
		return
	}
	f.x++
}

func (f *FigureVerticalDash) GoRight(field *Field) {
	if f.x == 0 {
		return
	}
	if field.rows[f.y]&(1<<(f.x-1)) != 0 ||
		field.rows[f.y+1]&(1<<(f.x-1)) != 0 ||
		field.rows[f.y+2]&(1<<(f.x-1)) != 0 ||
		field.rows[f.y+3]&(1<<(f.x-1)) != 0 {
		return
	}
	f.x--
}

func (f *FigureVerticalDash) DoRest(field *Field) {
	for y := f.y; y < f.y+4; y++ {
		field.rows[y] |= 1 << f.x
	}
	if f.y+3 > field.highestY {
		field.highestY = f.y + 3
	}
}

type FigurePlus struct {
	// coordinate of the center of the plus figure
	x, y int
}

func (f *FigurePlus) Init(field *Field) {
	f.x = 3
	f.y = field.highestY + 5
	field.Grow(f.y + 2)
}

func (f *FigurePlus) CanGoDown(field *Field) bool {
	if f.y == 1 {
		return false
	}
	mask2 := byte(0b1 << f.x)
	mask1 := byte(0b101) << (f.x - 1)
	return field.rows[f.y-2]&mask2 == 0 && field.rows[f.y-1]&mask1 == 0
}

func (f *FigurePlus) GoDown() {
	f.y--
}

func (f *FigurePlus) GoLeft(field *Field) {
	if f.x == 5 {
		return
	}
	if field.rows[f.y+1]&(1<<(f.x+1)) != 0 ||
		field.rows[f.y]&(1<<(f.x+2)) != 0 ||
		field.rows[f.y-1]&(1<<(f.x+1)) != 0 {
		return
	}
	f.x++
}

func (f *FigurePlus) GoRight(field *Field) {
	if f.x == 1 {
		return
	}
	if field.rows[f.y+1]&(1<<(f.x-1)) != 0 ||
		field.rows[f.y]&(1<<(f.x-2)) != 0 ||
		field.rows[f.y-1]&(1<<(f.x-1)) != 0 {
		return
	}
	f.x--
}

func (f *FigurePlus) DoRest(field *Field) {
	field.rows[f.y+1] |= 1 << f.x
	field.rows[f.y] |= 0b111 << (f.x - 1)
	field.rows[f.y-1] |= 0b1 << f.x
	if f.y+1 > field.highestY {
		field.highestY = f.y + 1
	}
}

type FigureAngle struct {
	// coordinate of the corner of the angle figure
	x, y int
}

func (f *FigureAngle) Init(field *Field) {
	f.x = 2
	f.y = field.highestY + 4
	field.Grow(f.y + 3)
}

func (f *FigureAngle) CanGoDown(field *Field) bool {
	if f.y == 0 {
		return false
	}
	return field.rows[f.y-1]&(0b111<<f.x) == 0
}

func (f *FigureAngle) GoDown() {
	f.y--
}

func (f *FigureAngle) GoLeft(field *Field) {
	if f.x == 4 {
		return
	}
	if field.rows[f.y]&(1<<(f.x+3)) != 0 ||
		field.rows[f.y+1]&(1<<(f.x+1)) != 0 ||
		field.rows[f.y+2]&(1<<(f.x+1)) != 0 {
		return
	}
	f.x++
}

func (f *FigureAngle) GoRight(field *Field) {
	if f.x == 0 {
		return
	}
	if field.rows[f.y]&(1<<(f.x-1)) != 0 ||
		field.rows[f.y+1]&(1<<(f.x-1)) != 0 ||
		field.rows[f.y+2]&(1<<(f.x-1)) != 0 {
		return
	}
	f.x--
}

func (f *FigureAngle) DoRest(field *Field) {
	field.rows[f.y] |= 0b111 << f.x
	field.rows[f.y+1] |= 0b1 << f.x
	field.rows[f.y+2] |= 0b1 << f.x
	if f.y+2 > field.highestY {
		field.highestY = f.y + 2
	}
}

type FigureSquare struct {
	// coordinate of the bottom left corner of the square figure
	x, y int
}

func (f *FigureSquare) Init(field *Field) {
	f.x = 4
	f.y = field.highestY + 4
	field.Grow(f.y + 2)
}

func (f *FigureSquare) CanGoDown(field *Field) bool {
	if f.y == 0 {
		return false
	}
	return field.rows[f.y-1]&(0b11<<(f.x-1)) == 0
}

func (f *FigureSquare) GoDown() {
	f.y--
}

func (f *FigureSquare) GoLeft(field *Field) {
	if f.x == 6 {
		return
	}
	if field.rows[f.y]&(1<<(f.x+1)) != 0 ||
		field.rows[f.y+1]&(1<<(f.x+1)) != 0 {
		return
	}
	f.x++
}

func (f *FigureSquare) GoRight(field *Field) {
	if f.x == 1 {
		return
	}
	if field.rows[f.y]&(1<<(f.x-2)) != 0 ||
		field.rows[f.y+1]&(1<<(f.x-2)) != 0 {
		return
	}
	f.x--
}

func (f *FigureSquare) DoRest(field *Field) {
	field.rows[f.y] |= 0b11 << (f.x - 1)
	field.rows[f.y+1] |= 0b11 << (f.x - 1)
	if f.y+1 > field.highestY {
		field.highestY = f.y + 1
	}
}

var figures = []Figure{&FigureHorizontalDash{}, &FigurePlus{},
	&FigureAngle{}, &FigureVerticalDash{}, &FigureSquare{}}

type Field struct {
	rows        []byte
	moves       string
	moveIdx     int
	state       State
	figuresRest int
	highestY    int
}

func (f *Field) Step() {
	f.state.Step()
}

// Grow to minimum n rows
func (f *Field) Grow(n int) {
	for len(f.rows) < n {
		f.rows = append(f.rows, 0)
	}
}

type State interface {
	Step()
}

type StateEmpty struct {
	field *Field
}

func (s *StateEmpty) Step() {
	//if s.field.highestY%2 == 0 {
	//	mid := s.field.highestY / 2
	//	if bytes.Equal(s.field.rows[:mid],
	//		s.field.rows[mid:s.field.highestY+1]) {
	//		log.Printf("found equal at %v", s.field.highestY)
	//	}
	//}
	nextFigure := figures[s.field.figuresRest%len(figures)]
	nextFigure.Init(s.field)
	s.field.state = &StateMoving{figure: nextFigure, field: s.field}
}

type StateMoving struct {
	figure Figure
	field  *Field
}

func (s *StateMoving) Step() {
	if s.field.moves[s.field.moveIdx%len(s.field.moves)] == '>' {
		s.figure.GoRight(s.field)
	} else if s.field.moves[s.field.moveIdx%len(s.field.moves)] == '<' {
		s.figure.GoLeft(s.field)
	} else {
		panic("unknown move")
	}
	s.field.moveIdx++

	if s.figure.CanGoDown(s.field) {
		s.figure.GoDown()
	} else {
		s.figure.DoRest(s.field)
		s.field.figuresRest++
		s.field.state = &StateEmpty{field: s.field}
	}
}

func main() {
	var moves string
	for txt := range internal.ReadLines("day17/input.txt") {
		if txt == "" {
			continue
		}
		moves = txt
	}

	state := &StateEmpty{}
	field := &Field{
		state:    state,
		moves:    moves,
		rows:     make([]byte, 1_000_000),
		highestY: -1,
	}
	state.field = field
	// 1000000 > 1514288
	start := time.Now()
	for field.figuresRest < 10_000_000 {
		field.Step()
	}
	//printField(field)
	fmt.Println(field.highestY + 1)
	fmt.Println("done in", time.Since(start))

	//start := time.Now()
	//t := 0
	//for i := 0; i < 1_000_000_000_000; i++ {
	//	if i%1_000_000_000 == 0 {
	//		fmt.Println(1, t, time.Since(start))
	//	}
	//	t += i * 2
	//}
	//fmt.Println(2, t, time.Since(start))

	return
	// 1.000.000.000.000
	for field.figuresRest < 1000000000000 {
		if field.figuresRest%100000000 == 0 {
			fmt.Printf("cnt=%v\n", field.figuresRest)
		}
		field.Step()
	}
	fmt.Println(field.highestY + 1)
}

func printField(field *Field) {
	for row := field.highestY; row >= 0; row-- {
		for col := 6; col >= 0; col-- {
			if field.rows[row]&(1<<uint(col)) != 0 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
