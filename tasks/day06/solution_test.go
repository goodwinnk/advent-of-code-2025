package day06

import (
	"AdventOfCode2025/internal/util"
	"testing"
)

type test1 struct {
	name  string
	input string
	want  TPart1
}

func run1(t *testing.T, tests []test1, testFun func(input string) TPart1) {
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
            123 328  51 64 
			 45 64  387 23 
			  6 98  215 314
			*   +   *   +  
            `,
			4277556,
		},
		{"Final",
			util.MustReadInput(day, "task.txt"),
			4693419406682,
		},
	}

	run1(t, tests, Part1Text)
}

func TestPart2(t *testing.T) {
	tests := []test1{
		{"sample",
			"\n" +
				"123 328  51 64 \n" +
				" 45 64  387 23 \n" +
				"  6 98  215 314\n" +
				"*   +   *   +  \n",
			3263827,
		},
		{"Final",
			util.MustReadInput(day, "task.txt"),
			9029931401920,
		},
	}

	run1(t, tests, Part2Text)
}
