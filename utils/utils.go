package utils

import "strings"

func SplitRow(row string) []string {
	return strings.Split(row, "")
}

func SplitStrings(input []string) [][]string {
	newSlice := [][]string{}
	for _, val := range input {
		parts := strings.Split(val, "")
		newSlice = append(newSlice, parts)
	}
	
	return newSlice
}