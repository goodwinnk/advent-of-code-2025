package day04

import (
	"AdventOfCode2025/internal/util"
	"fmt"
	"strings"
)

func Part1() int {
	text := util.MustReadInput(4, "task.txt")
	return Part1Text(text)
}

func countRoll(x int, y int, l int, maze []string) int {
	if x < 0 || x >= l || y < 0 || y >= len(maze) || maze[y][x] != '@' {
		return 0
	}

	return 1
}

func count3Column(x int, y int, l int, maze []string) int {
	return countRoll(x, y-1, l, maze) +
		countRoll(x, y, l, maze) +
		countRoll(x, y+1, l, maze)
}

func Part1Text(input string) int {
	maze, l := parse(input)
	// fmt.Println(strings.Join(maze, "\n"))

	if len(maze) == 0 {
		return 0
	}

	number := 0
	for y, line := range maze {
		for x := range line {
			if countRoll(x, y, l, maze) == 1 {
				nearby :=
					count3Column(x-1, y, l, maze) +
						count3Column(x, y, l, maze) +
						count3Column(x+1, y, l, maze) -
						1

				if nearby < 4 {
					number++
				}
			}
		}
	}

	return number
}

func parse(input string) ([]string, int) {
	input = strings.TrimSpace(strings.ReplaceAll(input, "\r\n", "\n"))
	lines := strings.Split(input, "\n")

	l := -1
	var result []string
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			continue
		}
		if l == -1 {
			l = len(trimmed)
		} else if l != len(trimmed) {
			panic(fmt.Sprintf("Bad input line, expected size is %d", l))
		}
		result = append(result, trimmed)
	}
	return result, l
}
