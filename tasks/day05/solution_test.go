package day05

import (
	"AdventOfCode2025/internal/util"
	"testing"
)

type test1 struct {
	name  string
	input string
	want  int
}

func run1(t *testing.T, tests []test1, testFun func(input string) int) {
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := testFun(tc.input); got != tc.want {
				t.Fatalf("output = %d, want %d", got, tc.want)
			}
		})
	}
}

type test2 struct {
	name  string
	input string
	want  int64
}

func run2(t *testing.T, tests []test2, testFun func(input string) int64) {
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := testFun(tc.input); got != tc.want {
				t.Fatalf("output = %d, want %d", got, tc.want)
			}
		})
	}
}

func TestPart1(t *testing.T) {
	tests := []test1{
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

	run1(t, tests, Part1Text)
}

func TestPart2(t *testing.T) {
	tests := []test2{
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
			14,
		},
		{"Final",
			util.MustReadInput(day, "task.txt"),
			339668510830757,
		},
	}
	run2(t, tests, Part2Text)
}
