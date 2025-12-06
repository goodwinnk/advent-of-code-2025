package day05

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
            3-5
			10-14
			16-20
			12-18
			
			1
			5
			8
			11
			17
			32
            `,
			3,
		},
		{"Final",
			util.MustReadInput(day, "task.txt"),
			513,
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
