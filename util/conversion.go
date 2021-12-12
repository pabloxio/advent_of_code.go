package util

import (
	"log"
	"strconv"
)

// StringToInt converts a valid int representation to int
func StringToInt(s string) int {
	value, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}

	return value
}

func SliceStringToInt(s []string) []int {
	var values []int

	for _, value := range s {
		values = append(values, StringToInt(value))
	}

	return values
}
