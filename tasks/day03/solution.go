package day03

import (
	"AdventOfCode2025/internal/util"
	"AdventOfCode2025/internal/util/coll"
	"fmt"
	"strings"
)

func Part1() int64 {
	text := util.MustReadInput(3, "task.txt")
	return Part1Text(text)
}

func Part2() int64 {
	text := util.MustReadInput(3, "task.txt")
	return Part2Text(text)
}

func maxJoltage2(bank []int) int64 {
	return maxJoltage(bank, 2)
}

func maxJoltage12(bank []int) int64 {
	return maxJoltage(bank, 12)
}

func maxJoltage(bank []int, n int) int64 {
	if len(bank) < n {
		panic(fmt.Sprintf("Invalid bank size: %d < %d", len(bank), n))
	}

	result := int64(0)
	prevIndex := -1
	for k := n - 1; k >= 0; k-- {
		prevIndex = coll.MaxIndex(bank[prevIndex+1:len(bank)-k]) + prevIndex + 1
		result = result*10 + int64(bank[prevIndex])
	}

	return result
}

func Part1Text(input string) int64 {
	banks := parse(input)
	return coll.Sum(coll.Map(banks, maxJoltage2))
}

func Part2Text(input string) int64 {
	banks := parse(input)
	return coll.Sum(coll.Map(banks, maxJoltage12))
}

func parse(input string) [][]int {
	input = strings.TrimSpace(strings.ReplaceAll(input, "\r\n", "\n"))
	lines := strings.Split(input, "\n")

	result := make([][]int, len(lines))
	for y, line := range lines {
		result[y] = make([]int, len(line))
		for x, ch := range line {
			result[y][x] = int(ch - '0')
		}
	}

	return result
}
