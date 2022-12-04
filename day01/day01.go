package day01

import (
	"bufio"
	"io"
	"sort"
	"strconv"
)

func Solve(input io.Reader) (string, string) {
	scanner := bufio.NewScanner(input)

	calories := []int{}
	count := 0

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			calories = append(calories, count)
			count = 0
			continue
		}

		calorie, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		count += calorie
	}

	calories = append(calories, count)
	sort.Ints(calories)

	max := calories[len(calories)-1]
	topThree := calories[len(calories)-1] + calories[len(calories)-2] + calories[len(calories)-3]

	return strconv.Itoa(max), strconv.Itoa(topThree)
}
