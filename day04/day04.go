package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		fmt.Printf("Something went wrong: %q", e)
		panic(e)
	}
}

type Puzzle struct {
	characters string
	word       string
	width      int
	height     int
}

func (p Puzzle) countWords() int {
	count := 0

	for i, _ := range p.characters {
		count += p.checkLeftRight(i)
		count += p.checkRightLeft(i)
		count += p.checkTopBottom(i)
		count += p.checkBottomTop(i)
		count += p.checkDiagDownRight(i)
		count += p.checkDiagUpRight(i)
		count += p.checkDiagDownLeft(i)
		count += p.checkDiagUpLeft(i)
	}
	return count
}

func (p *Puzzle) calculateSize() {
	for i, char := range p.characters {
		if char == '\n' || char == '\r' {
			p.width = i
			p.height = i + 1
			return
		}
	}
}

func (p Puzzle) checkLeftRight(i int) int {
	starting_col := i % (p.height)

	if starting_col <= (p.width - len(p.word)) {
		for j, char := range p.word {
			if char == rune(p.characters[i+j]) {
				continue
			} else {
				return 0
			}
		}
	} else {
		return 0
	}

	return 1
}

func (p Puzzle) checkRightLeft(i int) int {
	starting_col := i % (p.height)
	if starting_col >= (len(p.word) - 1) {
		for j, char := range p.word {
			if char == rune(p.characters[i-j]) {
				continue
			} else {
				return 0
			}
		}
	} else {
		return 0
	}

	return 1
}

func (p Puzzle) checkTopBottom(i int) int {
	starting_row := i / (p.height)
	starting_col := i % (p.height)

	if starting_row < (p.height - len(p.word)) {
		for j, char := range p.word {
			row := starting_row + j
			next_index := row*p.height + starting_col
			found_char := rune(p.characters[next_index])
			if char == found_char {
				continue
			} else {
				return 0
			}
		}
	} else {
		return 0
	}

	return 1
}

func (p Puzzle) checkBottomTop(i int) int {
	starting_row := i / (p.height)
	if starting_row >= (len(p.word) - 1) {
		for j, char := range p.word {
			row := starting_row - j
			next_index := row*(p.height) + i%(p.height)
			found_char := rune(p.characters[next_index])
			if char == found_char {
				continue
			} else {
				return 0
			}
		}
	} else {
		return 0
	}

	return 1
}

func (p Puzzle) checkDiagDownRight(i int) int {
	starting_row := i / (p.height)
	starting_col := i % (p.height)
	if starting_row < (p.height-len(p.word)) && starting_col <= p.width-(len(p.word)) {

		for j, char := range p.word {
			row := starting_row + j
			next_col := starting_col + j
			next_index := row*(p.height) + next_col
			found_char := rune(p.characters[next_index])
			if char == found_char {
				continue
			} else {
				return 0
			}
		}
	} else {
		return 0
	}

	return 1
}

func (p Puzzle) checkDiagDownLeft(i int) int {
	starting_row := i / (p.height)
	starting_col := i % (p.height)
	if starting_row <= (p.height-len(p.word)) && starting_col >= len(p.word)-1 {
		for j, char := range p.word {
			row := starting_row + j
			next_col := starting_col - j
			next_index := row*(p.height) + next_col
			found_char := rune(p.characters[next_index])
			if char == found_char {
				continue
			} else {
				return 0
			}
		}
	} else {
		return 0
	}

	return 1
}

func (p Puzzle) checkDiagUpRight(i int) int {
	starting_row := i / (p.height)
	starting_col := i % (p.height)
	if starting_row >= (len(p.word)-1) && starting_col <= p.width-(len(p.word)) {
		for j, char := range p.word {
			row := starting_row - j
			next_col := starting_col + j
			next_index := row*(p.height) + next_col
			found_char := rune(p.characters[next_index])
			if char == found_char {
				continue
			} else {
				return 0
			}
		}
	} else {
		return 0
	}

	return 1
}

func (p Puzzle) checkDiagUpLeft(i int) int {
	starting_row := i / (p.height)
	starting_col := i % (p.height)
	if starting_row >= (len(p.word)-1) && starting_col >= (len(p.word)-1) {
		for j, char := range p.word {
			row := starting_row - j
			next_col := starting_col - j
			next_index := row*(p.height) + next_col
			found_char := rune(p.characters[next_index])
			if char == found_char {
				continue
			} else {
				return 0
			}
		}
	} else {
		return 0
	}

	return 1
}

func main() {
	bytes, err := os.ReadFile("input_data/puzzle.txt")
	check(err)
	puzzle := Puzzle{characters: string(bytes), word: "XMAS"}
	puzzle.calculateSize()
	fmt.Println(puzzle.countWords())
}
