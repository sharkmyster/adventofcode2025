package main

import (
	"sort"
	"strconv"
	"strings"
)

func Exercise05(freshRanges []string, idList []int) int {
	ranges := make([][]int, len(freshRanges))
	for i, freshRange := range(freshRanges) {
		bounds := strings.Split(freshRange, "-")

		bottom, _ := strconv.Atoi(bounds[0])
		top, _ := strconv.Atoi(bounds[1])

		ranges[i] = []int{bottom, top}
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})

	total := 0

	for i := 0; i < len(ranges); i++ {
		if ranges[i][1] > ranges[i+1][0] {

		}
	}
	

	return total
}