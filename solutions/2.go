package solutions

import (
	"fmt"
	"strconv"
	"strings"
)

type navigationInput struct {
	direction string
	value     int
}

var solution2 = Solution{
	PartOne: solution2PartOne,
	PartTwo: solution2PartTwo,
}

func solution2PartOne() {
	fmt.Println("Day 2 - Part One")
	input := getInput("2")

	fmt.Println(navigate(input))
}

func solution2PartTwo() {
	fmt.Println("Day 2 - Part Two not implemented")
}

func navigate(input []string) int {
	convertedInput := convertInput(input)

	forwardChange := 0
	depthChange := 0

	for _, input := range convertedInput {
		switch input.direction {
		case "forward":
			forwardChange += input.value
		case "up":
			depthChange -= input.value
		case "down":
			depthChange += input.value
		}
	}

	return forwardChange * depthChange
}

func convertInput(input []string) []navigationInput {
	var convertedInput []navigationInput

	for _, value := range input {
		fields := strings.Fields(value)

		valueAsInt, _ := strconv.Atoi(fields[1])

		convertedInput = append(convertedInput, navigationInput{direction: fields[0], value: valueAsInt})
	}

	return convertedInput
}
