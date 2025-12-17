package day02

import (
	"strconv"
	"strings"
)

type Pair struct {
	RangeStart int
	RangeEnd   int
}

func Part1(input string) int {
	idPairs := getIdPairs(input)

	sum := calculateTotal(idPairs, isInvalidPart1)

	return sum
}

func Part2(input string) int {
	idPairs := getIdPairs(input)

	sum := calculateTotal(idPairs, isInvalidPart2)

	return sum
}

func calculateTotal(pairs []Pair, checkFn func(int) bool) int {
	sum := 0
	for _, pair := range pairs {
		for i := pair.RangeStart; i <= pair.RangeEnd; i++ {
			if checkFn(i) {
				sum += i
			}
		}
	}
	return sum
}

func getIdPairs(input string) []Pair {
	idStrings := strings.Split(input, ",")
	idPairs := make([]Pair, 0, len(idStrings))
	
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
	return idPairs
}

func isInvalidPart1(id int) bool{
	s := strconv.Itoa(id)
	idLength := len(s)
	index := idLength / 2

	return s[:index] == s[index:]
}

func isInvalidPart2(id int) bool {
	s := strconv.Itoa(id)
	idLength := len(s)
	index := idLength / 2

	for i := index; i > 0; i-- {
		if idLength%i != 0 {
			continue
		}
		if isRepeated(s, i) {
			return true
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
