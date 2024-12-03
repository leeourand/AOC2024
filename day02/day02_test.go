package main

import (
	"testing"
)

func TestIndexOfError(t *testing.T) {
	expectations := []struct {
		input []Level
		safe  bool
	}{
		{[]Level{1, 2, 3, 2}, true},
		{[]Level{5, 2, 3, 2}, true},
		{[]Level{1, 2, 2, 4}, true},
		{[]Level{1, 2, 7, 3}, true},
		{[]Level{1, 2, 7, 8}, false},
		{[]Level{9, 1, 2, 3}, true},
	}

	for i, exp := range expectations {
		reportIsSafe := reportIsSafe(Report(exp.input))

		if reportIsSafe != exp.safe {
			t.Errorf("Report %d expected to be %t but got %t", i, exp.safe, reportIsSafe)
		}
	}
}
