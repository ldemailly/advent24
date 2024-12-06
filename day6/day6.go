package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	north = iota
	east
	south
	west
	maxdir
)

type Guard struct {
	x, y      int
	direction int
}

func (g *Guard) TurnRight() {
	g.direction = (g.direction + 1) % maxdir
}

func (g *Guard) Move(backward bool) {
	delta := 1
	if backward {
		delta = -1
	}
	switch g.direction {
	case north:
		g.y = g.y - delta
	case east:
		g.x = g.x + delta
	case south:
		g.y = g.y + delta
	case west:
		g.x = g.x - delta
	}
}

func (g Guard) Print() string {
	switch g.direction {
	case north:
		return "^"
	case east:
		return ">"
	case south:
		return "v"
	case west:
		return "<"
	default:
		return "?"
	}
}

const (
	empty = iota
	obstacle
	visited
)

// happens to be square
type Map struct {
	m       [][]int
	l       int
	visited int
	startX  int
	startY  int
	steps   int
	loop    bool
}

func (m *Map) Outside(g *Guard) bool {
	return g.x < 0 || g.y < 0 || g.x >= m.l || g.y >= m.l
}

func (m *Map) Next(g *Guard) bool {
	m.steps++
	if m.m[g.y][g.x] == empty {
		m.visited++
	} else {
		if m.steps > m.l*m.l {
			m.loop = true
			return false
		}
	}
	m.m[g.y][g.x] = visited
	g.Move(false)
	if m.Outside(g) {
		return false // done
	}
	if m.m[g.y][g.x] == obstacle {
		g.Move(true) // backward/undo
		g.TurnRight()
	}
	return true
}

func (m *Map) Print(g Guard) string {
	var b strings.Builder
	for y := 0; y < m.l; y++ {
		for x := 0; x < m.l; x++ {
			if g.x == x && g.y == y {
				b.WriteString(g.Print())
				continue
			}
			switch m.m[y][x] {
			case empty:
				b.WriteString(".")
			case obstacle:
				b.WriteString("#")
			case visited:
				b.WriteString("*")
			default:
				log.Fatalf("Invalid state %q at %d,%d", m.m[y][x], x, y)
			}
		}
		b.WriteString("\n")
	}
	return b.String()
}

// copy without visited
func (m *Map) CleanClone() Map {
	var m2 Map
	m2.l = m.l
	m2.startX = m.startX
	m2.startY = m.startY
	m2.steps = 0
	m2.visited = 0
	m2.loop = false
	for y := 0; y < m.l; y++ {
		var row []int
		for x := 0; x < m.l; x++ {
			v := m.m[y][x]
			if v != obstacle {
				v = empty
			}
			row = append(row, v)
		}
		m2.m = append(m2.m, row)
	}
	return m2
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	var m Map
	var guard Guard
	y := 0
	for s.Scan() {
		line := s.Text()
		elems := strings.Split(line, "")
		var row []int
		for x, e := range elems {
			switch e {
			case "#":
				row = append(row, obstacle)
			case ".":
				row = append(row, empty)
			case "^":
				row = append(row, empty)
				guard.x = x
				guard.y = y
				fmt.Println("Guard at", x, y)
			default:
				log.Fatalf("Invalid char %q", e)
			}
		}
		m.m = append(m.m, row)
		y++
	}
	m.l = len(m.m)
	if len(m.m[0]) != m.l {
		log.Fatalf("Invalid matrix %d vs %d", len(m.m[0]), m.l)
	}
	initialGuard := guard
	m.startX = guard.x
	m.startY = guard.y
	for m.Next(&guard) {
	}
	fmt.Println(m.Print(guard))
	fmt.Println(m.visited)
	loops := 0
	for y := 0; y < m.l; y++ {
		for x := 0; x < m.l; x++ {
			if x == m.startX && y == m.startY {
				continue
			}
			if m.m[y][x] == visited {
				m2 := m.CleanClone()
				m2.m[y][x] = obstacle
				guard = initialGuard
				for m2.Next(&guard) {
				}
				if m2.loop {
					fmt.Println("Loop by adding obstacle at", x, y)
					loops++
				}
			}
		}
	}
	fmt.Println(loops)
}
