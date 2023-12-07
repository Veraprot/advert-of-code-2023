// Game 1: 3 green, 1 blue, 3 red; 3 blue, 1 green, 3 red; 2 red, 12 green, 7 blue; 1 red, 4 blue, 5 green; 7 green, 2 blue, 2 red
// Game 2: 1 green, 19 blue, 1 red; 8 blue, 4 red; 3 red, 6 blue; 1 green, 1 red, 12 blue
// Game 3: 3 green, 1 blue, 9 red; 1 blue, 2 green, 8 red; 1 blue, 2 red
// Game 4: 6 green, 2 red; 2 red, 16 green; 3 red, 1 blue

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckGame(t *testing.T) {
	test := func(input string, result int) {
		assert.Equal(t, result, checkGame(input))
	}
	// ill assume that all digits can be only positive
	test("Game 4: 6 green, 2 red; 2 red, 16 green; 3 red, 1 blue", 4)
}
