package day05

import (
	"slices"
	"testing"
)

func TestDay05Part1(t *testing.T) {
	want := 3
	got := Part1(TestInput05A, TestInput05B)

	if want != got {
		t.Fatalf("Day 5 Part 1 failed, want = %v, got = %v", want, got)
	}
}	

func TestDay05Part2(t *testing.T) {
	want := 14
	got := Part2(TestInput05A, TestInput05B)

	if want != got {
		t.Fatalf("Day 5 Part 2 failed, want = %v, got = %v", want, got)
	}
}

func TestConvertToRangeSlices(t *testing.T) {
	testRanges := []string{
		"5-10",
		"15-20",
	}

	tests := []struct {
		ranges []string
		want   [][]int
	}{
		{
			ranges: testRanges,
			want:   [][]int{{5, 10}, {15, 20}},
		},
	}

	for _, tt := range tests {
		got := convertToRangeSlices(tt.ranges)
		if !slices.EqualFunc(got, tt.want, slices.Equal) {
			t.Fatalf("convertToRangeSlices")
		}
	}
}
func TestCalculateRanges(t *testing.T) {

	tests := []struct {
		ranges [][]int
		want   [][]int
	}{
		{
			ranges: [][]int{
				{5, 10},
				{8, 15},
				{9, 14},
				{20, 25},
				{20, 24},
				{21, 23},
				{30, 42},
				{42, 55},
			},
			want: [][]int{{5, 15}, {20, 25}, {30, 55}},
		},
		{
			ranges: [][]int{
				{1, 2},
				{3, 4},
				{5, 6},
			},
			want: [][]int{{1, 6}},
		},
	}

	for _, tt := range tests {
		got := calculateRanges(tt.ranges)
		if !slices.EqualFunc(got, tt.want, slices.Equal) {
			t.Fatalf("calculateRanges: got %v want %v", got, tt.want)
		}
	}
}
