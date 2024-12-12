package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

var count1 = 0
var count2 = 0

var garden = make([][]string, 0)
var visited = make(map[coord]bool)

type coord struct {
	x int
	y int
}

func main() {
	readInput("input.txt")
	countScores()
	log.Printf("Solution 1: %d", count1)
	log.Printf("Solution 2: %d", count2)
}

func readInput(filename string) {
	file, _ := os.Open(filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	y := 0
	for scanner.Scan() {
		garden = append(garden, make([]string, 0))
		for _, c := range strings.Split(scanner.Text(), "") {
			garden[y] = append(garden[y], c)
		}
		y++
	}
}

func countScores() {
	for y, row := range garden {
		for x, c := range row {
			a, p := countPlot(coord{x: x, y: y}, c)
			count1 += a * p
		}
	}
}

func countPlot(pos coord, c string) (int, int) {
	if visited[pos] || garden[pos.y][pos.x] != c {
		return 0, 0
	}

	visited[pos] = true
	area := 1
	perimeter := 0

	directions := []coord{
		{y: pos.y - 1, x: pos.x}, // up
		{y: pos.y + 1, x: pos.x}, // down
		{y: pos.y, x: pos.x - 1}, // left
		{y: pos.y, x: pos.x + 1}, // right
	}

	for _, next := range directions {
		if isOutOfBounds(next) || garden[next.y][next.x] != c {
			perimeter++
		} else {
			a, p := countPlot(next, c)
			area += a
			perimeter += p
		}
	}

	return area, perimeter
}

func isOutOfBounds(next coord) bool {
	return next.x < 0 || next.x >= len(garden[0]) || next.y < 0 || next.y >= len(garden)
}
