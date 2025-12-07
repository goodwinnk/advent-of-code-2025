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
type TPart2 = TPart1

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

		//fmt.Println(y, level, result)
		//for _, row := range diagram {
		//	fmt.Println(string(row))
		//}
	}

	return result, nil
}

func Part2Text(input string) (TPart1, error) {
	return Part1Text(input)
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
