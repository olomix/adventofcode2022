package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/olomix/adventofcode2022/internal"
)

type coord struct {
	x, y, z int
}

func try[T any](fn func() (T, error)) T {
	res, err := fn()
	if err != nil {
		panic(err)
	}
	return res
}

func parseCoord(s string) coord {
	coordParts := strings.Split(s, ",")
	return coord{
		x: try(func() (int, error) { return strconv.Atoi(coordParts[0]) }),
		y: try(func() (int, error) { return strconv.Atoi(coordParts[1]) }),
		z: try(func() (int, error) { return strconv.Atoi(coordParts[2]) }),
	}
}

func maxDimension(cs []coord) int {
	m := 0
	for _, c := range cs {
		if c.x > m {
			m = c.x
		}
		if c.y > m {
			m = c.y
		}
		if c.z > m {
			m = c.z
		}
	}
	return m
}

type bitField struct {
	bits      []uint64
	dimension int
}

func newBitField(dimension int) *bitField {
	size := dimension * dimension * dimension
	return &bitField{bits: make([]uint64, (size+63)/64), dimension: dimension}
}

func (bf *bitField) isSet(x, y, z int) bool {
	if x < 0 || y < 0 || z < 0 ||
		x >= bf.dimension || y >= bf.dimension || z >= bf.dimension {
		return false
	}
	i := x*bf.dimension*bf.dimension + y*bf.dimension + z
	return bf.bits[i/64]&(1<<(i%64)) != 0
}

func (bf *bitField) setBit(x, y, z int) {
	i := x*bf.dimension*bf.dimension + y*bf.dimension + z
	bf.bits[i/64] |= 1 << (i % 64)
}

func traverseOpenAirField(openAirField, lavaField *bitField,
	x, y, z int) {
	if lavaField.isSet(x, y, z) {
		panic(1)
	}
	openAirField.setBit(x, y, z)
	if x < openAirField.dimension-1 && !openAirField.isSet(x+1, y, z) &&
		!lavaField.isSet(x+1, y, z) {
		traverseOpenAirField(openAirField, lavaField, x+1, y, z)
	}
	if x > 0 && !openAirField.isSet(x-1, y, z) &&
		!lavaField.isSet(x-1, y, z) {
		traverseOpenAirField(openAirField, lavaField, x-1, y, z)
	}
	if y < openAirField.dimension-1 && !openAirField.isSet(x, y+1, z) &&
		!lavaField.isSet(x, y+1, z) {
		traverseOpenAirField(openAirField, lavaField, x, y+1, z)
	}
	if y > 0 && !openAirField.isSet(x, y-1, z) &&
		!lavaField.isSet(x, y-1, z) {
		traverseOpenAirField(openAirField, lavaField, x, y-1, z)
	}
	if z < openAirField.dimension-1 && !openAirField.isSet(x, y, z+1) &&
		!lavaField.isSet(x, y, z+1) {
		traverseOpenAirField(openAirField, lavaField, x, y, z+1)
	}
	if z > 0 && !openAirField.isSet(x, y, z-1) &&
		!lavaField.isSet(x, y, z-1) {
		traverseOpenAirField(openAirField, lavaField, x, y, z-1)
	}
}

func openAirField(lavaField *bitField) *bitField {
	openField := newBitField(lavaField.dimension)
	for x := 0; x < lavaField.dimension; x++ {
		for y := 0; y < lavaField.dimension; y++ {
			for z := 0; z < lavaField.dimension; z++ {
				if x != 0 && x != lavaField.dimension-1 &&
					y != 0 && y != lavaField.dimension-1 &&
					z != 0 && z != lavaField.dimension-1 {
					continue
				}
				if lavaField.isSet(x, y, z) || openField.isSet(x, y, z) {
					continue
				}
				traverseOpenAirField(openField, lavaField, x, y, z)
			}
		}
	}
	return openField
}

func main() {
	var coords []coord
	for txt := range internal.ReadLines("day18/input.txt") {
		if txt == "" {
			continue
		}

		coords = append(coords, parseCoord(txt))
	}

	dimension := maxDimension(coords) + 1
	field := newBitField(dimension)
	surfaceArea := 0
	for _, c := range coords {
		incArea := 6
		if field.isSet(c.x-1, c.y, c.z) {
			incArea--
			surfaceArea--
		}
		if field.isSet(c.x+1, c.y, c.z) {
			incArea--
			surfaceArea--
		}
		if field.isSet(c.x, c.y-1, c.z) {
			incArea--
			surfaceArea--
		}
		if field.isSet(c.x, c.y+1, c.z) {
			incArea--
			surfaceArea--
		}
		if field.isSet(c.x, c.y, c.z-1) {
			incArea--
			surfaceArea--
		}
		if field.isSet(c.x, c.y, c.z+1) {
			incArea--
			surfaceArea--
		}

		field.setBit(c.x, c.y, c.z)
		surfaceArea += incArea
	}

	fmt.Println(surfaceArea)

	oaField := openAirField(field)

	for x := 0; x < oaField.dimension; x++ {
		for y := 0; y < oaField.dimension; y++ {
			for z := 0; z < oaField.dimension; z++ {
				if !field.isSet(x, y, z) && !oaField.isSet(x, y, z) {
					if field.isSet(x-1, y, z) {
						surfaceArea--
					}
					if field.isSet(x+1, y, z) {
						surfaceArea--
					}
					if field.isSet(x, y-1, z) {
						surfaceArea--
					}
					if field.isSet(x, y+1, z) {
						surfaceArea--
					}
					if field.isSet(x, y, z-1) {
						surfaceArea--
					}
					if field.isSet(x, y, z+1) {
						surfaceArea--
					}
				}
			}
		}
	}
	fmt.Println(surfaceArea)
}
