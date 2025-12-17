package day05

import (
	"sort"
	"strconv"
	"strings"
)

// correct answer 366181852921027

func Part1(freshRanges []string, idList []int) int {

	ranges := convertToRangeSlices(freshRanges)

	freshIngredients := 0

	for _, id := range idList {
		for _, idRange := range ranges {
			if id >= idRange[0] && id <= idRange[1] {
				freshIngredients += 1
				break
			}
		}
	}

	return freshIngredients
}

func Part2(freshRanges []string, idList []int) int {

	ranges := convertToRangeSlices(freshRanges)

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})

	newRanges := calculateRanges(ranges)

	total := 0

	for _, r := range newRanges {
		total += (r[1] - r[0]) + 1
	}

	return total

}

func calculateRanges(ranges [][]int) [][]int {
	if len(ranges) == 0 {
		return [][]int{}
	}

	merged := [][]int{}
	current := ranges[0]

	for i := 1; i < len(ranges); i++ {
		next := ranges[i]
		// Check if ranges overlap or are adjacent
		if current[1] >= next[0]-1 {
			// The next range might be smaller than the current range
			if next[1] > current[1] {
				current[1] = next[1]
			}
		} else {
			// No overlap, add current and move to next
			merged = append(merged, []int{current[0], current[1]})
			current = next
		}
	}
	
	// Add the last range
	merged = append(merged, []int{current[0], current[1]})

	return merged
}

func convertToRangeSlices(strSlice []string) [][]int {
	ranges := make([][]int, len(strSlice))

	for i, freshRange := range strSlice {
		bounds := strings.Split(freshRange, "-")

		bottom, _ := strconv.Atoi(bounds[0])
		top, _ := strconv.Atoi(bounds[1])

		ranges[i] = []int{bottom, top}
	}

	return ranges
}
