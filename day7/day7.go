package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Eq struct {
	result   int
	operands []int
}

func main() {
	var eqs []Eq
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		var eq Eq
		line := s.Text()
		lr := strings.Split(line, ": ")
		if len(lr) != 2 {
			log.Fatalf("Invalid line %q", line)
		}
		var err error
		eq.result, err = strconv.Atoi(lr[0])
		if err != nil {
			log.Fatalf("Invalid result %q: %v", lr[0], err)
		}
		elems := strings.Split(lr[1], " ")
		for _, e := range elems {
			n, err := strconv.Atoi(e)
			if err != nil {
				log.Fatalf("Invalid operand %q %q: %v", line, e, err)
			}
			eq.operands = append(eq.operands, n)
		}
		eqs = append(eqs, eq)
	}
	fmt.Println(eqs)
}
