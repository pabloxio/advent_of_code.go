// Usage: From root folder
// $ go run 2021/06/main.go path/to/your/input_file

package main

import (
	"aoc/util"
	"fmt"
	"os"
	"strings"
)

var cache map[string]int

func simulate(lanternfish []int, remainingDays int) int {
	if remainingDays == 0 {
		return len(lanternfish)
	}

	var length int
	for _, lf := range lanternfish {
		key := fmt.Sprintf("%d,%d", lf, remainingDays)
		value, hit_cache := cache[key]

		if hit_cache {
			length += value
		} else {
			var temp int
			if lf == 0 {
				temp = simulate([]int{6, 8}, remainingDays-1)
			} else {
				temp = simulate([]int{lf - 1}, remainingDays-1)
			}
			cache[key] = temp
			length += temp
		}
	}

	return length
}

func init() {
	cache = make(map[string]int)
}

func main() {
	days := []int{80, 256}

	for _, day := range days {
		fileContent := util.ReadFileString(os.Args[1])
		initialState := util.SliceStringToInt(strings.Split(fileContent[0], ","))
		fmt.Println("day", day, simulate(initialState, day))
	}
}
