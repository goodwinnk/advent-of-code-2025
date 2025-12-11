package day00

import (
	"AdventOfCode2025/internal/util"
	"bufio"
	"fmt"
	"strings"
)

const day = 0

type TPart1 = int
type TPart2 = TPart1

func Part1() (TPart1, error) {
	return Part1Text(util.Input(day))
}

func Part2() (TPart2, error) {
	return Part2Text(util.Input(day))
}

func Part1Text(input string) (TPart1, error) {
	graph, err := parse(input)
	if err != nil {
		return 0, fmt.Errorf("parsing input: %w", err)
	}

	println(len(graph))

	return 0, nil
}

func Part2Text(input string) (TPart2, error) {
	return 0, nil
}

func parse(input string) (map[string][]string, error) {
	graph := make(map[string][]string)

	sc := bufio.NewScanner(strings.NewReader(input))
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" {
			continue
		}

		// Expect format: name: a b c
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid line: %q", line)
		}

		name := strings.TrimSpace(parts[0])
		rhs := strings.TrimSpace(parts[1])
		if name == "" {
			return nil, fmt.Errorf("invalid line: %q", line)
		}

		if rhs == "" {
			graph[name] = nil
			continue
		}

		// split by spaces (multiple spaces allowed)
		tokens := strings.Fields(rhs)
		for _, t := range tokens {
			graph[name] = append(graph[name], t)
		}
	}

	return graph, nil
}
