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

const (
	GreenBackground = "\033[42m"
	// copied from 	fortio.org/terminal/ansipixels
	Reset     = "\033[0m"
	Green     = "\033[32m"
	FullPixel = '█'
	BluePixel = "\033[34m" + string(FullPixel) + Reset
	RedPixel  = "\033[31m" + string(FullPixel) + Reset
)

func (g Guard) Print() string {
	switch g.direction {
	case north:
		return GreenBackground + "⇑" + Reset
	case east:
		return GreenBackground + "⇒" + Reset
	case south:
		return GreenBackground + "⇓" + Reset
	case west:
		return GreenBackground + "⇐" + Reset
	default:
		panic("Invalid direction")
	}
}

const (
	empty = iota
	obstacle
	visited // + direction
)

// happens to be square
type Map struct {
	m          [][]int
	l          int
	visited    int
	startX     int
	startY     int
	steps      int
	loop       bool
	newX, newY int
}

func (m *Map) Outside(g *Guard) bool {
	return g.x < 0 || g.y < 0 || g.x >= m.l || g.y >= m.l
}

func (m *Map) Next(g *Guard) bool {
	m.steps++
	if m.m[g.y][g.x] == empty {
		m.visited++
		m.m[g.y][g.x] = visited + g.direction
	} else {
		if m.m[g.y][g.x] == visited+g.direction {
			// fmt.Println("Loop at", g.x, g.y, "after", m.steps, "steps")
			m.loop = true
			return false
		}
	}
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

func (m *Map) SetNew(x, y int) {
	m.newX = x
	m.newY = y
	m.m[y][x] = obstacle
}

func (m *Map) Print(g Guard) string {
	var b strings.Builder
	b.WriteString("\033[H")
	for y := 0; y < m.l; y++ {
		for x := 0; x < m.l; x++ {
			if g.x == x && g.y == y {
				b.WriteString(g.Print())
				continue
			}
			if x == m.startX && y == m.startY {
				b.WriteString(Green + "↑" + Reset)
				continue
			}
			switch m.m[y][x] {
			case empty:
				b.WriteRune(FullPixel)
			case obstacle:
				if x == m.newX && y == m.newY {
					b.WriteString(BluePixel)
				} else {
					b.WriteString(RedPixel)
				}
			case visited + north:
				b.WriteString("↑")
			case visited + east:
				b.WriteString("→")
			case visited + south:
				b.WriteString("↓")
			case visited + west:
				b.WriteString("←")
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
	fmt.Println(m.Print(guard))
	for m.Next(&guard) {
		fmt.Println(m.Print(guard))
	}
	fmt.Println(m.Print(guard))
	fmt.Println(m.visited)
	loops := 0
	for y := 0; y < m.l; y++ {
		for x := 0; x < m.l; x++ {
			if x == m.startX && y == m.startY {
				continue
			}
			if m.m[y][x] >= visited {
				m2 := m.CleanClone()
				m2.SetNew(x, y)
				guard = initialGuard
				for m2.Next(&guard) {
				}
				if m2.loop {
					fmt.Println("Loop by adding obstacle at", x, y)
					fmt.Println(m2.Print(guard))
					loops++
				}
			}
		}
	}
	fmt.Println(loops)
}
