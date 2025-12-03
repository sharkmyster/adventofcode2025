package main

import "strconv"

const (
	L    = "L"
	R    = "R"
	PIPS = 100
)

func Exercise01(input []string) int {
	position := 50
	password := 0
	ticks := 0

	for _, val := range input {

		letter, number := split(val)

		if letter == L {
			position, ticks = dec(position, number)
		}
		if letter == R {
			position, ticks = inc(position, number)
		}

		password += ticks
	}

	return password
}

func split(s string) (string, int) {
	letter := s[:1]
	numberString := s[1:]

	num, err := strconv.Atoi(numberString)
	if err != nil {
		panic("Not a number")
	}
	return letter, num
}

func dec(position int, num int) (int, int) {
	revs := num / PIPS
	rem := num % PIPS
	diff := position - rem

	if diff > 0 {
		return diff, revs
	}

	if position != 0 {
		revs += 1
	}

	return (diff + PIPS) % PIPS, revs
}

func inc(position int, num int) (int, int) {
	revs := num / PIPS
	rem := num % PIPS
	sum := position + rem

	if sum < PIPS {
		return sum, revs
	}
	return sum - PIPS, revs + 1
}
