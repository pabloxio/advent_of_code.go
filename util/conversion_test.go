package util

import "testing"

func TestToInt(t *testing.T) {
	var tests = []struct {
		t    string
		want int
	}{
		{"3", 3},
		{"190", 190},
		{"1000", 1000},
	}

	for _, test := range tests {
		got := ToInt(test.t)
		if got != test.want {
			t.Errorf("got %d, want %d", got, test.want)
		}
	}
}
