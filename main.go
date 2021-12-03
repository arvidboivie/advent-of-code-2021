package main

import (
	s "arbo/adventofcode/solutions"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, world!")

	var solution string

	if len(os.Args) > 1 {
		solution = os.Args[1]
	}

	if _, ok := s.SolutionMap[solution]; ok {
		s.SolutionMap[solution].PartOne()

		s.SolutionMap[solution].PartTwo()
	}

	os.Exit(1)
}
