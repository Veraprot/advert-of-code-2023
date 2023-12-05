package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("day 1")
	calibrationDocument := "eightqrssm9httwogqshfxninepnfrppfzhsc \n one111jxlmc7tvklrmhdpsix \n bptwone4sixzzppg \n ninezfzseveneight5kjrjvtfjqt5nineone"

	calibrationLines := strings.Split(calibrationDocument, "\n")
	fmt.Println(calibrationLines[0])
}

func GetCalibrationValue(s string) string {
	return s
}

// for each line run algo and return result
// add result to final result
// print result
