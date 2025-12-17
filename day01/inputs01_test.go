package day01

import "testing"

func TestDay01Part1(t *testing.T) {
	want := 3

	got := Part1(TestInput01)

	if want != got {
		t.Fatalf("Exercise01 wrong number of ticks")
	}
}

func TestDay01Part2(t *testing.T) {
	want := 6

	got := Part2(TestInput01)

	if want != got {
		t.Fatalf("Exercise01 wrong number of ticks")
	}
}