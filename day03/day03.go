package main

import (
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		fmt.Printf("Something went wrong: %q", e)
		panic(e)
	}
}

func main() {
	bytes, err := os.ReadFile("input_data/puzzle.txt")
	check(err)
	var sum int
	parsing := true

	for i, byte := range bytes {
		if byte == 'd' {
			if bytes[i+1] == 'o' {
				if bytes[i+2] == 'n' {
					if bytes[i+3] == '\'' {
						if bytes[i+4] == 't' {
							if bytes[i+5] == '(' {
								if bytes[i+6] == ')' {
									parsing = false
								}
							}
						}
					}
				}
				if bytes[i+2] == '(' {
					if bytes[i+3] == ')' {
						parsing = true
					}
				}
			}
		}

		if byte == 'm' && parsing {
			if bytes[i+1] == 'u' {
				if bytes[i+2] == 'l' {
					if bytes[i+3] == '(' {
						j := i + 4
						for isDigit(bytes[j]) {
							j++
						}
						strNum := bytes[(i + 4):j]
						num1, err := strconv.Atoi(string(strNum))
						check(err)

						if bytes[j] == ',' {
							j++
							start := j
							for isDigit(bytes[j]) {
								j++
							}
							num2, err := strconv.Atoi(string(bytes[(start):j]))
							check(err)

							if bytes[j] == ')' {
								sum += num1 * num2
							}
						}
					}
				}
			}
		}
	}
	fmt.Println(sum)
}

func isDigit(char byte) bool {
	return char >= '0' && char <= '9'
}
