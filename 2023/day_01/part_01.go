package main

import (
	"strconv"
	"strings"
	"unicode"
)

func Part01(input string) int {
	total := 0
	for _, line := range strings.Split(input, "\n") {
		digit := []rune{}
		for i := 0; i < len(line); i++ {
			char := rune(line[i])
			if unicode.IsDigit(char) {
				digit = append(digit, char)
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			char := rune(line[i])
			if unicode.IsDigit(char) {
				digit = append(digit, char)
				break
			}
		}
		digitString := string(digit)
		if digitString == "" {
			continue
		}
		parsedDigit, err := strconv.Atoi(digitString)
		if err != nil {
			return 0
		}
		total += parsedDigit
	}
	return total
}
