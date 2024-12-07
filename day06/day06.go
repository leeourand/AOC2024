package main

import (
	"fmt"
	"os"
)

type Guard struct {
	position *Position
	velocity *Velocity
	history  map[int]map[int]bool
}

type Position struct {
	x int
	y int
}

type Velocity struct {
	x int
	y int
}

type Map struct {
	obstacles map[int]map[int]bool
	guard     *Guard
	width     int
	height    int
}

func check(e error) {
	if e != nil {
		fmt.Printf("Something went wrong: %q", e)
		panic(e)
	}
}

func main() {
	bytes, err := os.ReadFile("input_data/puzzle.txt")
	check(err)

	m := buildMap(bytes)
	m.RunSim()
	fmt.Println(m.guard.HistoryCount())
}

func (m *Map) RunSim() {
	for true {
		if m.obstacleBlockingGuard() {
			m.guard.Rotate()
		} else {
			if m.guard.position.x <= m.width && m.guard.position.x >= 0 && m.guard.position.y <= m.height && m.guard.position.y >= 0 {
				m.guard.MoveGuard()
			} else {
				break
			}
		}
	}
}

func (m Map) obstacleBlockingGuard() bool {
	if _, exists := m.obstacles[m.guard.nextPosition().y]; exists {
		if _, blocked := m.obstacles[m.guard.nextPosition().y][m.guard.nextPosition().x]; blocked {
			return true
		}
	}
	return false
}

func (g Guard) nextPosition() *Position {
	var newPosition Position
	newPosition.x = g.position.x + g.velocity.x
	newPosition.y = g.position.y + g.velocity.y
	return &newPosition
}

func (g *Guard) MoveGuard() {
	g.RecordPosition()
	g.position = g.nextPosition()
}

func (g *Guard) Rotate() {
	prevX := g.velocity.x
	prevY := g.velocity.y
	g.velocity.x = prevY * -1
	g.velocity.y = prevX
}

func (g *Guard) RecordPosition() {
	if _, exists := g.history[g.position.y]; !exists {
		g.history[g.position.y] = make(map[int]bool)
	}
	g.history[g.position.y][g.position.x] = true
}

func (g Guard) HistoryCount() int {
	count := 0
	for _, h := range g.history {
		for range h {
			count++
		}
	}
	return count
}

func buildMap(bytes []byte) *Map {
	row := 0
	col := 0
	var m Map
	m.obstacles = make(map[int]map[int]bool)

	for i, char := range bytes {
		if char == '#' {
			if _, exists := m.obstacles[row]; !exists {
				m.obstacles[row] = make(map[int]bool)
			}
			m.obstacles[row][col] = true
		}

		if char == '\n' {
			if m.width == 0 {
				m.width = i
			}
			row++
			m.height = row
			col = -1
		}

		if char == '^' {
			m.guard = &Guard{position: &Position{x: col, y: row}, velocity: &Velocity{x: 0, y: -1}}
			m.guard.history = make(map[int]map[int]bool)
		}
		col++
	}
	return &m
}
