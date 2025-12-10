package main

import (
	"AdventOfCode2025/tasks/day02"
	"AdventOfCode2025/tasks/day03"
	"AdventOfCode2025/tasks/day04"
	"AdventOfCode2025/tasks/day05"
	"AdventOfCode2025/tasks/day06"
	"AdventOfCode2025/tasks/day07"
	"AdventOfCode2025/tasks/day08"
	"AdventOfCode2025/tasks/day09"
	"AdventOfCode2025/tasks/day10"
	"fmt"
	"strings"

	"AdventOfCode2025/tasks/day01"
)

func assume[T any](f func() (T, error)) T {
	v, err := f()
	if err != nil {
		panic(err)
	}
	return v
}

func day[P1 any, P2 any](day int, p1 P1, p2 P2) {
	fmt.Println("-- Day", day, strings.Repeat("-", 30))
	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
}

func main() {
	fmt.Println("Advent of Code 2025")
	day(1, day01.Part1(), day01.Part2())
	day(2, day02.Part1(), day02.Part2())
	day(3, day03.Part1(), day03.Part2())
	day(4, day04.Part1(), day04.Part2())
	day(5, day05.Part1(), day05.Part2())
	day(6, day06.Part1(), day06.Part2())
	day(7, assume(day07.Part1), assume(day07.Part2))
	day(8, assume(day08.Part1), assume(day08.Part2))
	day(9, assume(day09.Part1), assume(day09.Part2))
	day(10, assume(day10.Part1), assume(day10.Part2))
}
