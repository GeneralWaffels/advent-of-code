package main

import (
	"fmt"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var score int
	var i int
	file, err := os.ReadFile("input.txt")
	check(err)
	input := string(file)
	for i < len(input) {
		if input[i] == 'A' && input[i+2] == 'Y' {
			score += 4
		} else if input[i] == 'A' && input[i+2] == 'Z' {
			score += 8
		} else if input[i] == 'A' && input[i+2] == 'X' {
			score += 3
		} else if input[i] == 'B' && input[i+2] == 'Y' {
			score += 5
		} else if input[i] == 'B' && input[i+2] == 'X' {
			score += 1
		} else if input[i] == 'B' && input[i+2] == 'Z' {
			score += 9
		} else if input[i] == 'C' && input[i+2] == 'Y' {
			score += 6
		} else if input[i] == 'C' && input[i+2] == 'X' {
			score += 2
		} else if input[i] == 'C' && input[i+2] == 'Z' {
			score += 7
		}
		i += 4
	}
	fmt.Println(score)
}
