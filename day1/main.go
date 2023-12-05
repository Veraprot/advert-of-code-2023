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
	re := regexp.MustCompile("[0-9]")
	numValues := re.FindAllString(s, -1)

	if len(numValues) == 0 {
		return "0"
	}

	return numValues[0] + numValues[len(numValues)-1]
}
