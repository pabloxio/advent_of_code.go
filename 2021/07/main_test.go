package main

import (
	"aoc/util"
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	var tests = []struct {
		filename string
		position int
		want     float64
	}{
		{"input_test", 2, 37},
		{"input_test", 1, 41},
		{"input_test", 3, 39},
		{"input_test", 10, 71},
	}

	for _, test := range tests {
		fileContent := util.ReadFileString(test.filename)
		positions := util.SliceStringToInt(strings.Split(fileContent[0], ","))

		got := part1(positions, test.position)
		if got != test.want {
			t.Errorf("got %f, want %f", got, test.want)
		}
	}
}

func TestPart2(t *testing.T) {
	var tests = []struct {
		filename string
		position int
		want     float64
	}{
		{"input_test", 5, 168},
		{"input_test", 2, 206},
	}

	for _, test := range tests {
		fileContent := util.ReadFileString(test.filename)
		positions := util.SliceStringToInt(strings.Split(fileContent[0], ","))

		got := part2(positions, test.position)
		if got != test.want {
			t.Errorf("got %f, want %f", got, test.want)
		}
	}
}
