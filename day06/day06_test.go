package day06

import (
	"testing"
)

func TestDay06Part1(t *testing.T) {
	want := 4277556
	got := Part1(TestInput06)

	if want != got {
		t.Fatalf("Day 6 Part 1 failed, want = %v, got = %v", want, got)
	}
}	

func TestDay06Part2(t *testing.T) {
	want := 3263827
	got := Part2(TestInput06)

	if want != got {
		t.Fatalf("Day 6 Part 1 failed, want = %v, got = %v", want, got)
	}
}