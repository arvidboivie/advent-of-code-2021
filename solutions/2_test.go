package solutions

import "testing"

func TestNavigation(t *testing.T) {
	testData := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}

	expectedResult := 150

	destination := navigate(testData)

	if destination != expectedResult {
		t.Fatalf("Expected %d, got %d", expectedResult, destination)
	}

}

func TestAdvancedNavigation(t *testing.T) {
	testData := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}

	expectedResult := 900

	destination := navigateAdvanced(testData)

	if destination != expectedResult {
		t.Fatalf("Expected %d, got %d", expectedResult, destination)
	}
}
