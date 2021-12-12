// Usage: From root folder
// $ go run 2021/07/main.go path/to/your/input_file

package main

import (
	"aoc/util"
	"fmt"
	"math"
	"os"
	"strings"
)

func part1(positions []int, position int) float64 {
	var fuel float64

	for _, pos := range positions {
		fuel += math.Abs(float64(pos) - float64(position))
	}

	return fuel
}

func part2(positions []int, position int) float64 {
	var fuel float64

	for _, pos := range positions {
		fuel += util.Summation(math.Abs(float64(pos) - float64(position)))
	}

	return fuel
}

func main() {
	fileContent := util.ReadFileString(os.Args[1])
	positions := util.SliceStringToInt(strings.Split(fileContent[0], ","))

	minFuelPart1 := part1(positions, 0)
	minFuelPart2 := part2(positions, 0)

	var fuel float64
	for i := 1; i < len(positions); i++ {
		fuel = part1(positions, i)
		minFuelPart1 = math.Min(minFuelPart1, fuel)

		fuel = part2(positions, i)
		minFuelPart2 = math.Min(minFuelPart2, fuel)
	}

	fmt.Println("part1", minFuelPart1)
	fmt.Println("part2", int(minFuelPart2))
}
