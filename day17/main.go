package main

import (
	"fmt"

	"github.com/olomix/adventofcode2022/internal"
)

type figureKind uint8

const (
	HorizontalDash figureKind = iota
	VerticalDash
	Plus
	Angle
	Square
)

var figuresOrder = [...]figureKind{HorizontalDash, VerticalDash, Plus, Angle,
	Square}

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
	f.x = 2
	f.y = field.MaxTopOccupiedCoord() + 4
	field.Grow(f.y + 1)
}

func (f *FigureHorizontalDash) CanGoDown(field *Field) bool {
	if f.y == 0 {
		return false
	}
	for i := f.x; i <= f.x+3; i++ {
		if field.rows[f.y-1][i] {
			return false
		}
	}
	return true
}

func (f *FigureHorizontalDash) GoDown() {
	f.y--
}

func (f *FigureHorizontalDash) GoLeft(field *Field) {
	if f.x == 0 {
		return
	}
	if !field.rows[f.y][f.x-1] {
		f.x--
	}
}

func (f *FigureHorizontalDash) GoRight(field *Field) {
	if f.x >= 3 {
		return
	}
	if !field.rows[f.y][f.x+4] {
		f.x++
	}
}

func (f *FigureHorizontalDash) DoRest(field *Field) {
	for i := f.x; i < f.x+4; i++ {
		field.rows[f.y][i] = true
	}
}

type FigureVerticalDash struct {
	// coordinate of the bottom point
	x, y int
}

func (f *FigureVerticalDash) Init(field *Field) {
	f.x = 2
	f.y = field.MaxTopOccupiedCoord() + 4
	field.Grow(f.y + 4)
}

func (f *FigureVerticalDash) CanGoDown(field *Field) bool {
	if f.y == 0 {
		return false
	}
	return !field.rows[f.y-1][f.x]
}

func (f *FigureVerticalDash) GoDown() {
	f.y--
}

func (f *FigureVerticalDash) GoLeft(field *Field) {
	if f.x == 0 {
		return
	}
	for y := f.y; y < f.y+4; y++ {
		if field.rows[y][f.x-1] {
			return
		}
	}
	f.x--
}

func (f *FigureVerticalDash) GoRight(field *Field) {
	if f.x >= 6 {
		return
	}
	for y := f.y; y < f.y+4; y++ {
		if field.rows[y][f.x+1] {
			return
		}
	}
	f.x++
}

func (f *FigureVerticalDash) DoRest(field *Field) {
	for y := f.y; y < f.y+4; y++ {
		field.rows[y][f.x] = true
	}
}

type FigurePlus struct {
	// coordinate of the center of the plus figure
	x, y int
}

func (f *FigurePlus) Init(field *Field) {
	f.x = 3
	f.y = field.MaxTopOccupiedCoord() + 5
	field.Grow(f.y + 2)
}

func (f *FigurePlus) CanGoDown(field *Field) bool {
	if f.y == 1 {
		return false
	}
	if field.rows[f.y-2][f.x] || field.rows[f.y-1][f.x-1] || field.rows[f.y-1][f.x+1] {
		return false
	}
	return true
}

func (f *FigurePlus) GoDown() {
	f.y--
}

func (f *FigurePlus) GoLeft(field *Field) {
	if f.x == 1 {
		return
	}

	if !field.rows[f.y][f.x-2] &&
		!field.rows[f.y-1][f.x-1] &&
		!field.rows[f.y+1][f.x-1] {
		f.x--
	}
}

func (f *FigurePlus) GoRight(field *Field) {
	if f.x >= 5 {
		return
	}
	if !field.rows[f.y][f.x+2] &&
		!field.rows[f.y-1][f.x+1] &&
		!field.rows[f.y+1][f.x+1] {
		f.x++
	}
}

func (f *FigurePlus) DoRest(field *Field) {
	field.rows[f.y][f.x-1] = true
	field.rows[f.y][f.x] = true
	field.rows[f.y][f.x+1] = true
	field.rows[f.y-1][f.x] = true
	field.rows[f.y+1][f.x] = true
}

type FigureAngle struct {
	// coordinate of the corner of the angle figure
	x, y int
}

func (f *FigureAngle) Init(field *Field) {
	f.x = 4
	f.y = field.MaxTopOccupiedCoord() + 4
	field.Grow(f.y + 3)
}

func (f *FigureAngle) CanGoDown(field *Field) bool {
	if f.y == 0 {
		return false
	}
	for x := f.x - 2; x <= f.x; x++ {
		if field.rows[f.y-1][x] {
			return false
		}
	}
	return true
}

