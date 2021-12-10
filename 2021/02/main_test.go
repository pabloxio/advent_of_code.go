package main

import (
	"aoc/util"
	"testing"
)

func TestPart1(t *testing.T) {
	got := part1(util.ReadFileString("input_test"))
	want := 150

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	got := part2(util.ReadFileString("input_test"))
	want := 900

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
