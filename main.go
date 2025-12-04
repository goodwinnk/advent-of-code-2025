package main

import (
	"AdventOfCode2025/tasks/day02"
	"fmt"
	"strings"

	"AdventOfCode2025/tasks/day01"
)

func main() {
	fmt.Println("Advent of Code 2025")
	fmt.Println("-- Day 1", strings.Repeat("-", 30))
	fmt.Printf("Part 1: %d\n", day01.Part1())
	fmt.Printf("Part 2: %d\n", day01.Part2())

	fmt.Println("-- Day 2", strings.Repeat("-", 30))
	fmt.Printf("Part 1: %d\n", day02.Part1())
	fmt.Printf("Part 2: %d\n", day02.Part2())
}
