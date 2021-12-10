// Usage: From root folder
// $ go run 2021/01/main.go path/to/your/input_file
package main

import (
	"aoc/util"
	"fmt"
	"os"
)

func part1(report []int) int {
	var count int
	previous := report[0]

	for _, measurement := range report {
		current := measurement
		if current > previous {
			count++
		}
		previous = current
	}

	return count
}

func sum(subreport []int) int {
	var count int

	for _, measurement := range subreport {
		count += measurement
	}

	return count
}

func part2(report []int) int {
	var threeMeasurement []int

	for i := 0; i < len(report)-2; i++ {
		threeMeasurement = append(threeMeasurement, sum(report[i:i+3]))
	}

	return part1(threeMeasurement)
}

func main() {
	report := util.ReadFileInt(os.Args[1])

	fmt.Println("part1", part1(report))
	fmt.Println("part2", part2(report))
}
