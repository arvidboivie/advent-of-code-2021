package solutions

import (
	"bufio"
	"fmt"
	"os"
)

func readLines(path string) ([]string, error) {
	input, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	defer input.Close()

	var lines []string

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {

		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}
