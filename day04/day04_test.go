package main

import (
	"testing"
)

func TestLeftRight(t *testing.T) {
	input := `?????
?????
?XMAS
?????`

	puzzle := Puzzle{characters: input, word: "XMAS"}
	puzzle.calculateSize()
	if puzzle.countWords() != 1 {
		t.Errorf("Expected to find XMAS on left / right but did not")
	}
}

func TestTopDown(t *testing.T) {
	input := `X????
M????
A????
S????`

	puzzle := Puzzle{characters: input, word: "XMAS"}
	puzzle.calculateSize()
	if puzzle.countWords() != 1 {
		t.Errorf("Expected to find XMAS on top / down but did not")
	}
}

func TestBottomUp(t *testing.T) {
	input := `??S??
??A??
??M??
??X??`

	puzzle := Puzzle{characters: input, word: "XMAS"}
	puzzle.calculateSize()
	if puzzle.countWords() != 1 {
		t.Errorf("Expected to find XMAS on bottom / up but did not")
	}
}

func TestRightLeft(t *testing.T) {
	input := `?????
SAMX?
?????
?????`

	puzzle := Puzzle{characters: input, word: "XMAS"}
	puzzle.calculateSize()
	if puzzle.countWords() != 1 {
		t.Errorf("Expected to find XMAS on right / left but did not")
	}
}

func TestDiagRightUp(t *testing.T) {
	input := `???S?
??A??
?M???
X????`

	puzzle := Puzzle{characters: input, word: "XMAS"}
	puzzle.calculateSize()
	if puzzle.checkDiagUpRight(18) != 1 {
		t.Errorf("Expected to find XMAS on diag right/up but did not")
	}
}

func TestDiagRightDown(t *testing.T) {
	input := `X????
?M???
??A??
???S?`

	puzzle := Puzzle{characters: input, word: "XMAS"}
	puzzle.calculateSize()
	if puzzle.countWords() != 1 {
		t.Errorf("Expected to find XMAS on diag right/down but did not")
	}
}

func TestDiagLeftDown(t *testing.T) {
	input := `???X?
??M??
?A???
S????`

	puzzle := Puzzle{characters: input, word: "XMAS"}
	puzzle.calculateSize()
	if puzzle.countWords() != 1 {
		t.Errorf("Expected to find XMAS on diag left/down but did not")
	}
}

func TestDiagLeftUp(t *testing.T) {
	input := `???S?
??A??
?M???
X????`

	puzzle := Puzzle{characters: input, word: "XMAS"}
	puzzle.calculateSize()
	if puzzle.countWords() != 1 {
		t.Errorf("Expected to find XMAS on diag left/up but did not")
	}
}
