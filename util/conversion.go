package util

import (
	"log"
	"strconv"
)

func ToInt(s string) int {
	value, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}

	return value
}
