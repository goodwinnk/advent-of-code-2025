package day10

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
            [.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
			[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}
			[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}
            `,
			7,
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
            [.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
			[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}
			[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}
            `,
			33,
		},
		{"Final",
			util.MustReadInput(day, "task.txt"),
			0,
		},
	}

	run1(t, tests, Part2Text)
}
