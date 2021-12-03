package solutions

import (
	"fmt"
)

var solution1 = Solution{
	PartOne: partOne,
	PartTwo: partTwo,
}

func partOne() {
	fmt.Println("Advent of Code 1 - Part One")

	input := getInput("1")

	fmt.Println("Difference: ", getNumberOfIncreases(input))
}

func partTwo() {
	fmt.Println("Advent of Code 1 - Part Two")

	input := getInput("1")

	var ranges []int

	for i := 0; i < len(input); i++ {
		endWindow := i + 3
		if endWindow <= len(input) {
			currentRange := sumArray(input[i:endWindow])
			ranges = append(ranges, currentRange)
		}
	}
	fmt.Println("Difference: ", getNumberOfIncreases(ranges))
}

func getNumberOfIncreases(input []int) int {
	var difference int

	for key, value := range input[1:] {
		previousValue := input[key]
		value := value
		if value > previousValue {
			difference++
		}
	}

	return difference
}
