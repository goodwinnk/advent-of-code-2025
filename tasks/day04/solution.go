package day04

import (
	"AdventOfCode2025/internal/util"
	"fmt"
	"maps"
	"slices"
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

func countRoll(x int, y int, l int, maze [][]byte) int {
	if x < 0 || x >= l || y < 0 || y >= len(maze) || maze[y][x] != '@' {
		return 0
	}

	return 1
}

func Part1Text(input string) int {
	maze, l := parse(input)

	if len(maze) == 0 {
		return 0
	}

	accessible := findAccessible(maze, l, allChecks(maze, l))

	return len(accessible)
}

func Part2Text(input string) int {
	maze, l := parse(input)

	if len(maze) == 0 {
		return 0
	}

	sum := 0

	check := allChecks(maze, l)

	for {
		accessible := findAccessible(maze, l, check)
		sum += len(accessible)
		if len(accessible) == 0 {
			break
		}

		nextCheck := make(map[Cell]bool)
		for _, cell := range accessible {
			roundCells := round(cell)
			maze[cell.y][cell.x] = '.'
			for _, rc := range roundCells {
				nextCheck[rc] = true
			}
		}
		check = slices.Collect(maps.Keys(nextCheck))
	}

	return sum
}

type Cell struct {
	x int
	y int
}

func round(cell Cell) [8]Cell {
	return [8]Cell{
		{cell.x - 1, cell.y - 1},
		{cell.x - 1, cell.y},
		{cell.x - 1, cell.y + 1},
		{cell.x, cell.y - 1},
		{cell.x, cell.y + 1},
		{cell.x + 1, cell.y - 1},
		{cell.x + 1, cell.y},
		{cell.x + 1, cell.y + 1},
	}
}

func allChecks(maze [][]byte, l int) []Cell {
	checks := make([]Cell, l*len(maze))
	for y := 0; y < len(maze); y++ {
		for x := 0; x < l; x++ {
			checks[y*l+x] = Cell{x, y}
		}
	}
	return checks
}

func findAccessible(maze [][]byte, l int, check []Cell) []Cell {
	accessible := make([]Cell, 0)

	for _, cell := range check {
		if countRoll(cell.x, cell.y, l, maze) == 1 {
			roundCells := round(cell)
			nearby := 0
			for _, rc := range roundCells {
				if countRoll(rc.x, rc.y, l, maze) == 1 {
					nearby++
				}
			}

			if nearby < 4 {
				accessible = append(accessible, cell)
			}
		}
	}

	return accessible
}

func parse(input string) ([][]byte, int) {
	input = strings.TrimSpace(strings.ReplaceAll(input, "\r\n", "\n"))
	lines := strings.Split(input, "\n")

	l := -1
	maze := make([][]byte, 0)
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
		maze = append(maze, []byte(trimmed))
	}
	return maze, l
}
