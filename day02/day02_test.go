package day02

import (
	"testing"
)

func TestDay02Part1(t *testing.T) {
	got := Part1(TestInput02)
	want := 1227775554

	if got != want {
		t.Fatalf("Day 2 Part 1 failed!")
	}
}

func TestDay02Part2(t *testing.T) {
	got := Part2(TestInput02)
	want := 4174379265

	if got != want {
		t.Fatalf("Day 2 Part 1 failed!")
	}
}