package day01

import (
	"testing"
)

type Part int

const (
	Part1 Part = 1
	Part2 Part = 2
)

type Expect struct {
	input  []int
	result int
	part   Part
}

func TestDayOne(t *testing.T) {
	input := readLines("input.txt")

	expects := []Expect{
		Expect{[]int{+1, -2, +3, +1}, 3, Part1},
		Expect{[]int{+1, +1, +1}, 3, Part1},
		Expect{[]int{+1, +1, -2}, 0, Part1},
		Expect{[]int{-1, -2, -3}, -6, Part1},
		Expect{[]int{+1, -2, +3, +1}, 2, Part2},
		Expect{input, 529, Part1},
		Expect{input, 464, Part2},
	}

	for i, x := range expects {
		var res int
		if x.part == Part1 {
			res = sumFrequencies(x.input)
		} else {
			res = firstDupe(x.input)
		}
		if res != x.result {
			t.Errorf("01_p%d_t%d: Expected %d, got %d", x.part, i, x.result, res)
		}
	}
}
