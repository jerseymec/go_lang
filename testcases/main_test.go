package main

import "testing"

func TestSquares(t *testing.T) {
	testCases := []struct {
		description string
		input       int
		expected    int
	}{
		{"correctly squares a number", 4, 16},
		{"correctly squares another number", 3, 9},
		{"correctly square a negative number", -2, 4},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			if result := Squares(testCase.input); result != testCase.expected {
				t.Errorf("Incorrect value, got: %d, expected %d", testCase.input, testCase.expected)
			}

		})
	}
}
