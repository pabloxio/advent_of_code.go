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
		got := StringToInt(test.t)
		if got != test.want {
			t.Errorf("got %d, want %d", got, test.want)
		}
	}
}

func TestSliceStringToInt(t *testing.T) {
	var tests = []struct {
		t    []string
		want []int
	}{
		{[]string{"1", "2", "3"}, []int{1, 2, 3}},
		{[]string{"9", "8", "7"}, []int{9, 8, 7}},
	}

	for _, test := range tests {
		got := SliceStringToInt(test.t)
		for i, n := range got {
			if n != test.want[i] {
				t.Errorf("got %d, want %d", n, test.want[i])
			}
		}
	}
}
