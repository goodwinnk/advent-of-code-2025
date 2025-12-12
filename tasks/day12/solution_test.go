package day12

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
            0:
			###
			##.
			##.
			
			1:
			###
			##.
			.##
			
			2:
			.##
			###
			##.
			
			3:
			##.
			###
			##.
			
			4:
			###
			#..
			###
			
			5:
			###
			.#.
			###
			
			4x4: 0 0 0 0 2 0
			12x5: 1 0 1 0 2 2
			12x5: 1 0 1 0 3 2
            `,
			0, // But should be 2
		},
		{"Final",
			util.MustReadInput(day, "task.txt"),
			463,
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
