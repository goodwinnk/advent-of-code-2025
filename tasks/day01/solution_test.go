package day01

import (
	"AdventOfCode2025/internal/util"
	"testing"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"sample",
			`L68
			L30
			R48
			L5
			R60
			L55
			L1
			L99
			R14
			L82`,
			3},
		{"Final",
			util.MustReadInput(1, "task1.txt"),
			999,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := Part1Text(tc.input); got != tc.want {
				t.Fatalf("Part1 = %d, want %d", got, tc.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"simple", "R1000", 10},
		{"sample",
			`L68
			L30
			R48
			L5
			R60
			L55
			L1
			L99
			R14
			L82`,
			6},
		{"Final",
			util.MustReadInput(1, "task1.txt"),
			6099,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := Part2Text(tc.input); got != tc.want {
				t.Fatalf("Part2 = %d, want %d", got, tc.want)
			}
		})
	}
}
