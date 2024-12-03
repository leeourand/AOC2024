package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	. "aoc2024/day02/report"
)

func check(e error) {
	if e != nil {
		fmt.Printf("Something went wrong: %q", e)
		panic(e)
	}
}

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
