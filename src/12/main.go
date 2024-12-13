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
	x, y int
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
			set := make(map[coord]bool)
			a, p := countPlot(coord{x: x, y: y}, c, &set)
			count1 += a * p
			count2 += a * countCorners(set, c)
		}
	}
}
func countCorners(set map[coord]bool, c string) (corners int) {
	for p := range set {
		left := set[coord{x: p.x - 1, y: p.y}]
		right := set[coord{x: p.x + 1, y: p.y}]
		up := set[coord{x: p.x, y: p.y - 1}]
		down := set[coord{x: p.x, y: p.y + 1}]
		upLeft := set[coord{x: p.x - 1, y: p.y - 1}]
		upRight := set[coord{x: p.x + 1, y: p.y - 1}]
		downLeft := set[coord{x: p.x - 1, y: p.y + 1}]
		downRight := set[coord{x: p.x + 1, y: p.y + 1}]

		// outer corners
		if !left && !up { // thanks to RazarTuk on reddit for removing one condition and hereby fixing my code haha
			corners = addCorner(p, corners)
		}
		if !left && !down {
			corners = addCorner(p, corners)
		}
		if !right && !up {
			corners = addCorner(p, corners)
		}
		if !right && !down {
			corners = addCorner(p, corners)
		}

		// inner corners
		if left && up && !upLeft {
			corners = addCorner(p, corners)
		}
		if right && up && !upRight {
			corners = addCorner(coord{x: p.x + 1, y: p.y}, corners)
		}
		if left && down && !downLeft {
			corners = addCorner(coord{x: p.x, y: p.y + 1}, corners)
		}
		if right && down && !downRight {
			corners = addCorner(coord{x: p.x + 1, y: p.y + 1}, corners)
		}
	}

	if corners > 0 {
		log.Printf("%s has %d", c, corners)
	}
	return
}

func addCorner(p coord, corners int) int {
	log.Printf("corner at %d", p)
	corners++
	return corners
}

func countPlot(p coord, c string, joint *map[coord]bool) (int, int) {
	if visited[p] || garden[p.y][p.x] != c {
		return 0, 0
	}

	visited[p] = true
	(*joint)[p] = true
	area := 1
	perimeter := 0

	for _, next := range getDirs(p) {
		if isOutOfBounds(next) || garden[next.y][next.x] != c {
			perimeter++
		} else {
			a, p := countPlot(next, c, joint)
			area += a
			perimeter += p
		}
	}

	return area, perimeter
}

func getDirs(p coord) []coord {
	return []coord{
		{y: p.y - 1, x: p.x}, // up
		{y: p.y + 1, x: p.x}, // down
		{y: p.y, x: p.x - 1}, // left
		{y: p.y, x: p.x + 1}, // right
	}
}

func isOutOfBounds(next coord) bool {
	return next.x < 0 || next.x >= len(garden[0]) || next.y < 0 || next.y >= len(garden)
}
