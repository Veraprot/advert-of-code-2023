package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	// change to input after pass tests
	calibrationDocument := "eightqrssm9httwogqshfxninepnfrppfzhsc\none111jxlmc7tvklrmhdpsix\nbptwone4sixzzppg\nninezfzseveneight5kjrjvtfjqt5nineone"

	calibrationLines := strings.Split(calibrationDocument, "\n")
	countCalibrationValues(calibrationLines)
}

func countCalibrationValues(s []string) {
	for _, calibrationLine := range s {
		fmt.Println(calibrationLine)
		GetCalibrationValue(calibrationLine)
	}
}

func GetCalibrationValue(s string) string {
	re := regexp.MustCompile("[0-9]")
	numValues := re.FindAllString(s, -1)

	fmt.Println(numValues)

	if len(numValues) == 0 {
		return "0"
	}

	return numValues[0] + numValues[len(numValues)-1]
}
