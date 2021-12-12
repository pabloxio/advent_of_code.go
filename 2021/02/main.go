// Usage: From root folder
// $ go run 2021/02/main.go path/to/your/input_file
package main

import (
	"aoc/util"
	"fmt"
	"os"
	"strings"
)

func parseCommand(command string) (string, int) {
	c := strings.Split(command, " ")

	return c[0], util.StringToInt(c[1])
}

func part1(commands []string) int {
	var horizontalPosition, depth int

	for _, command := range commands {
		dir, value := parseCommand(command)
		switch dir {
		case "forward":
			horizontalPosition += value
		case "down":
			depth += value
		case "up":
			depth -= value
		}
	}

	return horizontalPosition * depth
}

func part2(commands []string) int {
	var horizontalPosition, depth, aim int

	for _, command := range commands {
		dir, value := parseCommand(command)
		switch dir {
		case "down":
			aim += value
		case "up":
			aim -= value
		case "forward":
			horizontalPosition += value
			depth += aim * value
		}
	}

	return horizontalPosition * depth
}

func main() {
	commands := util.ReadFileString(os.Args[1])

	fmt.Println("part1", part1(commands))
	fmt.Println("part2", part2(commands))
}
