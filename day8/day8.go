package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"fortio.org/sets"
)

type Point struct {
	x, y int
}

func (p Point) Anti(o Point) Point {
	return Point{2*p.x - o.x, 2*p.y - o.y}
}

func (p Point) Ok(max int) bool {
	return p.x >= 0 && p.y >= 0 && p.x < max && p.y < max
}

func AntiPoints(res []Point, ps []Point, max int) []Point {
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
		res = AntiPoints(res, ps[1:], max)
	}
	return res
}

func main() {
	inp, _ := io.ReadAll(os.Stdin)
	lines := strings.Split(strings.TrimSpace(string(inp)), "\n")
	height := len(lines)
	width := len(lines[0])
	if height != width {
		log.Fatalf("Non-square map: %d x %d", height, width)
	}
	max := height
	m := make(map[byte][]Point)
	for y := range height {
		l := lines[y]
		for x := range width {
			if l[x] == '.' {
				continue
			}
			m[l[x]] = append(m[l[x]], Point{x, y})
		}
	}
	res := sets.New[Point]()
	for k, v := range m {
		ap := AntiPoints(nil, v, max)
		fmt.Printf("%c: %d %v\n", k, len(ap), AntiPoints(nil, v, max))
		res.Add(ap...)
	}
	fmt.Println("Total:", len(res))
}
