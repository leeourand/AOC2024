package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		fmt.Printf("Something went wrong: %q", e)
		panic(e)
	}
}

type list []int

func main() {
	dat, err := os.ReadFile("input_data/puzzle.txt")
	check(err)
	var list1 list
	var list2 list
	buildLists(dat, &list1, &list2)

	sumDiffs := sumDiffs(list1, list2)
	sumCounts := sumCounts(list1, list2)

	fmt.Printf("Sum of diffs: %d\n", sumDiffs)
	fmt.Printf("Sum of counts: %d\n", sumCounts)
}

// Part 1
func sumDiffs(list1, list2 list) int {
	sort.Ints(list1)
	sort.Ints(list2)

	var sum int

	for i, ch := range list1 {
		diff := ch - list2[i]
		sum += abs(diff)
	}
	return sum
}

// Part 2
func sumCounts(list1, list2 list) int {
	sum := 0
	for _, ch := range list1 {
		count := 0
		for _, ch2 := range list2 {
			if ch == ch2 {
				count += 1
			}
		}
		sum += count * int(ch)
	}
	return sum
}

func buildLists(bytes []byte, list1 *list, list2 *list) {
	scanner := bufio.NewScanner(strings.NewReader(string(bytes)))
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Fields(line)

		if len(values) < 2 {
			continue
		}

		num1, err1 := strconv.Atoi(values[0])
		num2, err2 := strconv.Atoi(values[1])

		if err1 == nil && err2 == nil {
			*list1 = append(*list1, num1)
			*list2 = append(*list2, num2)
		}
	}
}

func abs(val int) int {
	if val >= 0 {
		return val
	} else {
		return -val
	}
}
