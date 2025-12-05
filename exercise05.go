package main

import (
	"fmt"
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

	for i := 0; i < len(ranges) - 1; i++ {
		start := i
		for j := start; j < len(ranges)-1; j++ {
			if (ranges[i][1] > ranges[i+1][0]) || ((ranges[i][1] + 1) == ranges[i+1][0]) {
				ranges[i][1] = ranges[i+1][1]
			}
		}
		
	}

	for i := 0; i < 3; i++ {
		fmt.Println(ranges[i])
	}
	

	return total
}