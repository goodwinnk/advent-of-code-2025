package day03

import (
	"AdventOfCode2025/internal/util"
	"testing"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int64
	}{
		{"sample",
			`987654321111111
			811111111111119
			234234234234278
			818181911112111`,
			357,
		},
		{"987654321111111", "987654321111111", 98},
		{"811111111111119", "811111111111119", 89},
		{"234234234234278", "234234234234278", 78},
		{"818181911112111", "818181911112111", 92},

		{"Final",
			util.MustReadInput(3, "task.txt"),
			17376,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := Part1Text(tc.input); got != tc.want {
				t.Fatalf("output = %d, want %d", got, tc.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int64
	}{
		{"sample",
			`987654321111111
			811111111111119
			234234234234278
			818181911112111`,
			3121910778619,
		},
		{"987654321111111", "987654321111111", 987654321111},
		{"811111111111119", "811111111111119", 811111111119},
		{"234234234234278", "234234234234278", 434234234278},
		{"818181911112111", "818181911112111", 888911112111},

		{"Final",
			util.MustReadInput(3, "task.txt"),
			172119830406258,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := Part2Text(tc.input); got != tc.want {
				t.Fatalf("output = %d, want %d", got, tc.want)
			}
		})
	}
}