func (f *FigureAngle) GoDown() {
	f.y--
}

func (f *FigureAngle) GoLeft(field *Field) {
	if f.x == 2 {
		return
	}

	if !field.rows[f.y][f.x-3] &&
		!field.rows[f.y+1][f.x-1] &&
		!field.rows[f.y+2][f.x-1] {
		f.x--
	}
}

func (f *FigureAngle) GoRight(field *Field) {
	if f.x >= 6 {
		return
	}
	if !field.rows[f.y][f.x+1] &&
		!field.rows[f.y+1][f.x+1] &&
		!field.rows[f.y+2][f.x+1] {
		f.x++
	}
}

func (f *FigureAngle) DoRest(field *Field) {
	field.rows[f.y][f.x-2] = true
	field.rows[f.y][f.x-1] = true
	field.rows[f.y][f.x] = true
	field.rows[f.y+1][f.x] = true
	field.rows[f.y+2][f.x] = true
}

type FigureSquare struct {
	// coordinate of the bottom left corner of the square figure
	x, y int
}

func (f *FigureSquare) Init(field *Field) {
	f.x = 2
	f.y = field.MaxTopOccupiedCoord() + 4
	field.Grow(f.y + 2)
	//fmt.Printf("init square at %d, %d\n", f.x, f.y)
}

func (f *FigureSquare) CanGoDown(field *Field) bool {
	if f.y == 0 {
		return false
	}
	return !field.rows[f.y-1][f.x] && !field.rows[f.y-1][f.x+1]
}

func (f *FigureSquare) GoDown() {
	f.y--
}

func (f *FigureSquare) GoLeft(field *Field) {
	if f.x == 0 {
		//fmt.Println("Square: can't go left 1")
		return
	}

	if !field.rows[f.y][f.x-1] && !field.rows[f.y+1][f.x-1] {
		//fmt.Println("Square: goes left")
		f.x--
	} else {
		//fmt.Println("Square: can't go left 2")
	}
}

func (f *FigureSquare) GoRight(field *Field) {
	if f.x >= 5 {
		//fmt.Println("Square: can't go right 1")
		return
	}
	if !field.rows[f.y][f.x+2] && !field.rows[f.y+1][f.x+2] {
		//fmt.Println("Square: goes right")
		f.x++
	} else {
		//fmt.Println("Square: can't go right 2")
	}
}

func (f *FigureSquare) DoRest(field *Field) {
	field.rows[f.y][f.x] = true
	field.rows[f.y][f.x+1] = true
	field.rows[f.y+1][f.x] = true
	field.rows[f.y+1][f.x+1] = true
}

var figures = []Figure{&FigureHorizontalDash{}, &FigurePlus{},
	&FigureAngle{}, &FigureVerticalDash{}, &FigureSquare{}}

type Field struct {
	rows        [][7]bool
	moves       string
	moveIdx     int
	state       State
	figuresRest int
}

func (f *Field) MaxTopOccupiedCoord() int {
	for i := len(f.rows) - 1; i >= 0; i-- {
		for j := 0; j < len(f.rows[i]); j++ {
			if f.rows[i][j] {
				return i
			}
		}
	}
	return -1
}

func (f *Field) Step() {
	f.state.Step()
}

// Grow to minimum n rows
func (f *Field) Grow(n int) {
	for len(f.rows) < n {
		f.rows = append(f.rows, [7]bool{})
	}
}

type State interface {
	Step()
}

type StateEmpty struct {
	field *Field
}

func (s *StateEmpty) Step() {
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
	for txt := range internal.ReadLines("day17/input2.txt") {
		if txt == "" {
			continue
		}
		moves = txt
	}

	state := &StateEmpty{}
	field := &Field{state: state, moves: moves}
	state.field = field
	for field.figuresRest < 2022 {
		field.Step()
	}
	printField(field)
	return
	fmt.Println(field.MaxTopOccupiedCoord() + 1)
	for field.figuresRest < 1000000000000 {
		if field.figuresRest%100000000 == 0 {
			fmt.Printf("cnt=%v\n", field.figuresRest)
		}
		field.Step()
	}
	fmt.Println(field.MaxTopOccupiedCoord() + 1)
}

func printField(field *Field) {
	for row := len(field.rows) - 1; row >= 0; row-- {
		for col := 0; col < len(field.rows[row]); col++ {
			if field.rows[row][col] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
