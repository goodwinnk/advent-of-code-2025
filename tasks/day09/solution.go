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

func Abs(i int64) int64 {
	if i < 0 {
		return -i
	}
	return i
}

func Part1Text(input string) (TPart1, error) {
	points, err := Parse(input)
	if err != nil {
		return 0, fmt.Errorf("parsing input: %w", err)
	}

	maxArea := int64(0)
	for i := 0; i < len(points); i++ {
		a := points[i]
		for j := i + 1; j < len(points); j++ {
			b := points[j]
			area := (Abs(a.X-b.X) + 1) * (Abs(a.Y-b.Y) + 1)
			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea, nil
}

type Point struct {
	Id   int
	X, Y int64
}

type StraightEdge struct {
	Id                         int
	leftX, rightX, topY, downY int64
}

type Rectange struct {
	leftX, topY, rightX, bottomY int64
}

func (r *Rectange) Intersect(edge StraightEdge) bool {
	if edge.rightX <= r.leftX || r.rightX <= edge.leftX {
		return false
	}

	if edge.downY <= r.topY || r.bottomY <= edge.topY {
		return false
	}

	return true
}

func Part2Text(input string) (TPart2, error) {
	points, err := Parse(input)
	if err != nil {
		return 0, fmt.Errorf("parsing input: %w", err)
	}

	previous := points[len(points)-1]
	edges := make([]StraightEdge, len(points))
	for _, p := range points {
		left := min(previous.X, p.X)
		right := max(previous.X, p.X)
		top := min(previous.Y, p.Y)
		down := max(previous.Y, p.Y)

		edge := StraightEdge{Id: p.Id, leftX: left, rightX: right, topY: top, downY: down}
		edges[p.Id] = edge

		previous = p
	}

	maxArea := int64(0)
	for i := 0; i < len(points); i++ {
		a := points[i]
		for j := i + 1; j < len(points); j++ {
			b := points[j]
			area := (Abs(a.X-b.X) + 1) * (Abs(a.Y-b.Y) + 1)
			if area > maxArea {
				r := Rectange{
					leftX:   min(a.X, b.X),
					topY:    min(a.Y, b.Y),
					rightX:  max(a.X, b.X),
					bottomY: max(a.Y, b.Y),
				}

				intersect := false
				for _, edge := range edges {
					if r.Intersect(edge) {
						intersect = true
						break
					}
				}

				if !intersect {
					maxArea = area
				}
			}
		}
	}

	return maxArea, nil
}

func (p Point) String() string {
	return fmt.Sprintf("%d(%d,%d)", p.Id, p.X, p.Y)
}

func Parse(input string) ([]Point, error) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	var points []Point

	id := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		var p Point
		p.Id = id
		if _, err := fmt.Sscanf(line, "%d,%d", &p.X, &p.Y); err != nil {
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
