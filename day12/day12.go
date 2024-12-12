package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"fortio.org/sets"
)

type Coord struct {
	x, y int
}

type Region struct {
	id        byte
	coords    sets.Set[Coord]
	perimeter int
}

type Map struct {
	plots   [][]byte
	seen    sets.Set[Coord]
	inner   sets.Set[Coord]
	regions []*Region
}

func (m *Map) Regions() {
	for y, row := range m.plots {
		for x, c := range row {
			co := Coord{x, y}
			if m.seen.Has(co) {
				continue
			}
			r := &Region{id: c, coords: sets.New[Coord]()}
			m.regions = append(m.regions, r)
			m.ExpandRegion(x, y, c, r)
		}
	}
}

func (m *Map) Edges(x, y int) int {
	edges := 0
	c := m.plots[y][x]
	if x == 0 || m.plots[y][x-1] != c {
		edges++
	}
	if y == 0 || m.plots[y-1][x] != c {
		edges++
	}
	if x == len(m.plots[0])-1 || m.plots[y][x+1] != c {
		edges++
	}
	if y == len(m.plots)-1 || m.plots[y+1][x] != c {
		edges++
	}
	return edges
}

func (m *Map) ExpandRegion(x, y int, c byte, r *Region) {
	if x < 0 || y < 0 || x >= len(m.plots) || y >= len(m.plots[0]) {
		return
	}
	if m.plots[y][x] != c {
		return
	}
	co := Coord{x, y}
	if m.seen.Has(co) {
		return
	}
	r.coords.Add(co)
	m.seen.Add(co)
	m.ExpandRegion(x+1, y, c, r)
	m.ExpandRegion(x, y+1, c, r)
	m.ExpandRegion(x, y-1, c, r)
	m.ExpandRegion(x-1, y, c, r)
	edges := m.Edges(x, y)
	r.perimeter += edges
	if edges == 0 {
		m.inner.Add(co)
	}
}

func Color(c, what byte) string {
	idx := int(c) - int('A')
	if idx < 0 || idx >= 26 {
		log.Fatalf("Invalid color %c %d", c, c)
	}
	color := 18 + idx*216/26
	return fmt.Sprintf("\033[48;5;%dm%c\033[0m", color, what)
}

func (m *Map) String() string {
	sum := 0
	var b strings.Builder
	for y, row := range m.plots {
		for x, c := range row {
			what := c
			if m.inner.Has(Coord{x, y}) {
				what = ' '
			}
			b.WriteString(Color(c, what))
		}
		b.WriteByte('\n')
	}
	for _, r := range m.regions {
		s := len(r.coords)
		p := r.perimeter
		cost := s * p
		sum += cost
		b.WriteString(fmt.Sprintf("Region %c: s %d p %d = %d\n", r.id, s, p, cost))
	}
	b.WriteString(fmt.Sprintf("Total: %d\n", sum))
	return b.String()
}

func main() {
	// read input list
	data, _ := io.ReadAll(os.Stdin)
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	// Part1
	m := Map{plots: make([][]byte, len(lines)), seen: sets.New[Coord](), inner: sets.New[Coord]()}
	for i, line := range lines {
		m.plots[i] = []byte(line)
	}
	m.Regions()
	fmt.Println(m.String())
}
