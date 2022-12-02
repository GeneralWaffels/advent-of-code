package main

import (
	"os"
	"fmt"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var score int
	var i int
	file, err := os.ReadFile("input")
	check(err)
	input := string(file)
	for i < len(input){
		if input[i] == 'A' && input[i + 2] == 'Y' {
			score += 8;
		} else if input[i] == 'A' && input[i + 2] == 'Z' {
			score += 3;
		} else if input[i] == 'A' && input[i + 2] == 'X' {
			score += 4;
		} else if input[i] == 'B' && input[i + 2] == 'Y' {
			score += 5;
		} else if input[i] == 'B' && input[i + 2] == 'X' {
			score += 1;
		} else if input[i] == 'B' && input[i + 2] == 'Z' {
			score += 9;
		} else if input[i] == 'C' && input[i + 2] == 'Y' {
			score += 2;
		} else if input[i] == 'C' && input[i + 2] == 'X' {
			score += 7;
		} else if input[i] == 'C' && input[i + 2] == 'Z' {
			score += 6;
		}
		i += 4;
	}
	fmt.Println(score)
}
