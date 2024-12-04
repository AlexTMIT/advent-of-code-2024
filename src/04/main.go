package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

const n = 140 // 10 for test, 140 for input
var grid = make([][]string, n)
var directions = []struct{ dx, dy int }{
	{1, 0},   // right
	{-1, 0},  // left
	{0, 1},   // down
	{0, -1},  // up
	{-1, 1},  // bottom left
	{-1, -1}, // top left
	{1, 1},   // bottom right
	{1, -1},  // top right
}

func main() {
	readInput("input.txt")
	count1, count2 := countOccurrences()
	log.Printf("Solution 1: %d\n", count1) // 2646
	log.Printf("Solution 2: %d\n", count2) // 2000
}

func readInput(filename string) {
	file, _ := os.Open(filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	row := 0
	for scanner.Scan() {
		grid[row] = strings.Split(scanner.Text(), "")
		row++
	}
}

func countOccurrences() (int, int) {
	count1 := 0
	count2 := 0

	for y, row := range grid {
		for x, char := range row {
			if char == "X" {
				for _, dir := range directions {
					if matchesPattern(x, y, dir.dx, dir.dy) {
						count1++
					}
				}
			}

			if char == "A" {
				if isXMASAt(x, y) {
					count2++
				}
			}
		}
	}

	return count1, count2
}

func matchesPattern(x, y, dx, dy int) bool {
	if !isValidPosition(x+3*dx, y+3*dy) {
		return false
	}
	return existsMAS(y, x, dy, dx)
}

func existsMAS(y, x, dy, dx int) bool {
	return grid[y+dy][x+dx] == "M" &&
		grid[y+2*dy][x+2*dx] == "A" &&
		grid[y+3*dy][x+3*dx] == "S"
}

func isValidPosition(x, y int) bool {
	return x >= 0 && x < n && y >= 0 && y < n
}

/*
* All code below is AI generated.
* Unfortunately, I did not have time to solve the problem.
* However, I wanted to see how it could be done by extending
* my code for the learning experience.
* For what it's worth, I probably would have done something
* very similar to this.
 */
func isXMASAt(x, y int) bool {
	// Check first diagonal (top-left to bottom-right)
	diag1 := false
	if isValidPosition(x-1, y-1) && isValidPosition(x+1, y+1) {
		l1 := grid[y-1][x-1]
		l3 := grid[y+1][x+1]
		if isMASSequence(l1, l3) {
			diag1 = true
		}
	}
	// Check second diagonal (top-right to bottom-left)
	diag2 := false
	if isValidPosition(x+1, y-1) && isValidPosition(x-1, y+1) {
		l1 := grid[y-1][x+1]
		l3 := grid[y+1][x-1]
		if isMASSequence(l1, l3) {
			diag2 = true
		}
	}
	return diag1 && diag2
}

func isMASSequence(l1, l3 string) bool {
	return (l1 == "M" && l3 == "S") || (l1 == "S" && l3 == "M")
}
