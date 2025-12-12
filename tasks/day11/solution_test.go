package day11

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
            aaa: you hhh
			you: bbb ccc
			bbb: ddd eee
			ccc: ddd eee fff
			ddd: ggg
			eee: out
			fff: out
			ggg: out
			hhh: ccc fff iii
			iii: out
            `,
			5,
		},
		{"Final",
			util.MustReadInput(day, "task.txt"),
			470,
		},
	}

	run1(t, tests, Part1Text)
}

func TestPart2(t *testing.T) {
	tests := []test1{
		{"sample",
			`
			svr: aaa bbb
			aaa: fft
			fft: ccc
			bbb: tty
			tty: ccc
			ccc: ddd eee
			ddd: hub
			hub: fff
			eee: dac
			dac: fff
			fff: ggg hhh
			ggg: out
			hhh: out
            `,
			2,
		},
		{"Final",
			util.MustReadInput(day, "task.txt"),
			384151614084875,
		},
	}

	run1(t, tests, Part2Text)
}
