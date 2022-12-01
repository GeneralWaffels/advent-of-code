package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//Read input file
	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	mostCalories1 := 0
	mostCalories2 := 0
	mostCalories3 := 0

	thisElfCalories := 0

	for sc.Scan() {
		snack, error := strconv.Atoi(sc.Text())
		thisElfCalories += snack

		//If error is different from nil then I found an empty line
		if error != nil {
			if thisElfCalories > mostCalories3 {
				mostCalories3 = thisElfCalories
			}
			if mostCalories3 > mostCalories2 {
				mostCalories3, mostCalories2 = mostCalories2, mostCalories3
			}
			if mostCalories2 > mostCalories1 {
				mostCalories2, mostCalories1 = mostCalories1, mostCalories2
			}
			//I start with a new elf
			thisElfCalories = 0
		}
	}
	fmt.Println(mostCalories1 + mostCalories2 + mostCalories3)
}
