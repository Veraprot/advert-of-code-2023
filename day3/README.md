--- Day 3: Gear Ratios ---
You and the Elf eventually reach a gondola lift station; he says the gondola lift will take you up to the water source, but this is as far as he can bring you. You go inside.

It doesn't take long to find the gondolas, but there seems to be a problem: they're not moving.

"Aaah!"

You turn around to see a slightly-greasy Elf with a wrench and a look of surprise. "Sorry, I wasn't expecting anyone! The gondola lift isn't working right now; it'll still be a while before I can fix it." You offer to help.

The engineer explains that an engine part seems to be missing from the engine, but nobody can figure out which one. If you can add up all the part numbers in the engine schematic, it should be easy to work out which part is missing.

The engine schematic (your puzzle input) consists of a visual representation of the engine. There are lots of numbers and symbols you don't really understand, but apparently any number adjacent to a symbol, even diagonally, is a "part number" and should be included in your sum. (Periods (.) do not count as a symbol.)

Here is an example engine schematic:

467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
In this schematic, two numbers are not part numbers because they are not adjacent to a symbol: 114 (top right) and 58 (middle right). Every other number is adjacent to a symbol and so is a part number; their sum is 4361.

Of course, the actual engine schematic is much larger. What is the sum of all of the part numbers in the engine schematic?


test notes: 

func test() {
	goFile, goErr := os.ReadFile("nums-go.txt")
	if goErr != nil {
		panic(goErr)
	}
	goStr := strings.Split(string(goFile), " ")

	jsFile, err := os.ReadFile("nums-js.txt")
	if err != nil {
		panic(err)
	}
	jsStr := strings.Split(string(jsFile), " ")

	for i, char := range jsStr {
		if goStr[i] != char {
			fmt.Println(i, char)
			fmt.Println(goStr[i], goStr[i-1])
			return
		}
	}
}

func Ints(input []string) []string {
	u := make([]string, 0, len(input))
	m := make(map[string]bool)

	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}

	return u
}

import fs from 'fs';

fs.readFile("./input.txt", (e, data) => {
  const symbols = []
  const numbers = []

  const lines = data.toString().split(/\n/)
  lines.forEach(line => {
    const symbolMatch = [...line.matchAll(/([^\d\.])/g)]
    if (symbolMatch.length == 0) symbols.push([])
    else symbols.push(symbolMatch.map(o => o.index))
  
    const numberMatch = [...line.matchAll(/(\d+)/g)]
    if (numberMatch.length == 0)  numbers.push([])
    else numbers.push(numberMatch.map(obj => {
      return {
        index: obj.index,
        number: parseInt(obj[0]),
        length: obj[0].length
      }
    }))
  })
  
  let numString = ""
  const partNumbersSum = numbers.reduce((acc, numArr, i) => {
    let rowAcc = numArr.reduce((racc, numObj) => {
      const { index, number, length } = numObj
      const currRowSymbolIndicies = symbols[i]
      const upperRowSymbolIndicies = symbols[i-1] || []
      const lowerRowSymbolIndicies = symbols[i+1] || []
      let isPartNumber = false
      for (let j = 0; j < length; j++) {
        const digitIndex = index + j
        if (
        [digitIndex-1, digitIndex+1].some(ni => currRowSymbolIndicies.includes(ni)) ||
        [digitIndex, digitIndex-1, digitIndex+1].some(ni => upperRowSymbolIndicies.includes(ni)) ||
        [digitIndex, digitIndex-1, digitIndex+1].some(ni => lowerRowSymbolIndicies.includes(ni))
        ) {
          isPartNumber = true
          numString += `${number} `
        }
      }
      return isPartNumber
      ? racc + number
      : racc
    }, 0)
    return acc + rowAcc
  }, 0)
  console.log(numString)
  console.log(partNumbersSum)
})