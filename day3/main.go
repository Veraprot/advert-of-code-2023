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
	engineSchema, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	engineRows := strings.Split(string(engineSchema), "\n")

	engineTable := [][]string{}

	for _, row := range engineRows {
		engineColumns := strings.Split(row, "")
		engineTable = append(engineTable, engineColumns)
	}

	traverseEngineTable(engineTable)
}

func traverseEngineTable(table [][]string) {
	result := 0
	for r, row := range table {
		tempNumberStr := ""
		isNumAdjacent := true
		for c, column := range row {
			isColumnNumeric := regexp.MustCompile(`\d`).MatchString(column)
			if isColumnNumeric {
				tempNumberStr += column
				if checkNeighborCellsSymbols(table, r, c) == true {
					isNumAdjacent = true
				}
			} else {
				if len(tempNumberStr) > 0 && isNumAdjacent {
					i, err := strconv.Atoi(tempNumberStr)
					if err != nil {
						panic(err)
					}

					result += i
				}
				tempNumberStr = ""
				isNumAdjacent = false
			}
		}
	}

	fmt.Println(result)
}

func checkNeighborCellsSymbols(table [][]string, row int, column int) bool {
	// neighbors := make([]int, 0)

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			newRow, newCol := row+i, column+j

			// Skip the cell itself
			if i == 0 && j == 0 {
				continue
			}

			// Check if the new position is within the bounds of the array
			if newRow >= 0 && newRow < len(table) && newCol >= 0 && newCol < len(table[0]) {
				pattern := "[^0-9.]"

				// Compile the regex
				re, err := regexp.Compile(pattern)

				if err != nil {
					fmt.Println("Error compiling regex:", err)
					return false
				}
				if re.MatchString(table[newRow][newCol]) {
					return true
				}
			}
		}
	}

	return false
}
