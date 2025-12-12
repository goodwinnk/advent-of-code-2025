package day11

import (
	"AdventOfCode2025/internal/util"
	"bufio"
	"fmt"
	"strings"
)

const day = 11

type TPart1 = int64
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

	paths := make(map[string]int64)
	paths["out"] = 1

	visited := make(map[string]bool)
	visited["out"] = true

	var dfs func(name string) int64
	dfs = func(name string) int64 {
		if !visited[name] {
			for _, next := range graph[name] {
				paths[name] += dfs(next)
			}
			visited[name] = true
		}

		return paths[name]
	}

	return dfs("you"), nil
}

func Part2Text(input string) (TPart2, error) {
	graph, err := parse(input)
	if err != nil {
		return 0, fmt.Errorf("parsing input: %w", err)
	}

	type Paths struct {
		none int64
		one  int64
		two  int64
	}

	add := func(p *Paths, op *Paths) Paths {
		return Paths{
			none: p.none + op.none,
			one:  p.one + op.one,
			two:  p.two + op.two,
		}
	}

	raise := func(p *Paths) Paths {
		if p.two > 0 {
			panic(fmt.Sprintf("already paths with two: %v", p))
		}
		return Paths{
			none: 0,
			one:  p.none,
			two:  p.one,
		}
	}

	paths := make(map[string]Paths)
	paths["out"] = Paths{none: 1, one: 0, two: 0}

	var dfs func(name string) Paths
	dfs = func(name string) Paths {
		p, ok := paths[name]
		if !ok {
			p = Paths{none: 0, one: 0, two: 0}

			for _, next := range graph[name] {
				nextP := dfs(next)
				p = add(&p, &nextP)
			}

			if name == "fft" || name == "dac" {
				p = raise(&p)
			}

			paths[name] = p
		}

		return p
	}

	r := dfs("svr")
	return r.two, nil
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
