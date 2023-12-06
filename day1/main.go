package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// change to input after pass tests
	calibrationDocument, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	calibrationLines := strings.Split(string(calibrationDocument), "\n")

	result := countCalibrationValues(calibrationLines)
	fmt.Println(result)
}

func countCalibrationValues(s []string) int {
	result := 0

	for _, calibrationLine := range s {
		lineMatches := findMatches(calibrationLine)
		lineValue := GetCalibrationValue(lineMatches)

		i, err := strconv.Atoi(lineValue)
		if err != nil {
			panic(err)
		}

		result += i
	}

	return result
}

func findMatches(text string) []string {
	pattern := `[0-9]|one|two|three|four|five|six|seven|eight|nine`
	re := regexp.MustCompile(pattern)

	var matches []string
	for start := 0; start < len(text); {
		loc := re.FindStringIndex(text[start:])
		if loc == nil {
			break
		}
		match := text[start+loc[0] : start+loc[1]]
		matches = append(matches, match)

		start += loc[0] + 1
	}
	return matches
}

func GetCalibrationValue(matches []string) string {
	digitMap := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	if len(matches) == 0 {
		return "0"
	}

	for i, digit := range matches {
		if digitMap[digit] != "" {
			matches[i] = digitMap[digit]
		}
	}

	return matches[0] + matches[len(matches)-1]
}
