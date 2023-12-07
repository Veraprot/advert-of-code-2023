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
		gameId, eligible := playGameRound(game)
		if eligible {
			eligibleGamesCount += gameId
		}
	}
	fmt.Println(eligibleGamesCount)
}

func playGameRound(s string) (int, bool) {
	re := regexp.MustCompile(`[;:]`)
	gameRoundsData := re.Split(s, -1)
	gameId := getGameId(gameRoundsData[0])
	rounds := gameRoundsData[1:]

	for i := 0; i < len(rounds); i++ {
		if checkGameRound(rounds[i]) == false {
			return gameId, false
		}
	}

	return gameId, true
}

func getTargetNum() map[string]int {
	gameTarget := map[string]int{"blue": 14, "green": 13, "red": 12}
	return gameTarget
}

func checkGameRound(s string) bool {
	formattedRounds := strings.TrimSpace(s)
	re := regexp.MustCompile(`[,]`)
	roundColors := re.Split(formattedRounds, -1)

	target := getTargetNum()
	for _, roundColor := range roundColors {
		formatted := strings.TrimSpace(roundColor)
		str := strings.Split(formatted, " ")
		cName := string(str[1])
		cValue, err := strconv.Atoi(str[0])

		if err != nil {
			panic(err)
		}

		if target[cName] < cValue {
			return false
		}
	}

	return true
}

func getGameId(g string) int {
	gameIdStr := string(g[5:])
	gameId, err := strconv.Atoi(gameIdStr)
	if err != nil {
		panic(err)
	}

	return gameId
}
