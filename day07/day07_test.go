package day07

import (
	"testing"
)

func TestDay07Part1(t *testing.T) {
	want := 21
	got := Part1(TestInput07)

	if want != got {
		t.Fatalf("Day 7 Part 1 failed, want = %v, got = %v", want, got)
	}
}	
func TestDay07Part2(t *testing.T) {
	want := 40
	got := Part2(TestInput07)

	if want != got {
		t.Fatalf("Day 7 Part 2 failed, want = %v, got = %v", want, got)
	}
}	

func TestDay07Part2Proper(t *testing.T) {
	want := 16716444407407
	got := Part2(Input07)

	if want != got {
		t.Fatalf("Day 7 Part 2 Proper failed, want = %v, got = %v", want, got)
	}
}	
