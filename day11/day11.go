package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Stone struct {
	n    int
	nStr string
}

func evenDigits(n Stone) bool {
	return len(n.nStr)%2 == 0
}

func leftRight(n Stone) []Stone {
	left := n.nStr[:len(n.nStr)/2]
	right := n.nStr[len(n.nStr)/2:]
	return []Stone{StoneFromStr(left), StoneFromStr(right)}
}

var One = Stone{1, "1"}

func StoneFromNum(n int) Stone {
	return Stone{n, strconv.Itoa(n)}
}

func StoneFromStr(s string) Stone {
	n, _ := strconv.Atoi(s)
	return Stone{n, strings.TrimLeft(s, "0")}
}

func ApplyRules(inp []Stone) []Stone {
	r := make([]Stone, 0, 2*len(inp)) // worst case is all gets doubled.
	for _, v := range inp {
		switch {
		case v.n == 0:
			r = append(r, One)
		case evenDigits(v):
			r = append(r, leftRight(v)...)
		default:
			r = append(r, StoneFromNum(2024*v.n))
		}
	}
	return r
}

func PrintStones(stones []Stone) {
	for _, s := range stones {
		fmt.Printf("%d ", s.n)
	}
	fmt.Println()
}

func main() {
	// read input list
	data, _ := io.ReadAll(os.Stdin)
	inp := strings.TrimSpace(string(data))
	inpList := strings.Split(inp, " ")
	// Split into list of ints.
	stones := make([]Stone, 0, len(inpList))
	for _, n := range inpList {
		stones = append(stones, StoneFromStr(n))
	}
	for n := range 25 {
		fmt.Println(n, stones)
		PrintStones(stones)
		stones = ApplyRules(stones)
	}
	// Part1
	PrintStones(stones)
	fmt.Println("Part 1", len(stones))
}
