package solutions

import (
	"fmt"
)

var solutionX = Solution{
	PartOne: solutionXPartOne,
	PartTwo: solutionXPartTwo,
}

func solutionXPartOne() {
	fmt.Println("Day X - Part One")

	input := getInputAsIntArray("X")

	fmt.Println("Difference: ", input)
}

func solutionXPartTwo() {
	// fmt.Println("Day X - Part Two")

	// input := getInputAsIntArray("X")

	// fmt.Println("Difference: ", input)
}
