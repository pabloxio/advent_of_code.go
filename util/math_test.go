package util

import "testing"

func TestSummation(t *testing.T) {
	var tests = []struct {
		number float64
		want   float64
	}{
		{11, 66},
		{4, 10},
		{9, 45},
	}

	for _, test := range tests {
		got := Summation(test.number)

		if got != test.want {
			t.Errorf("got %f, want %f", got, test.want)
		}
	}
}
