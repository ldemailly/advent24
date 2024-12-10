package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var ansi256Gray = []int{
	237, // 58
	239, // 78
	241, // 98
	243, // 118
	245, // 138
	247, // 158
	249, // 178
	251, // 198
	253, // 218
	15,  // 255 white
}

const (
	GreenBackground = "\033[42m"
	// Copied/inspired from my https://pkg.go.dev/fortio.org/terminal/ansipixels
	Reset          = "\033[0m"
	Green          = "\033[32m"
	FullPixel      = '█'
	BluePixel      = "\033[34m" + string(FullPixel) + Reset
	RedPixel       = "\033[31m" + string(FullPixel) + Reset
	ClearScreen    = "\033[2J"
	ClearEndOfLine = "\033[K"
)

func HeightToPixel(h int) string {
	if h < 0 || h >= len(ansi256Gray) {
		log.Fatalf("Invalid height %d", h)
	}
	return fmt.Sprintf("\033[38;5;%dm█", ansi256Gray[h])
}

// happens to be square.
type Map struct {
	heights [][]int
	l       int
}

func (m *Map) Print() string {
	var b strings.Builder
	// b.WriteString("\033[H")
	for y := range m.l {
		line := m.heights[y]
		for x := range m.l {
			b.WriteString(HeightToPixel(line[x]))
		}
		b.WriteString("\n")
	}
	return b.String()
}

var doPrint = true

func main() {
	//	flag.BoolVar(&doPrint, "print", false, "Print the map")
	//	flag.Parse()
	s := bufio.NewScanner(os.Stdin)
	var m Map
	y := 0
	for s.Scan() {
		line := s.Text()
		elems := strings.Split(line, "")
		var row []int
		for _, e := range elems {
			n, err := strconv.Atoi(e)
			if err != nil {
				log.Fatalf("Invalid height %q: %v", e, err)
			}
			row = append(row, n)
		}
		m.heights = append(m.heights, row)
		y++
	}
	m.l = len(m.heights)
	if len(m.heights[0]) != m.l {
		log.Fatalf("Invalid matrix %d vs %d", len(m.heights[0]), m.l)
	}
	if doPrint {
		// clear screen
		// fmt.Print(ClearScreen)
		fmt.Print(m.Print())
	}
}
