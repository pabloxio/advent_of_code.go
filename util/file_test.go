package util

import "testing"

func TestReadFileString(t *testing.T) {
	expected := []string{"100", "213", "349", "438"}

	var content []string
	content = ReadFileString("file_test")

	for i, line := range expected {
		if line != content[i] {
			t.Errorf("got %s, want %s", line, content[i])
		}
	}
}

func TestReadFileInt(t *testing.T) {
	expected := []int{100, 213, 349, 438}

	var content []int
	content = ReadFileInt("file_test")

	for i, num := range expected {
		if num != content[i] {
			t.Errorf("got %d, want %d", num, content[i])
		}
	}
}
