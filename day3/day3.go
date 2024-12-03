package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

func main() {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	re := regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)`)
	res := re.FindAllStringSubmatch(string(data), -1)
	sum := 0
	for _, r := range res {
		a, b := r[1], r[2]
		x, err := strconv.Atoi(a)
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(b)
		if err != nil {
			panic(err)
		}
		sum += x * y
	}
	fmt.Println(sum)
}
