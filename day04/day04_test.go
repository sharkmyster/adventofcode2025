package day04

import (
	"slices"
	"testing"
)

func TestSplitRow(t *testing.T) {
	got := splitRow("..@@")
	want := []string{".", ".", "@", "@"}

	if !slices.Equal(got, want) {
		t.Fatalf("CountRolls")
	}
}

// func BenchmarkSplitRow(b *testing.B)  {
// 	splitRow("..@@")
// }

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

func BenchmarkExercise04(b *testing.B) {
	Exercise04(Input04)
}
