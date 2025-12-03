package main

import (
	"fmt"
	"strings"
)

type Pair struct {
	A string
	B string
}

func Exercise02(TestInput02 string) {
	idStrings := strings.Split(TestInput02, ",")

	for _, idString := range(idStrings) {
		fmt.Println(idString)
	}
}