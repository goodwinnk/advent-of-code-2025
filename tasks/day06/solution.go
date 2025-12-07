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

	groups := make([]Group, 0, len(operations))
	for i, operation := range operations {
		values := make([]int, 0, len(numbers))
		for y := 0; y < len(numbers); y++ {
			values = append(values, numbers[y][i])
		}
		groups = append(groups, Group{values: values, operation: operation})
	}

	return sumGroups(groups)
}

func Part2Text(input string) TPart1 {
	groups, err := parse2(input)
	if err != nil {
		panic(err)
	}
	return sumGroups(groups)
}

func sumGroups(groups []Group) int64 {
	result := int64(0)

	for _, group := range groups {
		var groupValue int64
		if group.operation == Multiply {
			groupValue = 1
		} else {
			groupValue = 0
		}

		for _, value := range group.values {
			switch group.operation {
			case Sum:
				groupValue += int64(value)
			case Multiply:
				groupValue *= int64(value)
			}
		}

		result += groupValue
	}

	return result
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

type Group struct {
	values    []int
	operation Operation
}

func parse2(input string) (groups []Group, err error) {
	input = strings.TrimSpace(strings.ReplaceAll(input, "\r\n", "\n"))
	lines := strings.Split(input, "\n")

	nonEmptyLines := make([]string, 0, len(lines))
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed != "" {
			nonEmptyLines = append(nonEmptyLines, line)
		}
	}

	i := 0
	for i < len(nonEmptyLines[len(nonEmptyLines)-1]) {
		var operation Operation
		switch operationChar := nonEmptyLines[len(nonEmptyLines)-1][i]; operationChar {
		case '+':
			operation = Sum
		case '*':
			operation = Multiply
		default:
			return nil, fmt.Errorf("invalid operation: %q", operationChar)
		}

		values := make([]int, 0)
		for {
			number := 0
			for y := 0; y < len(nonEmptyLines)-1; y++ {
				var ch uint8
				if i >= len(nonEmptyLines[y]) {
					ch = '0'
				} else {
					ch = nonEmptyLines[y][i]
				}
				if ch >= '0' && ch <= '9' {
					number = number*10 + int(ch-'0')
				} else if ch != ' ' {
					return nil, fmt.Errorf("invalid character for number: %q", ch)
				}
			}
			i++
			if number == 0 {
				break
			}

			values = append(values, number)
		}

		groups = append(groups, Group{values: values, operation: operation})
	}

	return
}
