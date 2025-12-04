package day01

import (
	"AdventOfCode2025/internal/util"
	"strconv"
	"strings"
)

func Part1() int {
	text := util.MustReadInput(1, "task.txt")
	return Part1Text(text)
}

func Part1Text(input string) int {
	input = strings.ReplaceAll(input, "\r\n", "\n")
	input = strings.TrimSpace(input)
	return Part1Lined(strings.Split(input, "\n"))
}

func Part1Lined(lines []string) int {
	cipher := 100
	count := 0
	number := 50
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			continue
		}

		left := trimmed[0] == 'L'
		repeat, _ := strconv.Atoi(trimmed[1:])
		if left {
			number = (number - repeat + cipher) % cipher
		} else {
			number = (number + repeat) % cipher
		}

		if number == 0 {
			count++
		}
	}

	return count
}

func Part2() int {
	text := util.MustReadInput(1, "task.txt")
	return Part2Text(text)
}

func Part2Text(input string) int {
	input = strings.ReplaceAll(input, "\r\n", "\n")
	input = strings.TrimSpace(input)
	return Part2Lined(strings.Split(input, "\n"))
}

func Part2Lined(lines []string) int {
	cipher := 100
	count := 0
	number := 50
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			continue
		}

		left := trimmed[0] == 'L'
		repeat, _ := strconv.Atoi(trimmed[1:])

		count += repeat / cipher
		notFullRepeat := repeat % cipher

		if left {
			if number != 0 && notFullRepeat >= number {
				count++
			}
			number = (number - notFullRepeat + cipher) % cipher
		} else {
			if number+notFullRepeat >= cipher {
				count++
			}
			number = (number + notFullRepeat) % cipher
		}
	}

	return count
}
