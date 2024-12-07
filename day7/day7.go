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
	var sum int64
	for _, eq := range eqs {
		if eval(eq) {
			fmt.Println("Found", eq)
			sum += int64(eq.result)
		}
	}
	fmt.Printf("Sum: %d\n", sum)
}

func eval(eq Eq) bool {
	return TryPlus(eq.operands[1:], eq.operands[0], eq.result) ||
		TryTimes(eq.operands[1:], eq.operands[0], eq.result)
}

func TryPlus(operands []int, acc, target int) bool {
	// fmt.Printf("TryPlus(%v, %d, %d)\n", operands, acc, target)
	if len(operands) == 0 {
		return acc == target
	}
	acc += operands[0]
	return TryPlus(operands[1:], acc, target) || TryTimes(operands[1:], acc, target)
}

func TryTimes(operands []int, acc, target int) bool {
	// fmt.Printf("TryPlus(%v, %d, %d)\n", operands, acc, target)
	if len(operands) == 0 {
		return acc == target
	}
	acc *= operands[0]
	return TryPlus(operands[1:], acc, target) || TryTimes(operands[1:], acc, target)
}
