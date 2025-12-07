package day07

import (
	"AdventOfCode2025/internal/util"
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

const day = 7

type TPart1 = int
type TPart2 = int64

func input() string {
	return util.MustReadInput(day, "task.txt")
}

func Part1() (TPart1, error) {
	return Part1Text(input())
}

func Part2() (TPart2, error) {
	return Part2Text(input())
}

func Part1Text(input string) (result TPart1, err error) {
	diagram, err := parse(input)
	if err != nil {
		return 0, fmt.Errorf("parsing input: %w", err)
	}
	if len(diagram) < 2 {
		return 0, nil
	}

	indexByte := bytes.IndexByte(diagram[0], 'S')
	if indexByte == -1 {
		return 0, fmt.Errorf("no 'S' found in diagram")
	}

	diagram[1][indexByte] = '|'
	for y := 2; y < len(diagram); y++ {
		line := diagram[y]
		prevLine := diagram[y-1]
		for i, ch := range line {
			if ch == '.' && prevLine[i] == '|' {
				line[i] = '|'
			}
		}

		level := 0
		for i, ch := range line {
			if ch == '^' && prevLine[i] == '|' {
				level++
				if i-1 >= 0 && line[i-1] == '.' {
					line[i-1] = '|'
				}
				if i+1 < len(line) && line[i+1] == '.' {
					line[i+1] = '|'
				}
			}
		}
		result += level
	}

	return result, nil
}

func Part2Text(input string) (TPart2, error) {
	diagram, err := parse(input)
	if err != nil {
		return 0, fmt.Errorf("parsing input: %w", err)
	}
	if len(diagram) < 2 {
		return 0, nil
	}

	indexByte := bytes.IndexByte(diagram[0], 'S')
	if indexByte == -1 {
		return 0, fmt.Errorf("no 'S' found in diagram")
	}

	diagram[1][indexByte] = '|'
	path := make([]int64, len(diagram[1]))
	path[indexByte] = 1

	for y := 2; y < len(diagram); y++ {
		line := diagram[y]
		nextPath := make([]int64, len(line))

		prevLine := diagram[y-1]
		for i, ch := range line {
			if ch == '.' && prevLine[i] == '|' {
				line[i] = '|'
				nextPath[i] += path[i]
			}
		}

		for i, ch := range line {
			if ch == '^' && prevLine[i] == '|' {
				if i-1 >= 0 && (line[i-1] == '.' || line[i-1] == '|') {
					line[i-1] = '|'
					nextPath[i-1] += path[i]
				}
				if i+1 < len(line) && (line[i+1] == '.' || line[i+1] == '|') {
					line[i+1] = '|'
					nextPath[i+1] += path[i]
				}
			}
		}

		path = nextPath
	}

	result := int64(0)
	for _, p := range path {
		result += p
	}

	return result, nil
}

func parse(input string) (diagram [][]byte, err error) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		}

		diagram = append(diagram, []byte(line))
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scanning input: %w", err)
	}

	return
}
