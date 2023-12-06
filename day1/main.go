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
		lineValue := GetCalibrationValue(calibrationLine)

		i, err := strconv.Atoi(lineValue)
		if err != nil {
			panic(err)
		}

		result += i
	}

	return result
}

func GetCalibrationValue(s string) string {

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

	re := regexp.MustCompile("[0-9]|one|two|three|four|five|six|seven|eight|nine")

	digitValues := re.FindAllString(s, -1)

	if len(digitValues) == 0 {
		return "0"
	}

	for i, digit := range digitValues {
		if digitMap[digit] != "" {
			digitValues[i] = digitMap[digit]
		}
	}

	return digitValues[0] + digitValues[len(digitValues)-1]
}
