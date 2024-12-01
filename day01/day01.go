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
	build_lists(dat, &list1, &list2)

	sum := 0
	sort.Ints(list1)
	sort.Ints(list2)

	for i, ch := range list1 {
		diff := ch - list2[i]
		abs_diff := max(diff, -diff)
		sum += abs_diff
	}
	fmt.Println(sum)
}

func build_lists(bytes []byte, list1 *list, list2 *list) {
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
