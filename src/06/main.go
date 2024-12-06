package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

var count1 = 0
var count2 = 0
var grid = make([][]string, 0)
var start coord

type coord struct {
	x int
	y int
}

type guard struct {
	pos coord
	dir coord
}

// not proud of this one, boys
func main() {
	readInput("input.txt")
	partOne()
	partTwo()                              // for the love of god, do not read partTwo code
	log.Printf("Solution 1: %d\n", count1) // 5318
	log.Printf("Solution 2: %d\n", count2)
}

func simulateMovements() error {
	g, _ := findGuard()
	start = g.pos
	visitedStates := make(map[string]bool)

	for {
		state := fmt.Sprintf("%d,%d|%d,%d", g.pos.x, g.pos.y, g.dir.x, g.dir.y) // key
		if visitedStates[state] {
			return errors.New("loop detected")
		}

		visitedStates[state] = true
		err := moveGuard(&g)
		if err != nil {
			return nil
		}
	}
}

func moveGuard(g *guard) error {
	grid[g.pos.y][g.pos.x] = "X"

	next := nextPos(*g)
	c, err := evalNext(next)

	if err != nil {
		return err
	}

	if c != "#" {
		g.pos = next
		return nil
	}

	g.dir = turnRight(g.dir)

	return nil
}

func turnRight(c coord) coord {
	switch {
	case c.x == 0 && c.y == -1:
		return coord{x: 1, y: 0}
	case c.x == 1 && c.y == 0:
		return coord{x: 0, y: 1}
	case c.x == 0 && c.y == 1:
		return coord{x: -1, y: 0}
	default:
		return coord{x: 0, y: -1}
	}
}

func evalNext(nextPos coord) (string, error) {
	if nextPos.x < 0 || nextPos.x >= len(grid[0]) || nextPos.y < 0 || nextPos.y >= len(grid) {
		return "", errors.New("guard exits grid")
	}
	return grid[nextPos.y][nextPos.x], nil
}

func findGuard() (guard, error) {
	for y, row := range grid {
		for x, e := range row {
			if e == "^" {
				return guard{pos: coord{x: x, y: y}, dir: coord{x: 0, y: -1}}, nil
			}
		}
	}
	return guard{}, errors.New("couldn't find guard")
}

func nextPos(g guard) coord {
	return coord{x: g.pos.x + g.dir.x, y: g.pos.y + g.dir.y}
}

func readInput(filename string) {
	file, _ := os.Open(filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, strings.Split(scanner.Text(), ""))
	}
}

func partOne() {
	simulateMovements()

	for _, row := range grid {
		for _, e := range row {
			if e == "X" {
				count1++
			}
		}
	}
}

func printGrid() {
	for _, e := range grid {
		log.Println(e)
	}
}

// why are you here
// reading this code will burn your eyes and rot your soul
// turn away while you still have the chance
func partTwo() int {
	candidates := []coord{}

	for y, row := range grid {
		for x, cell := range row {
			if cell == "X" && !(x == start.x && y == start.y) { // only previous guard positions (thanks william)
				candidates = append(candidates, coord{x, y})
			}
		}
	}

	for _, candidate := range candidates {
		grid[candidate.y][candidate.x] = "#"
		grid[start.y][start.x] = "^"
		err := simulateMovements()
		grid[candidate.y][candidate.x] = "."

		if err != nil {
			count2++
		}
	}
	return count2
}
