package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//Reads input file input.txt and creates scan lines
	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	//Searches for the maximum sum of calories
	maxCalories := 0
	thisElfCalories := 0

	for sc.Scan() {
		snack, error := strconv.Atoi(sc.Text())
		thisElfCalories += snack

		//If error is different from nil then I found an empty line
		if error != nil {
			if thisElfCalories > maxCalories {
				maxCalories = thisElfCalories
			}
			//I start with a new elf
			thisElfCalories = 0
		}
	}
	fmt.Println(maxCalories)
}
