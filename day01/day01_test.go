package day01_test

import (
	"os"
	"testing"

	"github.com/brpratt/aoc2022/day01"
)

func TestSolveExample(t *testing.T) {
	file, err := os.Open("test_input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	part1, part2 := day01.Solve(file)

	part1Expected := "24000"
	part2Expected := "45000"

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

	part1, part2 := day01.Solve(file)

	part1Expected := "66616"
	part2Expected := "199172"

	if part1 != part1Expected {
		t.Errorf("[Part 1]: expected %s, got %s", part1Expected, part1)
	}

	if part2 != part2Expected {
		t.Errorf("[Part 2]: expected %s, got %s", part2Expected, part2)
	}
}
