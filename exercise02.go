package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Pair struct {
	RangeStart int
	RangeEnd   int
}

func Exercise02(TestInput02 string) {
	idStrings := strings.Split(TestInput02, ",")
	idPairs := make([]Pair, 0, len(idStrings))
	invalidIds := make([]int, 0, 1_000_000)

	for _, idString := range idStrings {
		parts := strings.Split(idString, "-")

		rangeStart, errA := strconv.Atoi(parts[0])

		if errA != nil {
			panic("Error parsing number first")
		}

		rangeEnd, errB := strconv.Atoi(parts[1])

		if errB != nil {
			panic("Error parsing number last")
		}
		idPairs = append(idPairs, Pair{RangeStart: (rangeStart), RangeEnd: (rangeEnd)})
	}

	for _, pair := range idPairs {

		for i := pair.RangeStart; i <= pair.RangeEnd; i++ {
			if isValid(i) {
				invalidIds = append(invalidIds, i)
			}
		}
	}

	sum := 0
	for _, value := range(invalidIds) {
		sum += value
	}

	fmt.Println(sum)
}

func isValid(id int) bool {
	s := strconv.Itoa(id)
	idLength := len(s)
	index := idLength/2

	for i := index; i > 0; i-- {
		if idLength % i != 0 {
			continue
		}
		if isRepeated(s, i) {
			return true;
		}
	}

	return false
}

func isRepeated(s string, seq int) bool {
	splits := len(s) / seq
	start := s[:seq]
	rest := s[seq:]

	for i := 0; i < splits-1; i++ {
		next := rest[:seq]
		rest = rest[seq:]

		if start != next {
			return false
		}
	}

	return true
}
