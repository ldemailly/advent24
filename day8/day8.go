package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"fortio.org/sets"
)

type Map struct {
	m   map[byte][]Point
	max int
	buf []byte
}

func (m Map) Set(c byte, p Point) {
	m.buf[p.y*(m.max+1)+p.x] = c
}

func (m Map) String() string {
	m.buf = make([]byte, m.max*(m.max+1))
	for y := range m.max {
		for x := range m.max {
			m.buf[y*(m.max+1)+x] = '.'
		}
		m.buf[y*(m.max+1)+m.max] = '\n'
	}
	for k, v := range m.m {
		for _, p := range v {
			m.Set(k, p)
		}
	}
	return string(m.buf)
}

type Point struct {
	x, y int
}

func (p Point) Anti(o Point) Point {
	return Point{2*p.x - o.x, 2*p.y - o.y}
}

func (p Point) Ok(max int) bool {
	return p.x >= 0 && p.y >= 0 && p.x < max && p.y < max
}

func AntiPoints1(res []Point, ps []Point, max int) []Point {
	p0 := ps[0]
	for _, p := range ps[1:] {
		a1 := p0.Anti(p)
		if a1.Ok(max) {
			res = append(res, a1)
		}
		a2 := p.Anti(p0)
		if a2.Ok(max) {
			res = append(res, a2)
		}
	}
	if len(ps) > 2 {
		res = AntiPoints1(res, ps[1:], max)
	}
	return res
}

func Anti2(from, to Point, max int) (res []Point) {
	for {
		next := from.Anti(to)
		if !next.Ok(max) {
			return res
		}
		res = append(res, next)
		to = from
		from = next
	}
}

func AntiPoints2(res []Point, ps []Point, max int) []Point {
	p0 := ps[0]
	for _, p := range ps[1:] {
		res = append(res, Anti2(p0, p, max)...)
		res = append(res, Anti2(p, p0, max)...)
	}
	if len(ps) > 2 {
		res = AntiPoints2(res, ps[1:], max)
	}
	return res
}

func main() {
	debug := false
	flag.BoolVar(&debug, "debug", false, "debug mode: only for uppercase input")
	flag.Parse()
	inp, _ := io.ReadAll(os.Stdin)
	lines := strings.Split(strings.TrimSpace(string(inp)), "\n")
	height := len(lines)
	width := len(lines[0])
	if height != width {
		log.Fatalf("Non-square map: %d x %d", height, width)
	}
	m := Map{m: make(map[byte][]Point), max: height}
	for y := range height {
		l := lines[y]
		for x := range width {
			if l[x] == '.' {
				continue
			}
			m.m[l[x]] = append(m.m[l[x]], Point{x, y})
		}
	}
	if debug {
		fmt.Println(m)
	}
	res := sets.New[Point]()
	for k, v := range m.m {
		ap := AntiPoints1(nil, v, m.max)
		res.Add(ap...)
		if debug {
			fmt.Printf("%c: %d %v\n", k, len(ap), ap)
		}
	}
	fmt.Println("Part1:", len(res))
	res.Clear()
	for k, v := range m.m {
		ap := AntiPoints2(v, v, m.max)
		res.Add(ap...)
		if !debug {
			continue
		}
		fmt.Printf("%c: %d unique %d raw %v\n", k, len(ap), len(sets.FromSlice(ap)), ap)
		lowerCase := k + 0x20
		if _, found := m.m[lowerCase]; found {
			log.Fatalf("Can't use -debug because found '%c' already preexisting when adding '%c'", lowerCase, k)
			continue
		}
		m.m[lowerCase] = ap
		fmt.Println(m)
		delete(m.m, lowerCase)
	}
	fmt.Println("Part2:", len(res))
}
