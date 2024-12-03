package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		fmt.Printf("Something went wrong: %q", e)
		panic(e)
	}
}

type Level int
type Report []Level

func main() {
	reports := buildReports()
	var safe int
	var safeWithTolerances int
	for _, r := range reports {
		if reportIsSafe(r) {
			safe++
		} else if reportIsSafeWithTolerances(r) {
			safeWithTolerances++
		}
	}
	fmt.Printf("Safe reports: %d\n", safe)
	fmt.Printf("Safe reports with tolerances: %d\n", safe+safeWithTolerances)
}

func buildReports() []Report {
	bytes, err := os.ReadFile("input_data/puzzle.txt")
	check(err)
	var reports []Report
	scanner := bufio.NewScanner(strings.NewReader(string(bytes)))
	for scanner.Scan() {
		var report Report
		line := scanner.Text()
		values := strings.Fields(line)

		for _, v := range values {
			num, err := strconv.Atoi(v)

			if err != nil {
				panic("Something went wrong")
			}
			report = append(report, Level(num))
		}
		reports = append(reports, report)
	}
	return reports
}

func reportIsSafeWithTolerances(r Report) bool {
	if hasError(r) {
		for i := range len(r) {
			tweakedReport := make([]Level, len(r)-1) // Assume ElementType is the correct type of elements in Report
			copy(tweakedReport, r[:i])
			copy(tweakedReport[i:], r[i+1:])

			if !hasError(tweakedReport) {
				return true
			}
		}
	} else {
		return true
	}
	return false
}

func reportIsSafe(r Report) bool {
	return !hasError(r)
}

func hasError(r Report) bool {
	var ascending bool

	for i, level := range r {
		if i > 0 {
			previous := r[i-1]
			if ascending {
				if level <= previous {
					return true
				} else if (level - previous) > 3 {
					return true
				}
			} else {
				if level >= previous {
					return true
				} else if (previous - level) > 3 {
					return true
				}
			}
		} else {
			ascending = level < r[i+1]
		}
	}
	return false
}
