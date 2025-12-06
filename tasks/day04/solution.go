package day04

import (
	"AdventOfCode2025/internal/util"
	"fmt"
	"strings"
)

func input() string {
	return util.MustReadInput(4, "task.txt")
}

func Part1() int {
	return Part1Text(input())
}

func Part2() int {
	return Part2Text(input())
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

	if len(maze) == 0 {
		return 0
	}

	accessible, _ := removeAccessible(maze, l)

	return accessible
}

func Part2Text(input string) int {
	maze, l := parse(input)

	if len(maze) == 0 {
		return 0
	}

	sum := 0

	for {
		accessible, next := removeAccessible(maze, l)
		sum += accessible
		maze = next

		if accessible == 0 {
			break
		}
	}

	// fmt.Println(strings.Join(maze, "\n"))

	return sum
}

func removeAccessible(maze []string, l int) (int, []string) {
	accessibleMap := make([]string, len(maze))

	number := 0
	for y, line := range maze {
		accessibleLine := []byte(line)

		for x, ch := range line {
			if countRoll(x, y, l, maze) == 1 {
				nearby :=
					count3Column(x-1, y, l, maze) +
						count3Column(x, y, l, maze) +
						count3Column(x+1, y, l, maze) -
						1

				if nearby < 4 {
					accessibleLine[x] = 'x'
					number++
				}
			} else {
				if ch == 'x' {
					accessibleLine[x] = '.'
				}
			}
		}

		accessibleMap[y] = string(accessibleLine)
	}

	return number, accessibleMap
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
