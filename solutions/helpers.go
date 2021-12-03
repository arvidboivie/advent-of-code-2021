package solutions

func sumArray(array []int) int {
	var sum int

	for _, value := range array {
		sum += value
	}

	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
