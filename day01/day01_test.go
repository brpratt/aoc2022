package day01_test

import (
	"strings"
	"testing"

	"github.com/brpratt/aoc2022/day01"
)

func TestSolve(t *testing.T) {
	input := strings.NewReader(`1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`)

	part1, part2 := day01.Solve(input)

	part1Expected := "24000"
	part2Expected := "45000"

	if part1 != part1Expected {
		t.Errorf("[Part 1]: expected %s, got %s", part1Expected, part1)
	}

	if part2 != part2Expected {
		t.Errorf("[Part 2]: expected %s, got %s", part2Expected, part2)
	}
}
