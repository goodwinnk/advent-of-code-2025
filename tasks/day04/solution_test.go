package day04

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
			`
            ..@@.@@@@.
			@@@.@.@.@@
			@@@@@.@.@@
			@.@@@@..@.
			@@.@@@@.@@
			.@@@@@@@.@
			.@.@.@.@@@
			@.@@@.@@@@
			.@@@@@@@@.
			@.@.@@@.@.
            `,
			13,
		},
		{"Final",
			util.MustReadInput(4, "task.txt"),
			1445,
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
		want  int
	}{
		{"sample",
			`
            ..@@.@@@@.
			@@@.@.@.@@
			@@@@@.@.@@
			@.@@@@..@.
			@@.@@@@.@@
			.@@@@@@@.@
			.@.@.@.@@@
			@.@@@.@@@@
			.@@@@@@@@.
			@.@.@@@.@.
            `,
			43,
		},
		{"Final",
			util.MustReadInput(4, "task.txt"),
			8317,
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
