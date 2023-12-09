package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
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

	// part 1
	// CountPartNumbers(engineTable)

	// part 2
	countGears(engineTable)
}

// part one
func CountPartNumbers(table [][]string) {
	result := 0
	for r, row := range table {
		tempNumberStr := ""
		isNumAdjacent := false
		for c, column := range row {
			isColumnNumeric := regexp.MustCompile(`\d`).MatchString(column)

			if isColumnNumeric {
				tempNumberStr += column
				pattern := "[^0-9.]"
				if checkNeighborCellsSymbols(table, r, c, pattern) == true {
					isNumAdjacent = true
				}

				// finish
				if c == len(row)-1 {
					if len(tempNumberStr) > 0 && isNumAdjacent {
						i, err := strconv.Atoi(tempNumberStr)
						if err != nil {
							panic(err)
						}
						result += i
						fmt.Print(i, " ")
					}
				}
			} else {
				if len(tempNumberStr) > 0 && isNumAdjacent {
					i, err := strconv.Atoi(tempNumberStr)
					if err != nil {
						panic(err)
					}
					result += i
					fmt.Print(i, " ")
				}
				tempNumberStr = ""
				isNumAdjacent = false
			}
		}
	}

	fmt.Println(result)
}

// part two (too lazy to refactor part 1, this is all a mess )
func countGears(table [][]string) {
	result := 0

	// holds location and num of times visited
	gearsMap := map[string][]int{}

	for r, row := range table {
		tempNumberStr := ""
		isNumAdjacent := false
		gearCoordinates := []string{}
		for c, column := range row {
			isColumnNumeric := regexp.MustCompile(`\d`).MatchString(column)

			if isColumnNumeric {
				tempNumberStr += column
				pattern := "[*]"

				hasNeighbors, coordinates := checkNeighborCells(table, r, c, pattern)
				gearCoordinates = coordinates
				if hasNeighbors == true {
					isNumAdjacent = true
					fmt.Println(tempNumberStr, coordinates)
				}

				// apply logic for adding gear here
				if c == len(row)-1 {
					if len(tempNumberStr) > 0 && isNumAdjacent {
						i, err := strconv.Atoi(tempNumberStr)
						if err != nil {
							panic(err)
						}

						for _, xy := range coordinates {
							if len(gearsMap[xy]) == 0 {
								gearsMap[xy] = []int{i}
							} else {
								gearsMap[xy] = append(gearsMap[xy], i)
							}
						}
					}
				}
			} else {
				// add logic for adding gear here
				if len(tempNumberStr) > 0 && isNumAdjacent {
					i, err := strconv.Atoi(tempNumberStr)
					if err != nil {
						panic(err)
					}
					for _, xy := range gearCoordinates {
						if len(gearsMap[xy]) == 0 {
							gearsMap[xy] = []int{i}
						} else {
							gearsMap[xy] = append(gearsMap[xy], i)
						}
					}
				}
				tempNumberStr = ""
				isNumAdjacent = false
			}
		}
	}

	// YET ANOTHER IF STATEMENT??? WHAAT NO WAY FML
	for _, partnumbers := range gearsMap {
		if len(partnumbers) == 2 {
			result += partnumbers[0] * partnumbers[1]
		}
	}

	fmt.Println(gearsMap)
	fmt.Println(result)
}

func checkNeighborCellsSymbols(table [][]string, row int, column int, pattern string) bool {
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

func checkNeighborCells(table [][]string, row int, column int, pattern string) (bool, []string) {
	neighbors := make([]string, 0)

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			newRow, newCol := row+i, column+j

			// Skip the cell itself
			if i == 0 && j == 0 {
				continue
			}

			// Check if the new position is within the bounds of the array
			if newRow >= 0 && newRow < len(table) && newCol >= 0 && newCol < len(table[0]) {
				// Compile the regex
				re, err := regexp.Compile(pattern)

				if err != nil {
					fmt.Println("Error compiling regex:", err)
					return false, neighbors
				}
				if re.MatchString(table[newRow][newCol]) {
					coordinates := strconv.Itoa(newRow) + strconv.Itoa(newCol)
					neighbors = append(neighbors, coordinates)
					return true, neighbors
				}
			}
		}
	}

	return false, neighbors
}
