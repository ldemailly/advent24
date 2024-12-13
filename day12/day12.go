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

type Vector struct {
	direction  Coord
	start, end Coord
}

type Region struct {
	id        byte
	coords    sets.Set[Coord]
	perimeter int
	vectors   [4][]Vector
	edges     int
}

type Map struct {
	plots   [][]byte
	seen    sets.Set[Coord]
	inner   sets.Set[Coord]
	regions []*Region
}

func (r *Region) AddVector(dirID int, x, y int) {
	// extend vector or new one
	if len(r.vectors[dirID]) == 0 {
		r.vectors[dirID] = append(r.vectors[dirID], Vector{start: Coord{x, y}, direction: directions[dirID]})
		r.edges++
		return
	}
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

const (
	westID = iota
	southID
	northID
	eastID
)

var (
	west  = Coord{-1, 0}
	south = Coord{0, 1}
	north = Coord{0, -1}
	east  = Coord{1, 0}
)

// West, South, North, East // optimized for fill order.
var directions = []Coord{west, south, north, east}

func (c Coord) SwapAxis() Coord {
	return Coord{c.y, c.x}
}

func (c Coord) SwapDirection() Coord {
	return Coord{-c.x, -c.y}
}

func (m *Map) IsEdge(c byte, x, y int, direction Coord) bool {
	xx := x + direction.x
	yy := y + direction.y
	return m.Outside(xx, yy) || m.plots[yy][xx] != c
}

func (m *Map) Edges(r *Region, x, y int) int {
	edges := 0
	c := m.plots[y][x]
	for did, d := range directions {
		if m.IsEdge(c, x, y, d) {
			r.AddVector(did, x, y)
			edges++
		}
	}
	return edges
}

func (m *Map) Outside(x, y int) bool {
	return x < 0 || y < 0 || x >= len(m.plots) || y >= len(m.plots[0])
}

func (m *Map) ExpandRegion(x, y int, c byte, r *Region) {
	if m.Outside(x, y) {
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
	for _, d := range directions {
		m.ExpandRegion(x+d.x, y+d.y, c, r)
	}
	perim := m.Edges(r, x, y)
	r.perimeter += perim
	if perim == 0 {
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
	sum1 := 0
	sum2 := 0
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
		cost1 := s * p
		sum1 += cost1
		e := r.edges
		cost2 := s * e
		sum2 += cost2
		b.WriteString(fmt.Sprintf("Region %c: s %d p %d = %d - e %d = %d\n", r.id, s, p, cost1, e, cost2))
	}
	b.WriteString(fmt.Sprintf("Total with perimeter : %d\n", sum1))
	b.WriteString(fmt.Sprintf("Total with edges     : %d\n", sum2))
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
