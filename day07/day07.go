package day07

import (
	"adventofcode2025/utils"
	"slices"
)

const pipe string = "|"
const dot string = "."
const junction = "^"
const startSymbol = "S"


func Part1(input []string) int {
	grid := utils.SplitStrings(input)
	rows := len(grid)
	cols := len(grid[0])
	countSplits := 0

	for j := range cols {
		if grid[0][j] == startSymbol && 1 < rows {
			grid[1][j] = pipe
		}
	}

	for i := 1; i < rows; i++ {
		for j := range cols {
			spot := grid[i][j]

			// Handle splitting
			if spot == junction && grid[i-1][j] == pipe {
				countSplits++
				if j > 0 {
					grid[i][j-1] = pipe
				}
				if j < cols-1 {
					grid[i][j+1] = pipe
				}
			}

			// Propagate pipes downward
			if spot == dot && grid[i-1][j] == pipe {
				grid[i][j] = pipe
			}
		}
	}

	return countSplits
}

func Part2(input []string) int {
	parsedInput := utils.SplitStrings(input)

	startColumn := slices.Index(parsedInput[0], startSymbol)

	countTachyon(parsedInput, 1, startColumn)

	return countTachyon(parsedInput, 1, startColumn)
}

var cache = make(map[int]map[int]int)

func countTachyon(parsedInput [][]string, row int, col int) int {
	if row == len(parsedInput)-1 {
		return 1
	}

	if col < 0 || col >= len(parsedInput[0]) {
		return 0
	}

	if cache[row] == nil {
		cache[row] = make(map[int]int)
	}

	if cached, exists := cache[row][col]; exists {
		return cached
	}

	var result int
	if parsedInput[row][col] == junction {
		result = countTachyon(parsedInput, row, col-1) + countTachyon(parsedInput, row, col+1)
	} else {
		result = countTachyon(parsedInput, row+1, col)
	}

	cache[row][col] = result
	return result
}

