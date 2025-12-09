package main

import (
	"fmt"
	"strconv"
)

func Exercise06(input []string) int {
	numbers := parseInput(input[:len(input)-1])
	operators := parseOperators(input[len(input)-1])

	numbers = transpose(numbers)

	grandTotal := 0

	for i, row := range numbers {
		operator := operators[i]

		// fmt.Printf("%v %v\n", operator, row)

		if operator == "*" {
			fmt.Printf("%v %v %v\n", operator, row, multiply(row))
			grandTotal += multiply(row)
		}

		if operator == "+" {
			fmt.Printf("%v %v %v\n", operator, row, add(row))
			grandTotal += add(row)
		}
	}

	fmt.Println(numbers)
	fmt.Println(operators)

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
	if len(nums) == 1 {
		return nums[0]
	}
	return nums[0] * multiply(nums[1:])
}

func add(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}
