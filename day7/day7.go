package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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
		if evalPart1(eq) {
			// fmt.Println("Found", eq)
			sum += int64(eq.result)
		}
	}
	fmt.Println("Part1:", sum)
	sum = 0
	for _, eq := range eqs {
		if evalPart2(eq) {
			// fmt.Println("Found", eq)
			sum += int64(eq.result)
		}
	}
	fmt.Println("Part2:", sum)
}

func evalPart1(eq Eq) bool {
	return TryPlus1(eq.operands[1:], eq.operands[0], eq.result) ||
		TryTimes1(eq.operands[1:], eq.operands[0], eq.result)
}

func TryPlus1(operands []int, acc, target int) bool {
	// fmt.Printf("TryPlus(%v, %d, %d)\n", operands, acc, target)
	if len(operands) == 0 {
		return acc == target
	}
	acc += operands[0]
	return TryPlus1(operands[1:], acc, target) || TryTimes1(operands[1:], acc, target)
}

func TryTimes1(operands []int, acc, target int) bool {
	// fmt.Printf("TryPlus(%v, %d, %d)\n", operands, acc, target)
	if len(operands) == 0 {
		return acc == target
	}
	acc *= operands[0]
	return TryPlus1(operands[1:], acc, target) || TryTimes1(operands[1:], acc, target)
}

// --- could change part1 too but... it's late already.

func evalPart2(eq Eq) bool {
	return TryPlus2(eq.operands[1:], eq.operands[0], eq.result) ||
		TryTimes2(eq.operands[1:], eq.operands[0], eq.result) ||
		TryConcat(eq.operands[1:], eq.operands[0], eq.result)
}

func TryPlus2(operands []int, acc, target int) bool {
	// fmt.Printf("TryPlus(%v, %d, %d)\n", operands, acc, target)
	acc += operands[0]
	if len(operands) == 1 {
		return acc == target
	}
	return TryPlus2(operands[1:], acc, target) || TryTimes2(operands[1:], acc, target) || TryConcat(operands[1:], acc, target)
}

func TryTimes2(operands []int, acc, target int) bool {
	// fmt.Printf("TryTimes(%v, %d, %d)\n", operands, acc, target)
	acc *= operands[0]
	if len(operands) == 1 {
		return acc == target
	}
	return TryPlus2(operands[1:], acc, target) || TryTimes2(operands[1:], acc, target) || TryConcat(operands[1:], acc, target)
}

func TryConcat(operands []int, acc, target int) bool {
	// fmt.Printf("TryConcat(%v, %d, %d)", operands, acc, target)
	acc = IntConcat2(acc, operands[0])
	// fmt.Printf(" -> %d\n", acc)
	if len(operands) == 1 {
		return acc == target
	}
	return TryPlus2(operands[1:], acc, target) || TryTimes2(operands[1:], acc, target) || TryConcat(operands[1:], acc, target)
}

// initial float math version:
func IntConcat1(a, b int) int {
	numDigits := int(math.Floor(math.Log10(float64(b))) + 1)
	return a*int(math.Pow10(numDigits)) + b
}

// Pure integer version:
func IntConcat2(a, b int) int {
	// Calculate the number of digits in `b`
	digits := 1
	temp := b
	for temp >= 10 {
		temp /= 10
		digits++
	}

	// Calculate the factor (10^digits) to shift `a` appropriately
	factor := 1
	for i := 0; i < digits; i++ {
		factor *= 10
	}

	// Concatenate a and b
	return a*factor + b
}
