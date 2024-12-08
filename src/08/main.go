package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

var maxX = 0
var maxY = 0

var antennas = make([][]coord, 256)
var antinodes = make(map[coord]bool)

type coord struct {
	x int
	y int
}

func main() {
	readInput("input.txt")
	locateAntinodes()
	log.Printf("Solution: %d", len(antinodes))
}

func readInput(filename string) {
	file, _ := os.Open(filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	y := 0

	for y = 0; scanner.Scan(); y++ {
		for x, l := range strings.Split(scanner.Text(), "") {
			if l == "." {
				continue
			}

			a := coord{x: x, y: y}
			antennas[l[0]] = append(antennas[l[0]], a)
		}

		maxX = len(scanner.Text())
	}

	maxY = y
}

func locateAntinodes() {
	for _, as := range antennas {
		for x, a := range as {
			antinodes[a] = true // part 2 (thanks to quintal william for adding this line haha)
			compute(a, as[x+1:])
		}
	}
}

func isOutOfBounds(c coord) bool {
	return c.x < 0 || c.x >= maxX || c.y < 0 || c.y >= maxY
}

func compute(a coord, antennas []coord) {
	if len(antennas) == 0 {
		return
	}

	diff := Diff(a, antennas[0])
	c1 := Subtract(antennas[0], diff)
	c2 := Add(a, diff)

	if !isOutOfBounds(c1) {
		antinodes[c1] = true
		markLines(diff, c1) // part 2
	}
	if !isOutOfBounds(c2) {
		antinodes[c2] = true
		markLines(diff, c2) // part 2
	}

	compute(a, antennas[1:])
	compute(antennas[0], antennas[1:])
}

func markLines(diff, b coord) {
	B := b

	for {
		next := Add(b, diff)
		if isOutOfBounds(next) {
			break
		}

		antinodes[next] = true
		b = next
	}

	b = B

	for {
		next := Subtract(b, diff)
		if isOutOfBounds(next) {
			break
		}

		antinodes[next] = true
		b = next
	}
}

func Diff(a, b coord) coord {
	return coord{x: a.x - b.x, y: a.y - b.y}
}

func Subtract(a, b coord) coord {
	return coord{x: a.x - b.x, y: a.y - b.y}
}

func Add(a, b coord) coord {
	return coord{x: a.x + b.x, y: a.y + b.y}
}
