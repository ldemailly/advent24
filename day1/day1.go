// Day one of advent of code 2024
// https://adventofcode.com/2024/day/1
// Seems to be about sorting 2 lists and subtracting and summing them.

package main

import (
	"fmt"
	"io"
	"log"
	"slices"
)

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	l1, l2 := readInput()
	if len(l1) != len(l2) {
		panic("Lists are not of the same length")
	}
	slices.Sort(l1)
	slices.Sort(l2)
	sum := 0
	for i := range len(l1) {
		// Unlike what is specified the lists can have elements at times before and after the other list
		/*		if l2[i] < l1[i] {
					log.Fatalf("At %d list2=%d is not greater than list1=%d", i, l2[i], l1[i])
				}
		*/
		sum += abs(l1[i] - l2[i]) // it's unclear if abs was needed or not, but it is to get the right answer
	}
	fmt.Println(sum)
}

func readInput() ([]int, []int) {
	l1 := []int{}
	l2 := []int{}
	lines := 0
	for {
		var n1, n2 int
		lines++
		n, err := fmt.Scanf("%d %d\n", &n1, &n2)
		if err == io.EOF {
			break
		}
		if n != 2 {
			log.Fatalf("Expected 2 numbers, got %d", n)
		}
		if err != nil {
			panic(err)
		}
		l1 = append(l1, n1)
		l2 = append(l2, n2)
	}
	return l1, l2
}
