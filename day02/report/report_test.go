package report

import (
	"testing"
)

func TestIndexOfError(t *testing.T) {
	expectations := []struct {
		input Report
		safe  bool
	}{
		{Report([]Level{1, 2, 3, 2}), true},
		{Report([]Level{5, 2, 3, 2}), true},
		{Report([]Level{1, 2, 2, 4}), true},
		{Report([]Level{1, 2, 7, 3}), true},
		{Report([]Level{1, 2, 7, 8}), false},
		{Report([]Level{9, 1, 2, 3}), true},
	}

	for i, exp := range expectations {
		reportIsSafe := exp.input.IsSafeWithTolerances()

		if reportIsSafe != exp.safe {
			t.Errorf("Report %d expected to be %t but got %t", i, exp.safe, reportIsSafe)
		}
	}
}
