package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	lotteryCards, err := os.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}

	card := strings.Split(string(lotteryCards), "\n")

	for _, c := range card {
		cardSections := strings.Split(string(c), "|")
		winningNumStrings := strings.Split(string(cardSections[0]), ":")[1]
		playerNumStrings := cardSections[1]
		formatCardString(playerNumStrings)

		winningNumbers := formatCardString(winningNumStrings)
		playerNumbers := formatCardString(playerNumStrings)
		cardPoints := countCardPoints(winningNumbers, playerNumbers)

		fmt.Println(winningNumStrings, "______", playerNumStrings)
		fmt.Println(cardPoints)
	}
}

func formatCardString(c string) []int {
	cardNumbers := make([]int, 0, len(c))
	formatted := strings.TrimSpace(c)
	arr := strings.Split(formatted, " ")

	for _, str := range arr {
		if str == "" {
			continue
		}
		trimmedEmptyNums := strings.TrimSpace(str)
		num, err := strconv.Atoi(trimmedEmptyNums)
		if err != nil {
			panic(err)
		}
		cardNumbers = append(cardNumbers, num)
	}

	return cardNumbers
}

func countCardPoints(wStr []int, pStr []int) int {
	sort.Ints(wStr)
	sort.Ints(pStr)

	cardPoints := 0
	winPointer := 0
	playerPointer := 0

	for winPointer < len(wStr) || playerPointer < len(pStr) {
		if wStr[winPointer] == pStr[playerPointer] {
			if cardPoints == 0 {
				cardPoints += 1
			} else {
				cardPoints *= 2
			}

			winPointer = movePointer(winPointer, wStr)
			playerPointer = movePointer(playerPointer, pStr)

		} else if wStr[winPointer] < pStr[playerPointer] {
			winPointer = movePointer(winPointer, wStr)

		} else if wStr[winPointer] > pStr[playerPointer] {
			playerPointer = movePointer(playerPointer, pStr)
		}

		// if either one of the pointers reaches the end of the list, let the other one catch up
		if winPointer == len(wStr)-1 && playerPointer < len(pStr)-1 {
			playerPointer += 1
		}

		if playerPointer == len(pStr)-1 && winPointer < len(wStr)-1 {
			winPointer += 1
		}

		// last iteration
		if winPointer == len(wStr)-1 && playerPointer == len(pStr)-1 {
			if wStr[winPointer] == pStr[playerPointer] {
				if cardPoints == 0 {
					cardPoints += 1
				} else {
					cardPoints *= 2
				}
			}

			winPointer += 1
			playerPointer += 1
		}
	}

	return cardPoints
}

func movePointer(pointerIndex int, arr []int) int {
	if pointerIndex == len(arr)-1 {
		return pointerIndex
	}
	return pointerIndex + 1
}
