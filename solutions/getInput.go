package solutions

import (
	"fmt"
	"os"
	"strconv"
)

func getInput(inputName string) []string {
	input, err := readLines("./solutions/" + inputName + ".input.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	return input
}

func getInputAsIntArray(inputName string) []int {
	input, err := readLines("./solutions/" + inputName + ".input.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	inputAsInt := make([]int, len(input))
	for i := range input {
		inputAsInt[i], _ = strconv.Atoi(input[i])
	}

	return inputAsInt
}
