package day07

import (
	"fmt"
	"strings"
)

const pipe string = "|"
const dot string = "."
const junction = "^"
const start = "S"


func Exercise07Part1(input []string)  {
	parsedInput := splitStrings(input)
	countSplits := 0

	for i, row := range parsedInput {
		for j, spot := range row {
			if spot == start {
				parsedInput[i+1] [j] = pipe 
			}
			if spot == junction {
				if parsedInput[i-1][j] == pipe {
					countSplits += 1
					parsedInput[i][j-1] = pipe
					parsedInput[i][j+1] = pipe
				}
			}
			if i > 0 && spot == dot && parsedInput[i-1][j] == pipe {
				parsedInput[i][j] = pipe
			}
		}
		fmt.Println(row)
	}

	fmt.Println(countSplits)
}

func Exercise07Part2() {
	
}

func splitStrings(input []string) [][]string {
	newSlice := [][]string{}
	for _, val := range input {
		parts := strings.Split(val, "")
		newSlice = append(newSlice, parts)
	}
	
	return newSlice
}