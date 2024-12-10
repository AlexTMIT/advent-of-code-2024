package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

var count1 = 0
var count2 = 0

var Map = make([][]int, 0)
var Trailheads = make([]coord, 0)

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
		Map = append(Map, make([]int, 0))
		for x, c := range strings.Split(scanner.Text(), "") {
			num, _ := strconv.Atoi(c)
			Map[y] = append(Map[y], num)

			if num == 0 {
				Trailheads = append(Trailheads, coord{x: x, y: y})
			}

		}
		y++
	}
}

func countScores() {
	for _, trailhead := range Trailheads {
		been := make(map[coord]bool)
		findPaths(trailhead, 0, been)
	}
}

func findPaths(pos coord, curr int, been map[coord]bool) {
	if curr == 9 {
		count2++
		if !been[pos] {
			been[pos] = true
			count1++
		}
		return
	}

	directions := []coord{
		{y: pos.y - 1, x: pos.x}, // up
		{y: pos.y + 1, x: pos.x}, // down
		{y: pos.y, x: pos.x - 1}, // left
		{y: pos.y, x: pos.x + 1}, // right
	}

	for _, next := range directions {
		if !isOutOfBounds(next) && Map[next.y][next.x] == curr+1 {
			findPaths(next, curr+1, been)
		}
	}
}

func isOutOfBounds(next coord) bool {
	return next.x < 0 || next.x >= len(Map[0]) || next.y < 0 || next.y >= len(Map)
}
