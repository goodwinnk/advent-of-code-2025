package day08

import (
	"AdventOfCode2025/internal/util"
	"AdventOfCode2025/internal/util/coll"
	"bufio"
	"cmp"
	"fmt"
	"maps"
	"slices"
	"strings"
)

const day = 8

type TPart1 = int64
type TPart2 = TPart1

func input() string {
	return util.MustReadInput(day, "task.txt")
}

func Part1() (TPart1, error) {
	return Part1Text(input(), 1000)
}

func Part2() (TPart2, error) {
	return Part2Text(input())
}

type Edge struct {
	from, to Point
	dist2    int64
}

func Part1Text(input string, number int) (result TPart1, err error) {
	points, err := parse(input)
	if err != nil {
		return 0, fmt.Errorf("parsing input: %w", err)
	}

	edges := coll.NewSmallestN[Edge, int64](number)
	for i := 0; i < len(points); i++ {
		a := points[i]
		for j := i + 1; j < len(points); j++ {
			b := points[j]
			distance2 := distance2(&a, &b)
			edge := Edge{a, b, distance2}
			edges.Push(edge, distance2)
		}
	}

	componentsNumber := make(map[Point]int, len(points))
	components := make(map[int]map[Point]bool, len(points))
	for _, p := range points {
		componentsNumber[p] = p.id
		components[p.id] = map[Point]bool{p: true}
	}

	for _, edge := range edges.PopAllAccending() {
		from, ok := componentsNumber[edge.from]
		if !ok {
			panic(fmt.Errorf("each point should be in a component %v", edge.from))
		}
		to, ok := componentsNumber[edge.to]
		if !ok {
			panic(fmt.Errorf("each point should be in a component %v", edge.from))
		}
		if from != to {
			var target int
			var source int
			if from < to {
				target, source = from, to
			} else {
				target, source = to, from
			}

			for p := range components[source] {
				componentsNumber[p] = target
			}
			maps.Copy(components[target], components[source])
			delete(components, source)
		}
	}

	sizes := coll.NewBiggestN[int, int](3)
	for _, c := range components {
		size := len(c)
		sizes.Push(size, size)
	}

	mult := int64(1)
	for _, v := range sizes.PopAllAccending() {
		mult *= int64(v)
	}

	return mult, nil
}

func distance2(a *Point, b *Point) int64 {
	dx := b.x - a.x
	dy := b.y - a.y
	dz := b.z - a.z

	return dx*dx + dy*dy + dz*dz
}

func distances(points []Point) []Edge {
	edges := make([]Edge, 0, len(points)*len(points)-1/2)
	for i := 0; i < len(points); i++ {
		a := points[i]
		for j := i + 1; j < len(points); j++ {
			b := points[j]
			distance2 := distance2(&a, &b)
			edge := Edge{a, b, distance2}
			edges = append(edges, edge)
		}
	}
	return edges
}

func Part2Text(input string) (TPart2, error) {
	points, err := parse(input)
	if err != nil {
		return 0, fmt.Errorf("parsing input: %w", err)
	}

	edges := distances(points)
	slices.SortStableFunc(edges, func(a, b Edge) int {
		return cmp.Compare(a.dist2, b.dist2)
	})

	componentsNumber := make(map[Point]int, len(points))
	components := make(map[int]map[Point]bool, len(points))
	for _, p := range points {
		componentsNumber[p] = p.id
		components[p.id] = map[Point]bool{p: true}
	}

	edgeIndex := 0
	var edge Edge
	for len(components) > 1 {
		edge = edges[edgeIndex]
		edgeIndex++

		from, ok := componentsNumber[edge.from]
		if !ok {
			panic(fmt.Errorf("each point should be in a component %v", edge.from))
		}
		to, ok := componentsNumber[edge.to]
		if !ok {
			panic(fmt.Errorf("each point should be in a component %v", edge.from))
		}
		if from != to {
			var target int
			var source int
			if from < to {
				target, source = from, to
			} else {
				target, source = to, from
			}

			for p := range components[source] {
				componentsNumber[p] = target
			}
			maps.Copy(components[target], components[source])
			delete(components, source)
		}
	}

	result := edge.from.x * edge.to.x
	return result, nil
}

type Point struct {
	id      int
	x, y, z int64
}

func (p Point) String() string {
	return fmt.Sprintf("%d(%d,%d,%d)", p.id, p.x, p.y, p.z)
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
		if _, err := fmt.Sscanf(line, "%d,%d,%d", &p.x, &p.y, &p.z); err != nil {
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
