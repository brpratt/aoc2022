package main

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/brpratt/aoc2022/day01"
	"github.com/brpratt/aoc2022/day02"
	"github.com/brpratt/aoc2022/day03"
	"github.com/brpratt/aoc2022/day04"
)

type Solution func(io.Reader) (string, string)

var solutions = []Solution{
	day01.Solve,
	day02.Solve,
	day03.Solve,
	day04.Solve,
}

func ExitWithUsage() {
	fmt.Println("Usage: aoc2022 <day>\n\t<day>: number between 1 and 25, inclusive")
	os.Exit(1)
}

func main() {
	if len(os.Args) != 2 {
		ExitWithUsage()
	}

	day, err := strconv.Atoi(os.Args[1])
	if err != nil {
		ExitWithUsage()
	}

	if day < 1 || day > 25 {
		ExitWithUsage()
	}

	if day > len(solutions) {
		fmt.Printf("No solution for day %02d\n", day)
		os.Exit(1)
	}

	input, err := os.Open(fmt.Sprintf("day%02d/input.txt", day))
	if err != nil {
		panic(err)
	}

	defer input.Close()

	part1, part2 := solutions[day-1](input)

	fmt.Printf("%s\n%s\n", part1, part2)
}
