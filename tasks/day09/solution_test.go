package day09

import (
	"AdventOfCode2025/internal/util"
	"testing"
)

type test1 struct {
	name  string
	input string
	want  TPart1
}

func run1(t *testing.T, tests []test1, testFun func(input string) (TPart1, error)) {
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got, err := testFun(tc.input); got != tc.want {
				if err != nil {
					t.Fatal(err)
				}
				t.Fatalf("output = %d, want %d", got, tc.want)
			}
		})
	}
}

func TestPart1(t *testing.T) {
	tests := []test1{
		{"sample",
			`
            7,1
			11,1
			11,7
			9,7
			9,5
			2,5
			2,3
			7,3
            `,
			50,
		},
		{"Final",
			util.MustReadInput(day, "task.txt"),
			0,
		},
	}

	run1(t, tests, Part1Text)
}

func TestPart2(t *testing.T) {
	tests := []test1{
		{"sample",
			`
            sample
            `,
			0,
		},
		{"Final",
			util.MustReadInput(day, "task.txt"),
			0,
		},
	}

	run1(t, tests, Part2Text)
}
