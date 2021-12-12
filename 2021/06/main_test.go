package main

import (
	"aoc/util"
	"strings"
	"testing"
)

func TestSimulate(t *testing.T) {
	var tests = []struct {
		filename string
		days     int
		want     int
	}{
		{"input_test", 1, 5},
		{"input_test", 2, 6},
		{"input_test", 18, 26},
		{"input_test", 80, 5934},
		{"input_test", 256, 26984457539},
	}

	for _, test := range tests {
		file := util.ReadFileString(test.filename)
		state := util.SliceStringToInt(strings.Split(file[0], ","))
		got := simulate(state, test.days)
		if got != test.want {
			t.Errorf("got %d, want %d", got, test.want)
		}
	}
}
