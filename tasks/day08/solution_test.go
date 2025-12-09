package day07

import (
	"AdventOfCode2025/internal/util"
	"testing"
)

type test1 struct {
	name   string
	input  string
	number int
	want   TPart1
}

func run1(t *testing.T, tests []test1, testFun func(input string, number int) (TPart1, error)) {
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got, err := testFun(tc.input, tc.number); got != tc.want {
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
            162,817,812
			57,618,57
			906,360,560
			592,479,940
			352,342,300
			466,668,158
			542,29,236
			431,825,988
			739,650,466
			52,470,668
			216,146,977
			819,987,18
			117,168,530
			805,96,715
			346,949,466
			970,615,88
			941,993,340
			862,61,35
			984,92,344
			425,690,689
            `,
			10,
			40,
		},
		{"Final",
			util.MustReadInput(day, "task.txt"),
			1000,
			24360,
		},
	}

	run1(t, tests, Part1Text)
}
