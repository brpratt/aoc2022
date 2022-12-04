package day02_test

import (
	"os"
	"testing"

	"github.com/brpratt/aoc2022/day02"
)

func TestSolveExample(t *testing.T) {
	file, err := os.Open("test_input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	part1, part2 := day02.Solve(file)

	part1Expected := "15"
	part2Expected := "12"

	if part1 != part1Expected {
		t.Errorf("[Part 1]: expected %s, got %s", part1Expected, part1)
	}

	if part2 != part2Expected {
		t.Errorf("[Part 2]: expected %s, got %s", part2Expected, part2)
	}
}

func TestSolveActual(t *testing.T) {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	part1, part2 := day02.Solve(file)

	part1Expected := "8890"
	part2Expected := "10238"

	if part1 != part1Expected {
		t.Errorf("[Part 1]: expected %s, got %s", part1Expected, part1)
	}

	if part2 != part2Expected {
		t.Errorf("[Part 2]: expected %s, got %s", part2Expected, part2)
	}
}
