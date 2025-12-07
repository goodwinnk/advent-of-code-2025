package day06

import (
	"AdventOfCode2025/internal/util"
	"fmt"
	"strconv"
	"strings"
)

const day = 6

type TPart1 = int64
type TPart2 = TPart1

func input() string {
	return util.MustReadInput(day, "task.txt")
}

func Part1() TPart1 {
	return Part1Text(input())
}

func Part2() TPart2 {
	return Part2Text(input())
}

func Part1Text(input string) TPart1 {
	numbers, operations, err := parse(input)
	if err != nil {
		panic(err)
	}

	result := int64(0)
	for i, operation := range operations {
		var column int64
		if operation == Multiply {
			column = 1
		} else {
			column = 0
		}

		for operandIndex := range numbers {
			value := int64(numbers[operandIndex][i])
			switch operation {
			case Sum:
				{
					column += value
				}
			case Multiply:
				{
					column *= value
				}
			}
		}
		result += column
	}

	return result
}

func Part2Text(input string) TPart1 {
	return Part1Text(input)
}

type Operation int

const (
	Sum Operation = iota
	Multiply
)

func parse(input string) (numbers [][]int, operations []Operation, err error) {
	input = strings.TrimSpace(strings.ReplaceAll(input, "\r\n", "\n"))
	lines := strings.Split(input, "\n")

	notEmptyLines := make([]string, 0, len(lines))
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed != "" {
			notEmptyLines = append(notEmptyLines, trimmed)
		}
	}

	for i := 0; i < len(notEmptyLines)-1; i++ {
		values := make([]int, 0)
		split := strings.Split(notEmptyLines[i], " ")
		for _, v := range split {
			trimmed := strings.TrimSpace(v)
			if trimmed == "" {
				continue
			}

			val, err := strconv.Atoi(trimmed)
			if err != nil {
				return nil, nil, fmt.Errorf("invalid end: %w", err)
			}
			values = append(values, val)
		}
		numbers = append(numbers, values)
	}

	split := strings.Split(notEmptyLines[len(notEmptyLines)-1], " ")
	for _, v := range split {
		trimmed := strings.TrimSpace(v)
		if trimmed == "" {
			continue
		}

		if trimmed == "+" {
			operations = append(operations, Sum)
		} else if trimmed == "*" {
			operations = append(operations, Multiply)
		} else {
			return nil, nil, fmt.Errorf("invalid operation: %q", trimmed)
		}
	}

	for _, values := range numbers {
		if len(values) != len(operations) {
			return nil, nil, fmt.Errorf("invalid number of values: %d", len(values))
		}
	}

	return
}
