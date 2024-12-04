package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

const n = 140 // 10 for test, 140 for input
var grid = make([][]string, n)

func main() {
	readInput("input.txt")
	count := countOccurrences()
	log.Printf("Solution 1: %d\n", count)
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

func countOccurrences() int {
	count := 0
	directions := []struct{ dx, dy int }{
		{1, 0},   // right
		{-1, 0},  // left
		{0, 1},   // down
		{0, -1},  // up
		{-1, 1},  // bottom left
		{-1, -1}, // top left
		{1, 1},   // bottom right
		{1, -1},  // top right
	}

	for y, row := range grid {
		for x, char := range row {
			if char != "X" {
				continue
			}
			for _, dir := range directions {
				if matchesPattern(x, y, dir.dx, dir.dy) {
					count++
				}
			}
		}
	}

	return count
}

func matchesPattern(x, y, dx, dy int) bool {
	if !isValidPosition(x+3*dx, y+3*dy) {
		return false
	}
	return grid[y+dy][x+dx] == "M" &&
		grid[y+2*dy][x+2*dx] == "A" &&
		grid[y+3*dy][x+3*dx] == "S"
}

func isValidPosition(x, y int) bool {
	return x >= 0 && x < n && y >= 0 && y < n
}
