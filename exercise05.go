package main

import (
	"sort"
	"strconv"
	"strings"
)

// correct answer 366181852921027

func Exercise05(freshRanges []string, idList []int) int {

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
	newRanges := [][]int{}
	penultimate := len(ranges) - 1

	for i := 0; i < penultimate; i++ {
		currentRange := ranges[i]
		nextRange := ranges[i+1]
		if currentRange[1] >= nextRange[0] || (currentRange[1]+1) == nextRange[0] {
			nextRange[0] = currentRange[0]
			if currentRange[1] > nextRange[1] {
				nextRange[1] = currentRange[1]
			} 
			currentRange[0] = 0
			currentRange[1] = 0
		}
	}

	for _, r := range ranges {
		if r[0] != 0 && r[1] != 0 {
			newRanges = append(newRanges, r)
		}
	}

	return newRanges
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
