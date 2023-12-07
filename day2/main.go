package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("input.txt")
	games := strings.Split(string(input), "\n")

	if err != nil {
		panic(err)
	}

	eligibleGamesCount := 0
	for _, game := range games {
		gameCount := playGameRound(game)
		eligibleGamesCount += gameCount
	}
	fmt.Println(eligibleGamesCount)
}

func playGameRound(s string) int {
	re := regexp.MustCompile(`[;:]`)
	gameRoundsData := re.Split(s, -1)
	rounds := gameRoundsData[1:]
	minMap := map[string]int{"blue": 0, "green": 0, "red": 0}
	for i := 0; i < len(rounds); i++ {
		colorMap := getRoundData(rounds[i])
		minMap = countMinColorCubes(minMap, colorMap)
	}

	gameCount := calculateGameValue(minMap)
	return gameCount
}

func getTargetNum() map[string]int {
	gameTarget := map[string]int{"blue": 14, "green": 13, "red": 12}
	return gameTarget
}

func countMinColorCubes(m, c map[string]int) map[string]int {
	result := m

	for color, value := range c {
		if m[color] < value {
			result[color] = value
		} else {
			result[color] = m[color]
		}
	}

	return result
}

func calculateGameValue(m map[string]int) int {
	result := 1
	for _, value := range m {
		result *= value
	}
	return result
}

func getRoundData(s string) map[string]int {
	formattedRounds := strings.TrimSpace(s)
	re := regexp.MustCompile(`[,]`)
	roundColors := re.Split(formattedRounds, -1)
	colorMap := map[string]int{}

	for _, roundColor := range roundColors {
		formatted := strings.TrimSpace(roundColor)
		str := strings.Split(formatted, " ")
		cName := string(str[1])
		cValue, err := strconv.Atoi(str[0])

		if err != nil {
			panic(err)
		}

		colorMap[cName] = cValue
	}

	return colorMap
}

func getGameId(g string) int {
	gameIdStr := string(g[5:])
	gameId, err := strconv.Atoi(gameIdStr)
	if err != nil {
		panic(err)
	}

	return gameId
}
