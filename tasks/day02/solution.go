package day02

import (
	"AdventOfCode2025/internal/util"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Part1() int64 {
	text := util.MustReadInput(2, "task.txt")
	return Part1Text(text)
}

func Part2() int64 {
	text := util.MustReadInput(2, "task.txt")
	return Part2Text(text)
}

type IdRange struct {
	start int
	end   int
}

func (t IdRange) String() string {
	return fmt.Sprintf("%d-%d", t.start, t.end)
}

func checkDouble(n int) bool {
	s := strconv.Itoa(n)
	if len(s)%2 == 1 {
		return false
	}

	powHalf := int(math.Pow10(len(s) / 2))
	return n%powHalf == n/powHalf
}

func Part1Text(input string) int64 {
	return countInvalid(input, checkDouble)
}

func checkMultiple(n int) bool {
	s := strconv.Itoa(n)
	sl := len(s)
	for l := 1; l <= sl/2; l++ {
		if sl%l != 0 {
			continue
		}
		pattern := s[:l]
		if strings.Repeat(pattern, sl/l) == s {
			return true
		}
	}
	return false
}

func Part2Text(input string) int64 {
	return countInvalid(input, checkMultiple)
}

func countInvalid(input string, check func(int) bool) int64 {
	ranges := parse(input)

	var r int64 = 0
	for _, idRange := range ranges {
		for i := idRange.start; i <= idRange.end; i++ {
			if check(i) {
				r += int64(i)
			}
		}
	}

	return r
}

func parse(input string) []IdRange {
	input = strings.ReplaceAll(input, "\r\n", "\n")
	idsStr := strings.Split(input, ",")

	result := make([]IdRange, len(idsStr))
	for i, idStr := range idsStr {
		trimmed := strings.TrimSpace(idStr)
		if trimmed == "" {
			continue

		}

		split := strings.Split(trimmed, "-")
		start, err := strconv.Atoi(split[0])
		if err != nil {
			panic(err)
		}
		end, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}

		result[i] = IdRange{
			start: start,
			end:   end,
		}
	}
	return result
}
