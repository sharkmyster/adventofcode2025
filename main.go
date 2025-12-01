package main

import (
	"fmt"
	"strconv"
)

const (
	L = "L"
	R = "R"
	PIPS = 100
)

// var input = []string{
// 	"L68",
// 	"L30",
// 	"R48",
// 	"L5",
// 	"R60",
// 	"L55",
// 	"L1",
// 	"L99",
// 	"R14",
// 	"L82",
// }

func main() {
	tests() 

	// password := exercise01(Input01)
	password := 1

	fmt.Println(password)
}

func tests()  {
		tests := []struct{
		a int
		b int
	}{
		{0, 1},
		{0, 2},
		{0, 100},
		{50, 1000},
	}
	fmt.Println("------- Testing --------")
	for _, test := range(tests) {
		fmt.Printf("%d - %d = %d\n", test.a, test.b, dec(test.a, test.b))
	}
	fmt.Println("------- End --------")
	fmt.Printf("\n\n")
}

func exercise01(input []string) int {
	position := 50
	password := 0

	for _, val := range(input) {

		letter, number := split(val)

		if letter == L {
			position = dec(position, number)
		
		}
		if letter == R {
			position = inc(position, number)
		
		}
		if position == 0 {
			password = password + 1
		}
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

func dec(position int, num int) int {
	return ((position - num) % PIPS + PIPS) % PIPS
}

func inc(position int, num int) int {
	return (position + num) % PIPS
}


