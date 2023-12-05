package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCalibrationValue(t *testing.T) {
	test := func(input string, result int) {
		assert.Equal(t, result, GetCalibrationValue(input))
	}
	// ill assume that all digits can be only positive
	test("1abc2", 12)
	test("pqr3stu8vwx", 38)
	test("a1b2c3d4e5f", 15)
	test("treb7uchet", 77)
	test("trebuchet", 0)
}
