package day03

import "testing"

func TestDay03Part1(t *testing.T) {
	got := Part1(TestInput03)
	want := 357

	if got != want {
		t.Fatalf("Day 3 Part 1 failed!")
	}
}

func TestDay03Part2(t *testing.T) {
	got := Part2(TestInput03)
	want := 3121910778619

	if got != want {
		t.Fatalf("Day 3 Part 1 failed!")
	}
}