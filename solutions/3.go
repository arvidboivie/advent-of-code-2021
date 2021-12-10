package solutions

import (
	"fmt"
	"strconv"
)

var solution3 = Solution{
	PartOne: solution3PartOne,
	PartTwo: solution3PartTwo,
}

func solution3PartOne() {
	fmt.Println("Day 3 - Part One")

	input := getInput("3")

	fmt.Println("Difference: ", getPowerConsumption(input))
}

func solution3PartTwo() {
	fmt.Println("Day 3 - Part Two not implemented")

	// input := getInputAsIntArray("3")

	// fmt.Println("Difference: ", input)
}

func getPowerConsumption(input []string) int {
	length := len(input[0])

	counters := make([]int, length)

	gamma := ""
	epsilon := ""

	for _, value := range input {
		for key, partValue := range value {
			inputValue, _ := strconv.Atoi(string(partValue))

			counters[key] += inputValue
		}
	}

	for _, value := range counters {
		if value > len(input)/2 {
			gamma += "1"
			epsilon += "0"
		} else {
			epsilon += "1"
			gamma += "0"
		}
	}

	gammaInt, _ := strconv.ParseInt(gamma, 2, 64)
	epsilonInt, _ := strconv.ParseInt(epsilon, 2, 64)

	return int(gammaInt * epsilonInt)
}
