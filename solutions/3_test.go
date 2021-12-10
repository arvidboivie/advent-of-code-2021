package solutions

import "testing"

func TestGetPowerConsumption(t *testing.T) {

	input := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}

	result := getPowerConsumption(input)

	expectedResult := 198

	if result != expectedResult {
		t.Fatalf("Expected %d, got %d", expectedResult, result)
	}
}
