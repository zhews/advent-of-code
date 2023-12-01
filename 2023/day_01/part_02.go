package main

import (
	"strconv"
	"strings"
)

func Part02(input string) int {
	total := 0
	digits := map[int]string{
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
		6: "six",
		7: "seven",
		8: "eight",
		9: "nine",
		0: "zero",
	}
	for _, line := range strings.Split(input, "\n") {
		first := 0
		firstIndex := -1
		last := 0
		lastIndex := -1
		for actual, subString := range digits {
			index := strings.Index(line, subString)
			if index != -1 && (firstIndex == -1 || index < firstIndex) {
				first = actual
				firstIndex = index
			}
			index = strings.Index(line, strconv.Itoa(actual))
			if index != -1 && (firstIndex == -1 || index < firstIndex) {
				first = actual
				firstIndex = index
			}
			index = strings.LastIndex(line, subString)
			if index != -1 && (lastIndex == -1 || index > lastIndex) {
				last = actual
				lastIndex = index
			}
			index = strings.LastIndex(line, strconv.Itoa(actual))
			if index != -1 && (lastIndex == -1 || index > lastIndex) {
				last = actual
				lastIndex = index
			}
		}
		current := 0
		current += first * 10
		current += last
		total += current
	}
	return total
}
