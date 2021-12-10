package util

import (
	"bufio"
	"log"
	"os"
)

// ReadFileString returns the filename content as a []string
func ReadFileString(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	var content []string

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		content = append(content, scanner.Text())
	}

	return content
}

// ReadFileInt returns the filename content as a []int
func ReadFileInt(filename string) []int {
	var content []int
	contentString := ReadFileString(filename)

	for _, n := range contentString {
		content = append(content, ToInt(n))
	}

	return content
}
