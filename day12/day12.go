package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"fortio.org/sets"
)

type Coord struct {
	x, y int
}

type Region struct {
	id     byte
	coords sets.Set[Coord]
}

type Map struct {
	plots   [][]byte
	seen    sets.Set[Coord]
	regions []Region
	// perim   int
}

func (m *Map) Regions() {
	for y, row := range m.plots {
		for x, c := range row {
			co := Coord{x, y}
			if m.seen.Has(co) {
				continue
			}
			r := Region{id: c, coords: sets.New[Coord]()}
			m.regions = append(m.regions, r)
			m.ExpandRegion(x, y, c, r)
		}
	}
}

func (m *Map) ExpandRegion(x, y int, c byte, r Region) {
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
}

func (m *Map) String() string {
	var b strings.Builder
	for _, r := range m.regions {
		b.WriteString(fmt.Sprintf("Region %c: %d\n", r.id, len(r.coords)))
	}
	return b.String()
}

func main() {
	// read input list
	data, _ := io.ReadAll(os.Stdin)
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	// Part1
	m := Map{plots: make([][]byte, len(lines)), seen: sets.New[Coord]()}
	for i, line := range lines {
		m.plots[i] = []byte(line)
	}
	m.Regions()
	fmt.Println(m.String())
}
