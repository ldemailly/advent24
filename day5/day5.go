package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rule struct {
	Before, After int
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	var rules []Rule
	fmt.Println("Rules:")
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
	fmt.Println(rules)
	fmt.Println("Data:")
	var data [][]int
	for s.Scan() {
		line := s.Text()
		elems := strings.Split(line, ",")
		var row []int
		for _, e := range elems {
			v, err := strconv.Atoi(e)
			if err != nil {
				panic(err)
			}
			row = append(row, v)
		}
		data = append(data, row)
	}
	fmt.Println(data)
}
