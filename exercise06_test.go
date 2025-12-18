package main

import (
	"slices"
	"testing"
)

// The rightmost problem is 4 + 431 + 623 = 1058
// The second problem from the right is 175 * 581 * 32 = 3253600
// The third problem from the right is 8 + 248 + 369 = 625
// Finally, the leftmost problem is 356 * 24 * 1 = 8544

func TestGetOperands(t *testing.T) {
	testInput := []string{
		"123 328  51 64 ",
		" 45 64  387 23 ",
		"  6 98  215 314",
	}

	tests := []struct {
		input []string
		want  [][]int
	}{
		{
			input: testInput,
			want:  [][]int{{4, 431, 623}, {175, 581, 32}, {8, 248, 369}, {356, 24, 1}},
		},
	}

	for _, tt := range tests {
		got := parseCephalopodNumbers(tt.input)
		if !slices.EqualFunc(got, tt.want, slices.Equal) {
			t.Fatalf("parseCephalopodNumbers(%v) = %v; want %v", tt.input, got, tt.want)
		}
	}
}
