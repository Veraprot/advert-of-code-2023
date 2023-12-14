package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	agricultureMap, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	category := strings.Split(string(agricultureMap), "\n\n")

	fmt.Println(category)
}
