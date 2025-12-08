package main

import "fmt"

func Exercise03(input []string) {
	intSlices := make([][]int, len(input))

	for i, row := range input {
		slice := toIntSlice(row)
		intSlices[i] = slice
	}

	sum := 0
	for _, slice := range intSlices {
		sum += findOptimalJoltageDigits(slice, 12)
	}

	fmt.Println(sum)
}

func findOptimalJoltageDigits(slice []int, numDigits int) int {
	digitPositions := make([]int, numDigits)

	for i := range numDigits {
		digitPositions[i] = len(slice) + i - numDigits
	}

	for j := range digitPositions {
		start := digitPositions[j]

		limit := -1
		if j > 0 {
			limit = digitPositions[j-1]
		}

		for i := start - 1; i > limit; i-- {
			if slice[i] >= slice[start] {
				start = i
			}
		}

		digitPositions[j] = start
	}

	num := 0
	for i := range digitPositions {
		num += slice[digitPositions[i]] * powInt(10, numDigits-1-i)
	}

	return num

}

func toIntSlice(s string) []int {
	ints := make([]int, len(s))

	for i := 0; i < len(s); i++ {
		ints[i] = int(s[i] - '0')
	}
	return ints
}

func powInt(a, b int) int {
	result := 1
	for i := 0; i < b; i++ {
		result *= a
	}
	return result
}
