package day12

import (
	"AdventOfCode2025/internal/util"
	"AdventOfCode2025/internal/util/coll"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

const day = 12

type TPart1 = int
type TPart2 = TPart1

func Part1() (TPart1, error) {
	return Part1Text(util.Input(day))
}

func Part2() (TPart2, error) {
	return Part2Text(util.Input(day))
}

func Part1Text(input string) (TPart1, error) {
	presents, err := parse(input)
	if err != nil {
		return 0, fmt.Errorf("parsing input: %w", err)
	}

	r := 0
	for _, region := range presents.regions {
		simple := int64((region.height / 3) * (region.width / 3))
		sum := coll.Sum(region.shapes)
		if sum <= simple {
			r++
		} else {
			minSpace := 0
			for i, shapeCount := range region.shapes {
				minSpace += presents.shapes[i].minSpace * shapeCount
			}

			area := region.width * region.height
			if minSpace <= area {
				fmt.Printf("HARD %v, sum=%v > simple=%v, area=%v >= minSpace=%v\n",
					region, sum, simple, area, minSpace,
				)
			}
		}
	}

	return r, nil
}

func Part2Text(input string) (TPart2, error) {
	return 0, nil
}

type shape struct {
	id       int
	minSpace int
	grid     []string
}

type region struct {
	width  int
	height int
	shapes []int
}

type presents struct {
	shapes  []shape
	regions []region
}

func parse(input string) (*presents, error) {
	result := &presents{}
	scanner := bufio.NewScanner(strings.NewReader(input))

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// An empty line or separator might indicate the end of a shapes section
		if line == "" {
			continue
		}

		// Check if this is a shape index line (e.g., "0:")
		if strings.HasSuffix(line, ":") {
			idStr := strings.TrimSuffix(line, ":")
			id, err := strconv.Atoi(idStr)
			if err != nil {
				return nil, fmt.Errorf("invalid shape index: %s", idStr)
			}

			var gridLines []string
			var minSpace int
			for {
				if scanner.Scan() {
					line = strings.TrimSpace(scanner.Text())
					if line == "" {
						break
					}
					minSpace += strings.Count(line, "#")
					gridLines = append(gridLines, line)
				}
			}
			result.shapes = append(result.shapes, shape{id: id, minSpace: minSpace, grid: gridLines})
		} else if strings.Contains(line, "x") && strings.Contains(line, ":") {
			region, err := parseRegion(line)
			if err != nil {
				return nil, err
			}
			result.regions = append(result.regions, region)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func parseRegion(line string) (region, error) {
	parts := strings.Split(line, ":")
	if len(parts) != 2 {
		return region{}, fmt.Errorf("invalid region format: %s", line)
	}

	dimensions := strings.Split(strings.TrimSpace(parts[0]), "x")
	if len(dimensions) != 2 {
		return region{}, fmt.Errorf("invalid dimensions: %s", parts[0])
	}

	width, err := strconv.Atoi(dimensions[0])
	if err != nil {
		return region{}, fmt.Errorf("invalid width: %s", dimensions[0])
	}

	height, err := strconv.Atoi(dimensions[1])
	if err != nil {
		return region{}, fmt.Errorf("invalid height: %s", dimensions[1])
	}

	quantitiesStr := strings.Fields(strings.TrimSpace(parts[1]))
	quantities := make([]int, len(quantitiesStr))
	for i, qStr := range quantitiesStr {
		q, err := strconv.Atoi(qStr)
		if err != nil {
			return region{}, fmt.Errorf("invalid quantity: %s", qStr)
		}
		quantities[i] = q
	}

	return region{
		width:  width,
		height: height,
		shapes: quantities,
	}, nil
}
