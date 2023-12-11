package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	lotteryCards, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	card := strings.Split(string(lotteryCards), "\n")
	cardPoints := map[int]int{}
	for i, c := range card {
		cardSections := strings.Split(string(c), "|")
		winningNumStrings := strings.Split(string(cardSections[0]), ":")[1]
		playerNumStrings := cardSections[1]
		formatCardString(playerNumStrings)

		winningNumbers := formatCardString(winningNumStrings)
		playerNumbers := formatCardString(playerNumStrings)
		cardPoints[i] = countCardPoints(winningNumbers, playerNumbers)
	}
	countScratchcards(cardPoints)
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

func countScratchcards(cardMap map[int]int) {
	copies := map[int]int{0: 0}

	for i := 0; i < len(cardMap); i++ {
		copies[i] += 1
		for j := 1; j <= cardMap[i]; j++ {
			copies[i+j] += copies[i]
		}
	}
	result := 0
	for _, value := range copies {
		result += value
	}

	fmt.Println(result)
}

func countCardPoints(wStr []int, pStr []int) int {
	sort.Ints(wStr)
	sort.Ints(pStr)

	cardPoints := 0
	winPointer := 0
	playerPointer := 0

	for winPointer < len(wStr) || playerPointer < len(pStr) {
		// last iteration
		if winPointer == len(wStr)-1 && playerPointer == len(pStr)-1 {
			if wStr[winPointer] == pStr[playerPointer] {
				cardPoints += 1
			}

			winPointer += 1
			playerPointer += 1
			break
		}

		// if either one of the pointers reaches the end of the list, let the other one catch up
		if winPointer == len(wStr)-1 && playerPointer < len(pStr)-1 {
			if wStr[winPointer] == pStr[playerPointer] {
				cardPoints += 1
			}
			playerPointer += 1
		} else if playerPointer == len(pStr)-1 && winPointer < len(wStr)-1 {
			if wStr[winPointer] == pStr[playerPointer] {
				cardPoints += 1
			}
			winPointer += 1
		} else {
			cardPoints, winPointer, playerPointer = checkNumbers(cardPoints, wStr, pStr, winPointer, playerPointer)
		}
	}

	return cardPoints
}
func checkNumbers(cardPoints int, wStr []int, pStr []int, winPointer int, playerPointer int) (int, int, int) {
	if wStr[winPointer] == pStr[playerPointer] {
		cardPoints += 1

		winPointer = movePointer(winPointer, wStr)
		playerPointer = movePointer(playerPointer, pStr)

	} else if wStr[winPointer] < pStr[playerPointer] {
		winPointer = movePointer(winPointer, wStr)

	} else if wStr[winPointer] > pStr[playerPointer] {
		playerPointer = movePointer(playerPointer, pStr)
	}

	return cardPoints, winPointer, playerPointer
}

func movePointer(pointerIndex int, arr []int) int {
	if pointerIndex == len(arr)-1 {
		return pointerIndex
	}
	return pointerIndex + 1
}
