package day09

import (
	"AdventOfCode2025/internal/util"
	"bufio"
	"fmt"
	"strings"
)

const day = 9

type TPart1 = int64
type TPart2 = TPart1

func Part1() (TPart1, error) {
	return Part1Text(util.Input(day))
}

func Part2() (TPart2, error) {
	return Part2Text(util.Input(day))
}

func abs(i int64) int64 {
	if i < 0 {
		return -i
	}
	return i
}

func Part1Text(input string) (TPart1, error) {
	points, err := parse(input)
	if err != nil {
		return 0, fmt.Errorf("parsing input: %w", err)
	}

	max := int64(0)
	for i := 0; i < len(points); i++ {
		a := points[i]
		for j := i + 1; j < len(points); j++ {
			b := points[j]
			area := (abs(a.x-b.x) + 1) * (abs(a.y-b.y) + 1)
			if area > max {
				max = area
			}
		}
	}

	return max, nil
}

type Point struct {
	id   int
	x, y int64
}

func Part2Text(input string) (TPart2, error) {
	return 0, nil
}

func (p Point) String() string {
	return fmt.Sprintf("%d(%d,%d)", p.id, p.x, p.y)
}

func parse(input string) ([]Point, error) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	var points []Point

	id := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		var p Point
		p.id = id
		if _, err := fmt.Sscanf(line, "%d,%d", &p.x, &p.y); err != nil {
			return nil, fmt.Errorf("invalid line: %s", line)
		}

		points = append(points, p)
		id++
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return points, nil
}
