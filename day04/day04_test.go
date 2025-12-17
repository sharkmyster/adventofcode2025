package day04

import (
	"slices"
	"testing"
)

func TestDay04Part1(t *testing.T) {
	want := 13
	got := Part1(TestInput04)

	if want != got {
		t.Fatalf("Day 4 Part 1 failed want = %v, got = %v", want, got)
	}
}

func TestDay04Part2(t *testing.T) {
	want := 43
	got := Part2(TestInput04)

	if want != got {
		t.Fatalf("Day 4 Part 1 failed want = %v, got = %v", want, got)
	}
}

func TestSplitRow(t *testing.T) {
	got := splitRow("..@@")
	want := []string{".", ".", "@", "@"}

	if !slices.Equal(got, want) {
		t.Fatalf("CountRolls")
	}
}

func TestCountRolls(t *testing.T) {
	testLocations := [][]string{
		{".", ".", ".", ".", "@", "@", "@", ".", "@", "."},
		{".", ".", ".", ".", "@", ".", "@", ".", "@", "."},
		{".", ".", ".", ".", "@", "@", "@", ".", "@", "."},
	}

	tests := []struct {
		locations [][]string
		x         int
		y         int
		want      int
	}{
		{
			locations: testLocations,
			x:         0,
			y:         0,
			want:      0,
		},
		{
			locations: testLocations,
			x:         0,
			y:         2,
			want:      0,
		},
		{
			locations: testLocations,
			x:         0,
			y:         2,
			want:      0,
		},
		{
			locations: testLocations,
			x:         9,
			y:         0,
			want:      2,
		},
		{
			locations: testLocations,
			x:         9,
			y:         2,
			want:      2,
		},
	}

	for _, tt := range tests {
		got := countAdjacentRolls(tt.locations, tt.x, tt.y)
		if got != tt.want {
			t.Errorf("countRolls = %d; want %d",
				got, tt.want)
		}
	}
}

// func BenchmarkExercise04(b *testing.B) {
// 	Exercise04(Input04)
// }
