package day04

import (
	"strings"
)

func splitRow(row string) []string {
	return strings.Split(row, "")
}

type Point struct {
	X int
	Y int
}

var north = Point{X: 0, Y: -1}
var northwest = Point{X: -1, Y: -1}
var northeast = Point{X: 1, Y: -1}
var west = Point{X: -1, Y: 0}
var east = Point{X: 1, Y: 0}
var southwest = Point{X: -1, Y: 1}
var southeast = Point{X: 1, Y: 1}
var south = Point{X: 0, Y: 1}

const roll string = "@"
const blank string = "."
const maxRolls int = 4


func Part1(input []string) int {
	rows := [][]string{}
	
	for _, row := range input {
		rows = append(rows, splitRow(row))
	}

	_, numRemovedRolls := removeRolls(rows)

	return numRemovedRolls
}

func Part2(input []string) int {
	rows := [][]string{}
	for _, row := range input {
		rows = append(rows, splitRow(row))
	}

	totalRemoved := 0
	for true {
		newRolls, numRemovedRolls := removeRolls(rows)
		if numRemovedRolls == 0 {
			break
		}
		totalRemoved += numRemovedRolls
		rows = Copy2D(newRolls)
	}

	return totalRemoved
}

func removeRolls(rows [][]string) ([][]string, int) {
	totalRemoved := 0

	newRows := Copy2D(rows)

	for y, row := range rows {
		for x, item := range row {
			if item == roll {
				adjacentRolls := countAdjacentRolls(rows, x, y)

				if adjacentRolls < maxRolls {
					newRows[y][x] = blank
					totalRemoved += 1
				}
			}
		}
	}
	return newRows, totalRemoved
}

func countAdjacentRolls(locationsMap [][]string, x int, y int) int {
	score := 0

	points := []Point{
		north,
		northwest,
		northeast,
		west,
		east,
		southwest,
		southeast,
		south,
	}

	width := len(locationsMap[0])
	height := len(locationsMap)

	for _, point := range points {
		adjacentX := x + point.X
		adjacentY := y + point.Y

		// fmt.Printf("{%d, %d} \n", adjacentY , adjacentX)

		if adjacentX < 0 || adjacentX >= width || adjacentY < 0 || adjacentY >= height {
			continue
		}
		if locationsMap[adjacentY][adjacentX] == "@" {
			score += 1
		}
	}

	return score
}

func Copy2D(src [][]string) [][]string {
	dst := make([][]string, len(src))
	for i := range src {
		dst[i] = make([]string, len(src[i]))
		copy(dst[i], src[i])
	}
	return dst
}
