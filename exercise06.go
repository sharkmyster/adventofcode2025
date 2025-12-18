package main

import (
	"strconv"
	"strings"
)

// Part 01 - 6725216329103
func Exercise06(input []string) int {
	numbers := parseInput(input[:len(input)-1])
	operators := parseOperators(input[len(input)-1])

	numbers = transpose(numbers)

	grandTotal := 0

	for i, row := range numbers {
		operator := operators[i]

		if operator == "*" {
			grandTotal += multiply(row)
		}

		if operator == "+" {

			grandTotal += add(row)
		}
	}

	return grandTotal
}

func Exercise06Part2(input []string) int {
	numbers := parseCephalopodNumbers(input[:len(input)-1])
	operators := parseOperators(input[len(input)-1])
	lenOperators := len(operators)

	grandTotal := 0
	for i, row := range numbers {
		operator := operators[lenOperators-1-i]

		if operator == "*" {
			grandTotal += multiply(row)
		}

		if operator == "+" {

			grandTotal += add(row)
		}
	}

	return grandTotal
}

func parseInput(input []string) [][]int {
	parsedNumbers := [][]int{}
	for i := 0; i < len(input); i++ {
		numbers := parseNumbers(input[i])
		parsedNumbers = append(parsedNumbers, numbers)
	}

	return parsedNumbers
}

func parseNumbers(input string) []int {
	numStr := ""
	nums := []int{}

	for i := 0; i < len(input); i++ {
		if input[i] == ' ' && numStr == "" {
			continue
		}
		if input[i] == ' ' && numStr != "" {
			num, _ := strconv.Atoi(numStr)
			nums = append(nums, num)
			numStr = ""
			continue
		}
		numStr += string(input[i])
	}

	if numStr != "" {
		num, _ := strconv.Atoi(numStr)
		nums = append(nums, num)
	}

	return nums
}

func parseOperators(input string) []string {
	operators := []string{}
	for i := 0; i < len(input); i++ {
		if input[i] == ' ' {
			continue
		}
		operators = append(operators, string(input[i]))
	}
	return operators
}

func transpose(matrix [][]int) [][]int {
	if len(matrix) == 0 {
		return [][]int{}
	}

	transposed := make([][]int, len(matrix[0]))
	for i := range transposed {
		transposed[i] = make([]int, len(matrix))
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			transposed[j][i] = matrix[i][j]
		}
	}

	return transposed
}

func multiply(nums []int) int {
	total := 1
	for _, n := range nums {
		total *= n
	}
	return total
}

func add(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func parseCephalopodNumbers(input []string) [][]int {
	height := len(input)
	width := len(input[0])

	numbersSlice := [][]int{}
	setOperands := []int{}
	for i := width - 1; i >= 0; i-- {
		numStr := ""
		for j := 0; j < height; j++ {
			numStr += string(input[j][i])
		}

		trimmed := strings.Trim(numStr, " ")

		if trimmed == "" {
			numbersSlice = append(numbersSlice, setOperands)
			setOperands = []int{}
			continue
		}

		num, _ := strconv.Atoi(trimmed)
		setOperands = append(setOperands, num)
	}

	numbersSlice = append(numbersSlice, setOperands)

	return numbersSlice
}
