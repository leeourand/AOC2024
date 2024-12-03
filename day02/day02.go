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
		if r.IsSafe() {
			safe++
		} else if r.IsSafeWithTolerances() {
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

func (r Report) IsSafeWithTolerances() bool {
	if r.hasError() {
		for i := range len(r) {
			tweakedReport := Report(make([]Level, len(r)-1))
			copy(tweakedReport, r[:i])
			copy(tweakedReport[i:], r[i+1:])

			if !tweakedReport.hasError() {
				return true
			}
		}
	} else {
		return true
	}
	return false
}

func (r Report) IsSafe() bool {
	return !r.hasError()
}

func (r Report) hasError() bool {
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
