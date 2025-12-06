package main

import (
	"AdventOfCode2025/tasks/day02"
	"AdventOfCode2025/tasks/day03"
	"AdventOfCode2025/tasks/day04"
	"AdventOfCode2025/tasks/day05"
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

	fmt.Println("-- Day 3", strings.Repeat("-", 30))
	fmt.Printf("Part 1: %d\n", day03.Part1())
	fmt.Printf("Part 2: %d\n", day03.Part2())

	fmt.Println("-- Day 4", strings.Repeat("-", 30))
	fmt.Printf("Part 1: %d\n", day04.Part1())
	fmt.Printf("Part 2: %d\n", day04.Part2())

	fmt.Println("-- Day 5", strings.Repeat("-", 30))
	fmt.Printf("Part 1: %d\n", day05.Part1())
	fmt.Printf("Part 2: %d\n", day05.Part2())
}
