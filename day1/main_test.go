package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCalibrationValue(t *testing.T) {
	test := func(input, result string) {
		assert.Equal(t, result, GetCalibrationValue(input))
	}
	// ill assume that all digits can be only positive
	test("1abc2", "12")
	test("pqr3stu8vwx", "38")
	test("a1b2c3d4e5f", "15")
	test("treb7uchet", "77")
	test("trebuchet", "0")
}

func TestCountCalibrationValues(t *testing.T) {
	test := func(input []string, result int) {
		assert.Equal(t, result, countCalibrationValues(input))
	}
	// ill assume that all digits can be only positive
	test([]string{}, 0)
	test([]string{"1abc2", "pqr3stu8vwx"}, 50)
	test([]string{"pqr3stu8vwx"}, 38)
	test([]string{"treb7uchet", "0"}, 77)

	// test string digits
	test([]string{"two1nine"}, 29)
	test([]string{"eightwothree"}, 83)
	test([]string{"abcone2threexyz"}, 13)
	test([]string{"xtwone3four"}, 24)
	test([]string{"4nineeightseven2"}, 42)
	test([]string{"zoneight234"}, 14)
	test([]string{"7pqrstsixteen"}, 76)
}
