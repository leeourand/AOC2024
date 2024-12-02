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
	dat, err := os.ReadFile("input_data/puzzle.txt")
	check(err)
	var reports []Report
	buildReports(dat, &reports)
	var safe int
	for _, r := range reports {
		if reportIsSafe(r) {
			safe++
		}
	}
	fmt.Printf("Safe reports: %d\n", safe)
}

func buildReports(bytes []byte, reports *[]Report) {
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
		*reports = append(*reports, report)
	}
}

func reportIsSafe(r Report) bool {
	var ascending bool

	for i, level := range r {
		if i > 0 {
			previous := r[i-1]
			if ascending {
				if level <= previous {
					return false
				} else if (level - previous) > 3 {
					return false
				}
			} else {
				if level >= previous {
					return false
				} else if (previous - level) > 3 {
					return false
				}
			}
		} else {
			ascending = level < r[i+1]
		}
	}
	return true
}
