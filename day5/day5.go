package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"fortio.org/sets"
)

type Rule struct {
	Before, After int
}

type Data struct {
	Row []int // raw row
	Set sets.Set[int]
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	var rules []Rule
	// fmt.Println("Rules:")
	for s.Scan() {
		line := s.Text()
		if line == "" {
			break // 2nd section
		}
		var r Rule
		_, err := fmt.Sscanf(line, "%d|%d", &r.Before, &r.After)
		if err != nil {
			panic(err)
		}
		rules = append(rules, r)
	}
	// fmt.Println(rules)
	// fmt.Println("Data:")
	var data []Data
	for s.Scan() {
		line := s.Text()
		elems := strings.Split(line, ",")
		var row Data
		for _, e := range elems {
			v, err := strconv.Atoi(e)
			if err != nil {
				panic(err)
			}
			row.Row = append(row.Row, v)
		}
		row.Set = sets.FromSlice(row.Row)
		data = append(data, row)
	}
	// fmt.Println(data)

	// Part 1
	for i, row := range data {
		rowOk := true
		for j, r := range rules {
			if !(row.Set.Has(r.Before) && row.Set.Has(r.After)) {
				// need both to be there for the rule to matter
				continue
			}
			i1 := slices.Index(row.Row, r.Before)
			i2 := slices.Index(row.Row, r.After)
			if i1 > i2 {
				fmt.Printf("Row %d elimited by rule %d: %v: %d is before %d\n", i+1, j+1, row.Row, r.Before, r.After)
				rowOk = false
				break
			}
		}
		if rowOk {
			fmt.Printf("Row %d is ok: %v\n", i+1, row.Row)
		}
	}
}
