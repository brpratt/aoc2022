package day04

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type interval struct {
	min int
	max int
}

func parseIntervals(input string) (int1 interval, int2 interval) {
	s := strings.Split(input, ",")

	nums1 := strings.Split(s[0], "-")
	nums2 := strings.Split(s[1], "-")

	int1.min, _ = strconv.Atoi(nums1[0])
	int1.max, _ = strconv.Atoi(nums1[1])

	int2.min, _ = strconv.Atoi(nums2[0])
	int2.max, _ = strconv.Atoi(nums2[1])

	return
}

func contains(int1, int2 interval) bool {
	if int1.min <= int2.min && int1.max >= int2.max {
		return true
	}

	return int2.min <= int1.min && int2.max >= int1.max
}

func overlaps(int1, int2 interval) bool {
	return !(int1.max < int2.min || int2.max < int1.min)
}

func Solve(input io.Reader) (string, string) {
	scanner := bufio.NewScanner(input)

	partA := 0
	partB := 0

	for scanner.Scan() {
		line := scanner.Text()

		int1, int2 := parseIntervals(line)

		if contains(int1, int2) {
			partA += 1
		}

		if overlaps(int1, int2) {
			partB += 1
		}
	}

	return strconv.Itoa(partA), strconv.Itoa(partB)
}
