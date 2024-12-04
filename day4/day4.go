package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// X M A S so 3 after the X

func HasXMASright(matrix [][]string, i, j int) int {
	if j+3 >= len(matrix[i]) {
		return 0
	}
	if matrix[i][j+1] == "M" && matrix[i][j+2] == "A" && matrix[i][j+3] == "S" {
		return 1
	}
	return 0
}

func HasXMASleft(matrix [][]string, i, j int) int {
	if j < 3 {
		return 0
	}
	if matrix[i][j-1] == "M" && matrix[i][j-2] == "A" && matrix[i][j-3] == "S" {
		return 1
	}
	return 0
}

func HasXMASup(matrix [][]string, i, j int) int {
	if i < 3 {
		return 0
	}
	if matrix[i-1][j] == "M" && matrix[i-2][j] == "A" && matrix[i-3][j] == "S" {
		return 1
	}
	return 0
}

func HasXMASdown(matrix [][]string, i, j int) int {
	if i+3 >= len(matrix) {
		return 0
	}
	if matrix[i+1][j] == "M" && matrix[i+2][j] == "A" && matrix[i+3][j] == "S" {
		return 1
	}
	return 0
}

func HasXMASDiagUpRight(matrix [][]string, i, j int) int {
	if i < 3 || j+3 >= len(matrix[i]) {
		return 0
	}
	if matrix[i-1][j+1] == "M" && matrix[i-2][j+2] == "A" && matrix[i-3][j+3] == "S" {
		return 1
	}
	return 0
}

func HasXMASDiagUpLeft(matrix [][]string, i, j int) int {
	if i < 3 || j < 3 {
		return 0
	}
	if matrix[i-1][j-1] == "M" && matrix[i-2][j-2] == "A" && matrix[i-3][j-3] == "S" {
		return 1
	}
	return 0
}

func HasXMASDiagDownRight(matrix [][]string, i, j int) int {
	if i+3 >= len(matrix) || j+3 >= len(matrix[i]) {
		return 0
	}
	if matrix[i+1][j+1] == "M" && matrix[i+2][j+2] == "A" && matrix[i+3][j+3] == "S" {
		return 1
	}
	return 0
}

func HasXMASDiagDownLeft(matrix [][]string, i, j int) int {
	if i+3 >= len(matrix) || j < 3 {
		return 0
	}
	if matrix[i+1][j-1] == "M" && matrix[i+2][j-2] == "A" && matrix[i+3][j-3] == "S" {
		return 1
	}
	return 0
}

func main() {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	dataStr := string(data)
	lines := strings.Split(strings.TrimSpace(dataStr), "\n")
	width := len(lines[0])
	height := len(lines)
	fmt.Printf("%d x %d\n", width, height)
	matrix := make([][]string, height)
	for i, line := range lines {
		// log.Printf("At %d: %s", i, line)
		matrix[i] = strings.Split(line, "")
		if len(matrix[i]) != width {
			log.Fatalf("At %d: Invalid matrix %d vs %d", i, len(matrix[i]), width)
		}
	}
	// Part 1
	sum := 0
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] != "X" {
				continue
			}
			sum += HasXMASright(matrix, i, j)
			sum += HasXMASleft(matrix, i, j)
			sum += HasXMASup(matrix, i, j)
			sum += HasXMASdown(matrix, i, j)
			sum += HasXMASDiagUpRight(matrix, i, j)
			sum += HasXMASDiagUpLeft(matrix, i, j)
			sum += HasXMASDiagDownRight(matrix, i, j)
			sum += HasXMASDiagDownLeft(matrix, i, j)
		}
	}
	fmt.Println(sum)
}
